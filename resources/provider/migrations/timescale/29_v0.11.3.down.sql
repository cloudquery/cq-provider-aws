-- aws_redshift_clusters: revert PK
TRUNCATE TABLE aws_redshift_clusters CASCADE;
ALTER TABLE IF EXISTS aws_redshift_clusters DROP CONSTRAINT aws_redshift_clusters_pk;
ALTER TABLE IF EXISTS aws_redshift_clusters ADD CONSTRAINT aws_redshift_clusters_pk PRIMARY KEY (cq_fetch_date,account_id,id);

-- aws_redshift_snapshots: revert PK, drop cluster_cq_id, drop arn
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP CONSTRAINT aws_redshift_snapshots_pk;
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY (cq_fetch_date,cluster_identifier,cluster_create_time);
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey;
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP COLUMN IF EXISTS cluster_cq_id;
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP COLUMN IF EXISTS arn;

-- aws_redshift_event_subscriptions: revert PK, drop arn
TRUNCATE TABLE aws_redshift_event_subscriptions CASCADE;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions DROP CONSTRAINT aws_redshift_event_subscriptions_pk;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY (cq_fetch_date,account_id,id);
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions DROP COLUMN IF EXISTS arn;

-- aws_redshift_subnet_groups: revert PK
TRUNCATE TABLE aws_redshift_subnet_groups CASCADE;
ALTER TABLE IF EXISTS aws_redshift_subnet_groups DROP CONSTRAINT aws_redshift_subnet_groups_pk;
ALTER TABLE IF EXISTS aws_redshift_subnet_groups ADD CONSTRAINT aws_redshift_subnet_groups_pk PRIMARY KEY (cq_fetch_date,account_id, region, cluster_subnet_group_name);


-- Resource: wafv2.managed_rule_groups
TRUNCATE TABLE aws_wafv2_managed_rule_groups;
ALTER TABLE IF EXISTS aws_wafv2_managed_rule_groups DROP CONSTRAINT aws_wafv2_managed_rule_groups_pk;
ALTER TABLE IF EXISTS aws_wafv2_managed_rule_groups ADD CONSTRAINT aws_wafv2_managed_rule_groups_pk PRIMARY KEY (cq_fetch_date,account_id,region,vendor_name,name);
ALTER TABLE IF EXISTS "aws_wafv2_managed_rule_groups" DROP COLUMN IF EXISTS "scope";

-- Resource: wafv2.rule_groups
TRUNCATE TABLE aws_wafv2_rule_groups;
ALTER TABLE IF EXISTS "aws_wafv2_rule_groups" DROP COLUMN IF EXISTS "scope";

-- Resource: wafv2.web_acls
DELETE FROM aws_wafv2_web_acls;
ALTER TABLE IF EXISTS "aws_wafv2_web_acls" DROP COLUMN IF EXISTS "scope";