DROP TABLE IF EXISTS aws_access_analyzer_analyzer_archive_rules;

-- Resource: ec2.images
ALTER TABLE IF EXISTS "aws_ec2_images" DROP COLUMN IF EXISTS "last_launched_time";

-- Resource: ec2.security_groups
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs DROP CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk;
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs ADD CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk PRIMARY KEY (cq_id);

-- Resource: workspaces.directories
DROP TABLE IF EXISTS aws_workspaces_directories;

-- Resource: workspaces.workspaces
DROP TABLE IF EXISTS aws_workspaces_workspaces;

-- Resource: redshift.event_subscriptions
DROP TABLE IF EXISTS aws_redshift_event_subscriptions;

-- Resource: redshift.clusters
DROP TABLE IF EXISTS aws_redshift_snapshot_accounts_with_restore_access;
DROP TABLE IF EXISTS aws_redshift_snapshots;
