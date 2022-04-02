DROP TABLE IF EXISTS aws_access_analyzer_analyzer_archive_rules;

-- Resource: ec2.images
ALTER TABLE IF EXISTS "aws_ec2_images" DROP COLUMN IF EXISTS "last_launched_time";

-- Resource: ec2.security_groups
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs DROP CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk;
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs ADD CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk PRIMARY KEY (security_group_ip_permission_cq_id,group_id,user_id);


-- Resource: ecs.clusters
DROP TABLE IF EXISTS aws_ecs_cluster_task_attachments;
DROP TABLE IF EXISTS aws_ecs_cluster_task_containers;
DROP TABLE IF EXISTS aws_ecs_cluster_tasks;
ALTER TABLE IF EXISTS "aws_ecs_cluster_service_deployments"
    DROP COLUMN IF EXISTS "platform_family";
ALTER TABLE IF EXISTS "aws_ecs_cluster_service_task_sets"
    DROP COLUMN IF EXISTS "platform_family";
ALTER TABLE IF EXISTS "aws_ecs_cluster_services"
    DROP COLUMN IF EXISTS "platform_family";
DROP TABLE IF EXISTS aws_ecs_cluster_container_instance_health_status_details;
ALTER TABLE IF EXISTS "aws_ecs_cluster_container_instances"
    DROP COLUMN IF EXISTS "health_status_overall_status";