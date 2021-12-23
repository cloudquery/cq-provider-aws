package driftexport

import (
	"context"
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
	"testing"
	"text/template"

	"github.com/Masterminds/squirrel"
	"github.com/cloudquery/cq-provider-aws/resources"
	"github.com/cloudquery/cq-provider-sdk/provider/schema"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const testPrefix = "cq-testing"
const testSuffix = "integration"

type tableInfo struct {
	idFormat string
	sqlPred  interface{}
}

var exportedTables = map[string]tableInfo{
	"accessanalyzer.analyzers": {
		"{{.V.name}}",
		nil,
	},
	"apigateway.api_keys": {
		"",
		nil,
	},
	"apigateway.client_certificates": {
		"",
		nil,
	},
	"apigateway.rest_apis": {
		"",
		nil,
	},
	"apigateway.usage_plans": {
		"",
		nil,
	},
	"apigateway.vpc_links": {
		"",
		nil,
	},
	"apigatewayv2.apis": {
		"",
		nil,
	},
	"apigatewayv2.vpc_links": {
		"",
		nil,
	},
	"autoscaling.launch_configurations": {
		"{{.V.launch_configuration_name}}",
		squirrel.Eq{"launch_configuration_name": fmt.Sprintf("lc-%s-%s", testPrefix, testSuffix)},
	},
	"aws_apigateway_rest_api_authorizers": {
		"",
		nil,
	},
	"aws_apigateway_rest_api_deployments": {
		"",
		nil,
	},
	"aws_apigateway_rest_api_documentation_parts": {
		"{{.V.rest_api_id}}/{{.V.id}}",
		nil,
	},
	"aws_apigateway_rest_api_documentation_versions": {
		"{{.V.rest_api_id}}/{{.V.version}}",
		nil,
	},
	"aws_apigateway_rest_api_models": {
		"",
		nil,
	},
	"aws_apigateway_rest_api_request_validators": {
		"",
		nil,
	},
	"aws_apigateway_rest_api_resources": {
		"",
		map[string]interface{}{"path_part": "gateway_resource_1"}, // some other paths get created automatically (?) but tf state has only explicit one
	},
	"aws_apigateway_rest_api_stages": {
		"ags-{{.Parent.V.id}}-{{.V.stage_name}}",
		nil,
	},
	"aws_apigateway_usage_plan_api_stages": {
		"{{.V.usage_plan_id}}|{{.V.api_id}}|{{.V.stage}}",
		nil,
	},
	"aws_apigateway_usage_plan_keys": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_authorizers": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_deployments": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_integration_responses": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_integrations": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_models": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_route_responses": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_routes": {
		"",
		nil,
	},
	"aws_apigatewayv2_api_stages": {
		"",
		nil,
	},
	"aws_autoscaling_launch_configuration_block_device_mappings": {
		"{{.Parent.V.launch_configuration_name}}|{{.V.device_name}}",
		nil,
	},
	"aws_cloudfront_distribution_cache_behaviors": {
		"{{.Parent.V.id}}|{{.V.path_pattern}}|{{.V.target_origin_id}}|{{.V.viewer_protocol_policy}}",
		nil,
	},
	"aws_cloudfront_distribution_custom_error_responses": {
		"{{.Parent.V.id}}|{{.V.error_code}}|{{.V.response_code}}|{{.V.response_page_path}}",
		nil,
	},
	"aws_cloudfront_distribution_origins": {
		`{{split_part .V.s3_origin_config_origin_access_identity "/" 3}}`,
		nil,
	},
	"aws_cloudtrail_trail_event_selectors": {
		"{{.Parent.V.name}}|{{.V.include_management_events}}|{{.V.read_write_type}}",
		nil,
	},
	"aws_cloudwatch_alarm_metrics": {
		"{{.Parent.V.name}}|{{.V.id}}",
		nil,
	},
	"aws_cloudwatchlogs_filter_metric_transformations": {
		"{{.Parent.V.name}}|{{.V.metric_namespace}}|{{.V.metric_name}}",
		nil,
	},
	"aws_directconnect_gateway_associations": {
		"ga-{{.V.gateway_id}}{{.V.associated_gateway_id}}",
		nil,
	},
	"aws_ec2_ebs_volume_attachments": {
		"{{.V.instance_id}}|{{.V.volume_id}}|{{.V.device}}",
		nil,
	},
	"aws_iam_user_access_keys": {
		"",
		nil,
	},
	"aws_iam_user_attached_policies": {
		"{{.Parent.V.user_name}}:user_{{.V.policy_name}}",
		nil,
	},
	"aws_iam_user_groups": {
		"{{.V.group_name}}",
		nil,
	},
	"aws_iam_user_policies": {
		"{{.Parent.V.user_name}}:{{.V.policy_name}}",
		nil,
	},
	// commented out because:
	// * extra objects are present beyond those in tf state file
	// * they could be filtered out based on parent, but parent table is not supported in drift
	//	 (squirrel.Eq{"name": fmt.Sprintf("lambda_layer%s%s", testPrefix, testSuffix)})
	//
	// "aws_lambda_layer_versions": {
	// 	"{{.Parent.V.arn}}:{{.V.version}}",
	// 	nil,
	// },
	"cloudfront.cache_policies": {
		"",
		map[string]interface{}{"name": fmt.Sprintf("cache_policy%s-%s", testPrefix, testSuffix)},
	},
	"cloudfront.distributions": {
		"{{.V.id}}",
		nil,
	},
	"cloudtrail.trails": {
		"{{.V.name}}",
		nil,
	},
	"cloudwatch.alarms": {
		"{{.V.name}}",
		map[string]interface{}{"name": fmt.Sprintf("cw-alarm%s-%s", testPrefix, testSuffix)},
	},
	"cloudwatchlogs.filters": {
		"{{.V.name}}|{{.V.log_group_name}}",
		nil,
	},
	"cognito.identity_pools": {
		"",
		nil,
	},
	"cognito.user_pools": {
		"",
		map[string]interface{}{"name": fmt.Sprintf("cognito_user_pool%s-%s", testPrefix, testSuffix)},
	},
	"config.configuration_recorders": {
		"{{.V.name}}",
		map[string]interface{}{"name": fmt.Sprintf("config-cr-%s-%s", testPrefix, testSuffix)},
	},
	"config.conformance_packs": {
		"{{.V.conformance_pack_name}}",
		nil,
	},
	"directconnect.connections": {
		"",
		nil,
	},
	"directconnect.gateways": {
		"",
		nil,
	},
	"directconnect.lags": {
		"",
		nil,
	},
	"directconnect.virtual_interfaces": {
		"",
		nil,
	},
	"ec2.customer_gateways": {
		"",
		nil,
	},
	"ec2.ebs_volumes": {
		"",
		map[string]interface{}{"size": 5},
	},
	"ec2.flow_logs": {
		"",
		nil,
	},
	"ec2.images": {
		"{{.V.tags.Ec2ImageBuilderArn}}",
		nil,
	},
	"ec2.instances": {
		"",
		squirrel.Or{
			squirrel.Expr("tags ->> 'Name' IS NOT DISTINCT FROM ?", fmt.Sprintf("ec2_instance%s", testSuffix)),
			squirrel.Expr("tags ->> 'Name' IS NOT DISTINCT FROM ?", fmt.Sprintf("ecs_ec2_instance%s", testSuffix)),
			squirrel.Expr("tags ->> 'Name' IS NOT DISTINCT FROM ?", fmt.Sprintf("elbv1-instance-1%s", testSuffix)),
		},
	},
	"ec2.internet_gateways": {
		"",
		nil,
	},
	"ec2.nat_gateways": {
		"",
		nil,
	},
	"ec2.network_acls": {
		"",
		nil,
	},
	"ec2.route_tables": {
		"",
		nil,
	},
	"ec2.security_groups": {
		"",
		squirrel.Or{
			squirrel.Eq{"group_name": fmt.Sprintf("apigw-sg-%s-%s", testPrefix, testSuffix)},
			squirrel.Eq{"group_name": fmt.Sprintf("dms-sg-%s_%s", testPrefix, testSuffix)},
			squirrel.Eq{"group_name": fmt.Sprintf("aws_ec2_instances_sg_%s%s", testPrefix, testSuffix)},
			squirrel.Eq{"group_name": fmt.Sprintf("ec2-sg-%s%s", testPrefix, testSuffix)},
			squirrel.Eq{"group_name": fmt.Sprintf("ecs_clusters_sg%s%s", testPrefix, testSuffix)},
			squirrel.Eq{"group_name": fmt.Sprintf("aws_emr_clusters_security_group%s%s", testPrefix, testSuffix)},
			squirrel.Eq{"group_name": fmt.Sprintf("mq_test_sg_%s_%s", testPrefix, testSuffix)},
		},
	},
	"ec2.subnets": {
		"",
		nil,
	},
	"ec2.transit_gateways": {
		"",
		nil,
	},
	"ec2.vpc_endpoints": {
		"",
		nil,
	},
	"ec2.vpc_peering_connections": {
		"",
		nil,
	},
	"ec2.vpcs": {
		"",
		nil,
	},
	"ec2.vpn_gateways": {
		"",
		nil,
	},
	"ecr.repositories": {
		"{{.V.name}}",
		map[string]interface{}{"name": fmt.Sprintf("ecr_repositories_%s%s", testPrefix, testSuffix)},
	},
	"ecs.clusters": {
		"{{.V.arn}}",
		nil,
	},
	"efs.filesystems": {
		"",
		nil,
	},
	"eks.clusters": {
		"{{.V.name}}",
		nil,
	},
	"elasticbeanstalk.environments": {
		"",
		nil,
	},
	"elasticsearch.domains": {
		"{{.V.arn}}",
		nil,
	},
	"elbv1.load_balancers": {
		"",
		nil,
	},
	"elbv2.load_balancers": {
		"",
		squirrel.Or{
			squirrel.Eq{"name": fmt.Sprintf("apigateway-nlb-%s", testSuffix)},
			squirrel.Eq{"name": fmt.Sprintf("elbv2-%s", testSuffix)},
		},
	},
	"elbv2.target_groups": {
		"",
		squirrel.Or{
			squirrel.Eq{"name": fmt.Sprintf("elbv2-tg-%s", testSuffix)},
			squirrel.Eq{"name": fmt.Sprintf("lbv2target%s", testPrefix)},
		},
	},
	"emr.clusters": {
		"{{.V.id}}",
		squirrel.Eq{"state": "WAITING"},
	},
	"fsx.backups": {
		"",
		nil,
	},
	"iam.groups": {
		"{{.V.name}}",
		map[string]interface{}{"name": fmt.Sprintf("aws_iam_group%s%s", testPrefix, testSuffix)},
	},
	"iam.openid_connect_identity_providers": {
		"",
		nil,
	},
	"iam.policies": {
		"{{.V.arn}}",
		map[string]interface{}{"name": []string{
			fmt.Sprintf("iam_policy_%s%s", testPrefix, testSuffix),
			fmt.Sprintf("policy%s%s", testPrefix, testSuffix),
			fmt.Sprintf("lambda_%s%s", testPrefix, testSuffix),
		}},
	},
	"iam.roles": {
		"{{.V.name}}",
		nil,
	},
	"iam.saml_identity_providers": {
		"",
		nil,
	},
	"iam.server_certificates": {
		"",
		nil,
	},
	"iam.users": {
		"{{.V.user_name}}",
		nil,
	},
	"kms.keys": {
		"{{.V.id}}",
		map[string]interface{}{"description": fmt.Sprintf("kms-key-%s%s", testPrefix, testSuffix)},
	},
	"lambda.functions": {
		"{{.V.name}}",
		squirrel.Or{
			squirrel.Eq{"name": fmt.Sprintf("function_%s%s", testPrefix, testSuffix)},
			squirrel.Eq{"name": fmt.Sprintf("secretsmanager-secret-rotation-function-%s%s", testPrefix, testSuffix)},
		},
	},
	"mq.brokers": {
		"",
		nil,
	},
	"rds.clusters": {
		"{{.V.db_cluster_identifier}}",
		nil,
	},
	"rds.db_subnet_groups": {
		"{{.V.name}}",
		map[string]interface{}{"name": fmt.Sprintf("rds_db_subnet%s%s", testPrefix, testSuffix)},
	},
	"rds.instances": {
		"{{.V.db_name}}",
		map[string]interface{}{"user_instance_id": fmt.Sprintf("rdsclusterdb%s", testSuffix)},
	},
	"redshift.clusters": {
		"",
		nil,
	},
	"redshift.subnet_groups": {
		"",
		nil,
	},
	"route53.health_checks": {
		"",
		nil,
	},
	"route53.hosted_zones": {
		"",
		nil,
	},
	"route53.reusable_delegation_sets": {
		`{{split_part .V.id "/" 3}}`,
		nil,
	},
	"s3.buckets": {
		"",
		map[string]interface{}{
			"name": []string{
				fmt.Sprintf("cf-buc-%s-%s", testPrefix, testSuffix),
				fmt.Sprintf("cloudtrail-buc-%s-%s", testPrefix, testSuffix),
				fmt.Sprintf("cloudtrail-target-buc-%s-%s", testPrefix, testSuffix),
				fmt.Sprintf("codebuild%s%s", testPrefix, testSuffix),
				fmt.Sprintf("ec2-fl-buck%s%s", testPrefix, testSuffix),
				fmt.Sprintf("ec2-images-logs-%s%s", testPrefix, testSuffix),
				fmt.Sprintf("elbv1-bucket%s%s", testPrefix, testSuffix),
				fmt.Sprintf("emr-cluster-logs%s%s", testPrefix, testSuffix),
				fmt.Sprintf("kinesis-firehose-bucket-%s%s", testPrefix, testSuffix),
				fmt.Sprintf("bucket-%s%s", testPrefix, testSuffix),
			},
		},
	},
	"sns.subscriptions": {
		"",
		nil,
	},
	"sns.topics": {
		"",
		map[string]interface{}{
			"display_name": []string{
				fmt.Sprintf("ag-topic-%s", testSuffix),
				fmt.Sprintf("lambda_func_user-updates-topic%s%s", testPrefix, testSuffix),
				fmt.Sprintf("lambda_func_errors-topic%s%s", testPrefix, testSuffix),
				fmt.Sprintf("sns-tests-topic-%s.fifo", testSuffix),
				fmt.Sprintf("sns-test2-%s-%s", testPrefix, testSuffix),
			},
		},
	},
	"sqs.queues": {
		"{{.V.url}}",
		nil,
	},
	"waf.rule_groups": {
		"",
		nil,
	},
	"waf.rules": {
		"",
		nil,
	},
	"waf.web_acls": {
		"",
		nil,
	},
	"wafv2.rule_groups": {
		"",
		nil,
	},
	"wafv2.web_acls": {
		"",
		nil,
	},
}

func TestExportDataForDriftE2E(t *testing.T) {
	if os.Getenv("DRIFT_DATA_EXPORT") == "" {
		t.Skip()
	}
	provider := resources.Provider()
	ctx := context.Background()
	pool := setupDatabase(t, ctx)
	defer pool.Close()
	conn, err := pool.Acquire(ctx)
	require.Nil(t, err)
	defer conn.Release()

	all, err := loadExistingResources(ctx, conn, provider.ResourceMap)
	require.Nil(t, err)

	idBuilders := makeIDBuilders(provider.ResourceMap)
	known := exportRecords(all, idBuilders, nil)

	b, err := yaml.Marshal(known)
	require.Nil(t, err)
	require.Nil(t, ioutil.WriteFile("output.yaml", b, 0644))
}

func coalesceString(values ...string) string {
	for _, v := range values {
		if v != "" {
			return v
		}
	}
	return ""
}

func setupDatabase(t *testing.T, ctx context.Context) *pgxpool.Pool {
	dbCfg, err := pgxpool.ParseConfig(coalesceString(os.Getenv("DATABASE_URL"), "host=localhost user=postgres password=pass dbname=postgres port=5432"))
	require.Nil(t, err)
	dbCfg.LazyConnect = true
	pool, err := pgxpool.ConnectConfig(ctx, dbCfg)
	require.Nil(t, err)
	return pool
}

type record struct {
	data     map[string]interface{}
	children map[string][]record
}

func loadExistingResources(ctx context.Context, conn *pgxpool.Conn, resourceMap map[string]*schema.Table) (map[string][]record, error) {
	r := make(map[string][]record)
	for resourceName, table := range resourceMap {
		var err error
		r[resourceName], err = loadTable(ctx, conn, table, resourceName, nil)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func createFilter(parent *record, table *schema.Table, resource string, q squirrel.SelectBuilder) (squirrel.SelectBuilder, error) {
	if info, ok := exportedTables[resource]; ok && info.sqlPred != nil {
		q = q.Where(info.sqlPred)
	}
	if parent != nil {
		id, ok := parent.data["cq_id"]
		if !ok {
			return q, fmt.Errorf("cq_id field is missing in a parent of %s", table.Name)
		}
		var fk string
		for _, col := range table.Columns {
			if strings.HasSuffix(col.Name, "_cq_id") {
				fk = col.Name
				break
			}
		}
		if fk == "" {
			return q, fmt.Errorf("foreign key not found for the table %s", table.Name)
		}
		q = q.Where(squirrel.Eq{fk: id})
	}
	return q, nil
}

func loadTable(ctx context.Context, conn *pgxpool.Conn, table *schema.Table, resource string, parent *record) ([]record, error) {
	var result []record
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	q := psql.Select(fmt.Sprintf("to_jsonb(%s)", table.Name)).From(table.Name)
	q, err := createFilter(parent, table, resource, q)
	if err != nil {
		return nil, err
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var row record
		if err := rows.Scan(&row.data); err != nil {
			return nil, err
		}
		if isInterestingRecord(row.data) {
			result = append(result, row)
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	for i, row := range result {
		row.children = make(map[string][]record)
		for _, relation := range table.Relations {
			var err error
			row.children[relation.Name], err = loadTable(ctx, conn, relation, relation.Name, &row)
			if err != nil {
				return nil, err
			}
		}
		result[i] = row
	}
	return result, nil
}

func isInterestingRecord(row map[string]interface{}) bool {
	tagsValue, ok := row["tags"]
	if !ok {
		return true
	}
	tags, ok := tagsValue.(map[string]interface{})
	if !ok {
		return true
	}
	return reflect.DeepEqual(tags["Type"], "integration_test") && reflect.DeepEqual(tags["TestId"], testSuffix)
}

func makeIDBuilders(resourceMap map[string]*schema.Table) map[string]func([]map[string]interface{}) string {
	b := make(map[string]func([]map[string]interface{}) string)
	for name, table := range resourceMap {
		info, ok := exportedTables[name]
		if ok {
			b[name] = makeIDBuilder(info.idFormat, table)
		}
		children := make(map[string]*schema.Table)
		for _, t := range table.Relations {
			children[t.Name] = t
		}
		for name, f := range makeIDBuilders(children) {
			b[name] = f
		}
	}
	return b
}

func makeIDBuilder(expr string, table *schema.Table) func([]map[string]interface{}) string {
	ignored := ignoredColumns(table)
	if expr == "" {
		keys := make([]string, 0, len(table.Options.PrimaryKeys))
		for _, k := range table.Options.PrimaryKeys {
			if _, ok := ignored[k]; !ok {
				keys = append(keys, k)
			}
		}
		return func(recs []map[string]interface{}) string {
			values := make([]string, 0, len(keys))
			for _, k := range keys {
				values = append(values, fmt.Sprintf("%s", recs[0][k]))
			}
			return strings.Join(values, "|")
		}
	}
	return func(rows []map[string]interface{}) string {
		t := template.New("id")
		t.Funcs(template.FuncMap{"split_part": sqlSplitPart})
		template.Must(t.Parse(expr))
		type record struct {
			V      map[string]interface{}
			Parent *record
		}
		var last *record
		for i := len(rows) - 1; i >= 0; i-- {
			r := record{V: rows[i], Parent: last}
			last = &r
		}
		var sb strings.Builder
		if err := t.Execute(&sb, last); err != nil {
			panic(err)
		}
		return sb.String()
	}
}

// sqlSplitPart is analog to Postgres' split_part function.
func sqlSplitPart(s string, sep string, n int) string {
	items := strings.Split(s, sep)
	if n-1 < len(items) {
		return items[n-1]
	}
	return ""
}

func exportRecords(all map[string][]record, idBuilder map[string]func([]map[string]interface{}) string, parents []map[string]interface{}) map[string][]string {
	known := make(map[string][]string)
	for name, recs := range all {
		for _, rec := range recs {
			recs := append([]map[string]interface{}{rec.data}, parents...)
			if _, ok := exportedTables[name]; ok {
				if _, ok := idBuilder[name]; !ok {
					panic(fmt.Sprintf("id builder not found for %s", name))
				}
				id := idBuilder[name](recs)
				known[name] = append(known[name], id)
			}
			for k, v := range exportRecords(rec.children, idBuilder, recs) {
				known[k] = append(known[k], v...)
			}
		}
	}
	return known
}

func ignoredColumns(t *schema.Table) map[string]struct{} {
	ignored := make(map[string]struct{})
	for _, c := range t.Columns {
		m := c.Meta()
		if m == nil || m.Resolver == nil {
			continue
		}
		switch m.Resolver.Name {
		case "schema.ParentIdResolver",
			"github.com/cloudquery/cq-provider-aws/client.ResolveAWSAccount",
			"github.com/cloudquery/cq-provider-aws/client.ResolveAWSRegion":
			ignored[c.Name] = struct{}{}
		}
	}
	return ignored
}
