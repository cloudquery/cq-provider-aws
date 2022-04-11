-- Autogenerated by migration tool on 2022-04-04 17:12:37
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: codepipeline.pipelines
DROP TABLE IF EXISTS aws_codepipeline_pipeline_stage_actions;
DROP TABLE IF EXISTS aws_codepipeline_pipeline_stages;
DROP TABLE IF EXISTS aws_codepipeline_pipelines;

-- Resource: codepipeline.webhooks
DROP TABLE IF EXISTS aws_codepipeline_webhook_filters;
DROP TABLE IF EXISTS aws_codepipeline_webhooks;

-- Resource: sns.subscriptions
ALTER TABLE IF EXISTS aws_sns_subscriptions DROP CONSTRAINT aws_sns_subscriptions_pk;
ALTER TABLE IF EXISTS aws_sns_subscriptions ADD CONSTRAINT aws_sns_subscriptions_pk PRIMARY KEY (arn);

-- Resource: mq.brokers
ALTER TABLE IF EXISTS aws_mq_brokers DROP CONSTRAINT aws_mq_brokers_pk;
ALTER TABLE IF EXISTS "aws_mq_brokers" RENAME COLUMN id TO broker_id;
ALTER TABLE IF EXISTS aws_mq_brokers ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (account_id,broker_id);
ALTER TABLE IF EXISTS aws_mq_broker_configurations DROP CONSTRAINT aws_mq_broker_configurations_pk;
ALTER TABLE IF EXISTS aws_mq_broker_configurations ADD CONSTRAINT aws_mq_broker_configurations_pk PRIMARY KEY (broker_cq_id,id);
DROP TABLE IF EXISTS aws_mq_broker_configuration_revisions;
ALTER TABLE IF EXISTS aws_mq_broker_users DROP CONSTRAINT aws_mq_broker_users_pk;
ALTER TABLE IF EXISTS aws_mq_broker_users ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (broker_cq_id,username);

-- Resource: iam.virtual_mfa_devices
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices DROP CONSTRAINT aws_iam_virtual_mfa_devices_pk;
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices ADD CONSTRAINT aws_iam_virtual_mfa_devices_pk PRIMARY KEY (cq_id);

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