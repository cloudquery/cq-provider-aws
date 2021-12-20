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

var exportedTables = map[string]string{
	"accessanalyzer.analyzers":                                   "{{.V.name}}",
	"apigateway.api_keys":                                        "",
	"apigateway.client_certificates":                             "",
	"apigateway.rest_apis":                                       "",
	"apigateway.usage_plans":                                     "",
	"apigateway.vpc_links":                                       "",
	"apigatewayv2.apis":                                          "",
	"apigatewayv2.vpc_links":                                     "",
	"autoscaling.launch_configurations":                          "{{.V.launch_configuration_name}}",
	"aws_apigateway_rest_api_authorizers":                        "",
	"aws_apigateway_rest_api_deployments":                        "",
	"aws_apigateway_rest_api_documentation_parts":                "{{.V.rest_api_id}}/{{.V.id}}",
	"aws_apigateway_rest_api_documentation_versions":             "{{.V.rest_api_id}}/{{.V.version}}",
	"aws_apigateway_rest_api_models":                             "",
	"aws_apigateway_rest_api_request_validators":                 "",
	"aws_apigateway_rest_api_resources":                          "",
	"aws_apigateway_rest_api_stages":                             "ags{{.Parent.V.id}}-{{.V.stage_name}}",
	"aws_apigateway_usage_plan_api_stages":                       "{{.V.usage_plan_id}}{{.V.api_id}}{{.V.stage}}",
	"aws_apigateway_usage_plan_keys":                             "",
	"aws_apigatewayv2_api_authorizers":                           "",
	"aws_apigatewayv2_api_deployments":                           "",
	"aws_apigatewayv2_api_integration_responses":                 "",
	"aws_apigatewayv2_api_integrations":                          "",
	"aws_apigatewayv2_api_models":                                "",
	"aws_apigatewayv2_api_route_responses":                       "",
	"aws_apigatewayv2_api_routes":                                "",
	"aws_apigatewayv2_api_stages":                                "",
	"aws_autoscaling_launch_configuration_block_device_mappings": "{{.Parent.V.launch_configuration_name}}{{.V.device_name}}",
	"aws_cloudfront_distribution_cache_behaviors":                "{{.Parent.V.id}}{{.V.path_pattern}}{{.V.target_origin_id}}{{.V.viewer_protocol_policy}}",
	"aws_cloudfront_distribution_custom_error_responses":         "{{.Parent.V.id}}{{.V.error_code}}{{.V.response_code}}{{.V.response_page_path}}",
	"aws_cloudfront_distribution_origins":                        "", // sql("SPLIT_PART(c.s3_origin_config_origin_access_identity,'/', 3)")
	"aws_cloudtrail_trail_event_selectors":                       "{{.Parent.V.name}}{{.V.include_management_events}}{{.V.read_write_type}}",
	"aws_cloudwatch_alarm_metrics":                               "{{.Parent.V.name}}{{.V.id}}",
	"aws_cloudwatchlogs_filter_metric_transformations":           "{{.Parent.V.name}}{{.V.metric_namespace}}{{.V.metric_name}}",
	"aws_directconnect_gateway_associations":                     "ga-{{.V.gateway_id}}{{.V.associated_gateway_id}}",
	"aws_ec2_ebs_volume_attachments":                             "{{.V.instance_id}}{{.V.volume_id}}{{.V.device}}",
	"aws_iam_user_access_keys":                                   "",
	"aws_iam_user_attached_policies":                             "{{.Parent.V.user_name}}:user_{{.V.policy_name}}",
	"aws_iam_user_groups":                                        "{{.V.group_name}}",
	"aws_iam_user_policies":                                      "{{.Parent.V.user_name}}:{{.V.policy_name}}",
	"aws_lambda_layer_versions":                                  "{{.Parent.V.arn}}:{{.V.version}}",
	"cloudfront.cache_policies":                                  "",
	"cloudfront.distributions":                                   "{{.V.id}}",
	"cloudtrail.trails":                                          "{{.V.name}}",
	"cloudwatch.alarms":                                          "{{.V.name}}",
	"cloudwatchlogs.filters":                                     "{{.V.name}}{{.V.log_group_name}}",
	"cognito.identity_pools":                                     "",
	"cognito.user_pools":                                         "",
	"config.configuration_recorders":                             "{{.V.name}}",
	"config.conformance_packs":                                   "{{.V.conformance_pack_name}}",
	"directconnect.connections":                                  "",
	"directconnect.gateways":                                     "",
	"directconnect.lags":                                         "",
	"directconnect.virtual_interfaces":                           "",
	"ec2.customer_gateways":                                      "",
	"ec2.ebs_volumes":                                            "",
	"ec2.flow_logs":                                              "",
	"ec2.images":                                                 "{{.V.tags.Ec2ImageBuilderArn}}",
	"ec2.instances":                                              "",
	"ec2.internet_gateways":                                      "",
	"ec2.nat_gateways":                                           "",
	"ec2.network_acls":                                           "",
	"ec2.route_tables":                                           "",
	"ec2.security_groups":                                        "",
	"ec2.subnets":                                                "",
	"ec2.transit_gateways":                                       "",
	"ec2.vpc_endpoints":                                          "",
	"ec2.vpc_peering_connections":                                "",
	"ec2.vpcs":                                                   "",
	"ec2.vpn_gateways":                                           "",
	"ecr.repositories":                                           "{{.V.name}}",
	"ecs.clusters":                                               "{{.V.arn}}",
	"efs.filesystems":                                            "",
	"eks.clusters":                                               "{{.V.name}}",
	"elasticbeanstalk.environments":                              "",
	"elasticsearch.domains":                                      "{{.V.arn}}",
	"elbv1.load_balancers":                                       "",
	"elbv2.load_balancers":                                       "",
	"elbv2.target_groups":                                        "",
	"emr.clusters":                                               "",
	"fsx.backups":                                                "",
	"iam.groups":                                                 "{{.V.name}}",
	"iam.openid_connect_identity_providers":                      "",
	"iam.policies":                                               "{{.V.arn}}",
	"iam.roles":                                                  "{{.V.name}}",
	"iam.saml_identity_providers":                                "",
	"iam.server_certificates":                                    "",
	"iam.users":                                                  "{{.V.user_name}}",
	"kms.keys":                                                   "{{.V.id}}",
	"lambda.functions":                                           "{{.V.name}}",
	"mq.brokers":                                                 "",
	"rds.clusters":                                               "{{.V.db_cluster_identifier}}",
	"rds.db_subnet_groups":                                       "{{.V.name}}",
	"rds.instances":                                              "{{.V.db_name}}",
	"redshift.clusters":                                          "",
	"redshift.subnet_groups":                                     "",
	"route53.health_checks":                                      "",
	"route53.hosted_zones":                                       "",
	"route53.reusable_delegation_sets":                           "", // sql("SPLIT_PART(c.id, '/', 3)")
	"s3.buckets":                                                 "",
	"sns.subscriptions":                                          "",
	"sns.topics":                                                 "",
	"sqs.queues":                                                 "{{.V.url}}",
	"waf.rule_groups":                                            "",
	"waf.rules":                                                  "",
	"waf.web_acls":                                               "",
	"wafv2.rule_groups":                                          "",
	"wafv2.web_acls":                                             "",
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
		r[resourceName], err = loadTable(ctx, conn, table, nil)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func loadTable(ctx context.Context, conn *pgxpool.Conn, table *schema.Table, parent *record) ([]record, error) {
	var result []record
	psql := squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	q := psql.Select(fmt.Sprintf("to_jsonb(%s)", table.Name)).From(table.Name)
	if parent != nil {
		id, ok := parent.data["cq_id"]
		if !ok {
			return nil, fmt.Errorf("cq_id field is missing in a parent of %s", table.Name)
		}
		var fk string
		for _, col := range table.Columns {
			if strings.HasSuffix(col.Name, "_cq_id") {
				fk = col.Name
				break
			}
		}
		if fk == "" {
			return nil, fmt.Errorf("foreign key not found for the table %s", table.Name)
		}
		q = q.Where(squirrel.Eq{fk: id})
	}
	sql, args, err := q.ToSql()
	if err != nil {
		return nil, err
	}
	rows, err := conn.Query(ctx, sql, args...)
	if err != nil {
		return nil, err
	}
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
			row.children[relation.Name], err = loadTable(ctx, conn, relation, &row)
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
	return reflect.DeepEqual(tags["Type"], "integration_test") && reflect.DeepEqual(tags["TestId"], "integration")
}

func makeIDBuilders(resourceMap map[string]*schema.Table) map[string]func([]map[string]interface{}) string {
	b := make(map[string]func([]map[string]interface{}) string)
	for name, table := range resourceMap {
		keys, ok := exportedTables[name]
		if ok {
			b[name] = makeIDBuilder(keys, table)
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
	if expr == "" {
		keys := make([]string, 0, len(table.Options.PrimaryKeys))
		for _, k := range table.Options.PrimaryKeys {
			if k != "account_id" && k != "region" && k != "cq_id" && !strings.HasSuffix(k, "_cq_id") {
				keys = append(keys, k)
			}
		}
		return func(recs []map[string]interface{}) string {
			values := make([]string, 0, len(keys))
			for _, k := range keys {
				values = append(values, fmt.Sprintf("%s", recs[0][k]))
			}
			return strings.Join(values, "")
		}
	}
	return func(rows []map[string]interface{}) string {
		t := template.Must(template.New("id").Parse(expr))
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
		t.Execute(&sb, last)
		return sb.String()
	}
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
