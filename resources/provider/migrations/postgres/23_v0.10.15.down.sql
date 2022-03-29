DROP TABLE IF EXISTS aws_access_analyzer_analyzer_archive_rules;

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