-- Autogenerated by migration tool on 2022-04-04 17:12:37
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: codepipeline.pipelines
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipelines" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "created" timestamp without time zone,
    "arn" text,
    "updated" timestamp without time zone,
    "name" text,
    "role_arn" text,
    "artifact_store_location" text,
    "artifact_store_type" text,
    "artifact_store_encryption_key_id" text,
    "artifact_store_encryption_key_type" text,
    "artifact_stores" jsonb,
    "version" integer,
    CONSTRAINT aws_codepipeline_pipelines_pk PRIMARY KEY(arn),
    UNIQUE(cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipeline_stages" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "pipeline_cq_id" uuid,
    "stage_order" integer,
    "name" text,
    "blockers" jsonb,
    CONSTRAINT aws_codepipeline_pipeline_stages_pk PRIMARY KEY(cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (pipeline_cq_id) REFERENCES aws_codepipeline_pipelines(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipeline_stage_actions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "pipeline_stage_cq_id" uuid,
    "category" text,
    "owner" text,
    "provider" text,
    "version" text,
    "name" text,
    "configuration" jsonb,
    "input_artifacts" text[],
    "namespace" text,
    "output_artifacts" text[],
    "region" text,
    "role_arn" text,
    "run_order" integer,
    CONSTRAINT aws_codepipeline_pipeline_stage_actions_pk PRIMARY KEY(cq_id),
    UNIQUE(cq_id),
    FOREIGN KEY (pipeline_stage_cq_id) REFERENCES aws_codepipeline_pipeline_stages(cq_id) ON DELETE CASCADE
);

-- Resource: sns.subscriptions
ALTER TABLE IF EXISTS aws_sns_subscriptions DROP CONSTRAINT aws_sns_subscriptions_pk;
ALTER TABLE IF EXISTS aws_sns_subscriptions ADD CONSTRAINT aws_sns_subscriptions_pk PRIMARY KEY (endpoint,owner,protocol,arn,topic_arn);

-- Resource: mq.brokers
ALTER TABLE IF EXISTS aws_mq_brokers
    DROP CONSTRAINT aws_mq_brokers_pk;
ALTER TABLE IF EXISTS "aws_mq_brokers"
    RENAME COLUMN broker_id TO id;
ALTER TABLE IF EXISTS aws_mq_brokers
    ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (account_id, id);
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    DROP CONSTRAINT aws_mq_broker_configurations_pk;
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    ADD CONSTRAINT aws_mq_broker_configurations_pk PRIMARY KEY (cq_id);
CREATE TABLE IF NOT EXISTS "aws_mq_broker_configuration_revisions"
(
    "cq_id"                      uuid NOT NULL,
    "cq_meta"                    jsonb,
    "broker_configuration_cq_id" uuid,
    "configuration_id"           text,
    "created"                    timestamp without time zone,
    "data"                       jsonb,
    "description"                text,
    CONSTRAINT aws_mq_broker_configuration_revisions_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (broker_configuration_cq_id) REFERENCES aws_mq_broker_configurations (cq_id) ON DELETE CASCADE
);
ALTER TABLE IF EXISTS aws_mq_broker_users
    DROP CONSTRAINT aws_mq_broker_users_pk;
ALTER TABLE IF EXISTS aws_mq_broker_users
    ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (cq_id);


-- Resource: ecs.clusters
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_tasks"
(
    "cq_id"                         uuid NOT NULL,
    "cq_meta"                       jsonb,
    "cluster_cq_id"                 uuid,
    "attributes"                    jsonb,
    "availability_zone"             text,
    "capacity_provider_name"        text,
    "cluster_arn"                   text,
    "connectivity"                  text,
    "connectivity_at"               timestamp without time zone,
    "container_instance_arn"        text,
    "cpu"                           text,
    "created_at"                    timestamp without time zone,
    "desired_status"                text,
    "enable_execute_command"        boolean,
    "ephemeral_storage_size_in_gib" integer,
    "execution_stopped_at"          timestamp without time zone,
    "group"                         text,
    "health_status"                 text,
    "inference_accelerators"        jsonb,
    "last_status"                   text,
    "launch_type"                   text,
    "memory"                        text,
    "overrides"                     jsonb,
    "platform_family"               text,
    "platform_version"              text,
    "pull_started_at"               timestamp without time zone,
    "pull_stopped_at"               timestamp without time zone,
    "started_at"                    timestamp without time zone,
    "started_by"                    text,
    "stop_code"                     text,
    "stopped_at"                    timestamp without time zone,
    "stopped_reason"                text,
    "stopping_at"                   timestamp without time zone,
    "tags"                          jsonb,
    "arn"                           text,
    "task_definition_arn"           text,
    "version"                       bigint,
    CONSTRAINT aws_ecs_cluster_tasks_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (cluster_cq_id) REFERENCES aws_ecs_clusters (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_task_attachments"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "cluster_task_cq_id" uuid,
    "details"            jsonb,
    "id"                 text,
    "status"             text,
    "type"               text,
    CONSTRAINT aws_ecs_cluster_task_attachments_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (cluster_task_cq_id) REFERENCES aws_ecs_cluster_tasks (cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_task_containers"
(
    "cq_id"              uuid NOT NULL,
    "cq_meta"            jsonb,
    "cluster_task_cq_id" uuid,
    "container_arn"      text,
    "cpu"                text,
    "exit_code"          integer,
    "gpu_ids"            text[],
    "health_status"      text,
    "image"              text,
    "image_digest"       text,
    "last_status"        text,
    "managed_agents"     jsonb,
    "memory"             text,
    "memory_reservation" text,
    "name"               text,
    "network_bindings"   jsonb,
    "network_interfaces" jsonb,
    "reason"             text,
    "runtime_id"         text,
    "task_arn"           text,
    CONSTRAINT aws_ecs_cluster_task_containers_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (cluster_task_cq_id) REFERENCES aws_ecs_cluster_tasks (cq_id) ON DELETE CASCADE
);
ALTER TABLE IF EXISTS "aws_ecs_cluster_services"
    ADD COLUMN IF NOT EXISTS "platform_family" text;
ALTER TABLE IF EXISTS "aws_ecs_cluster_service_deployments"
    ADD COLUMN IF NOT EXISTS "platform_family" text;
ALTER TABLE IF EXISTS "aws_ecs_cluster_service_task_sets"
    ADD COLUMN IF NOT EXISTS "platform_family" text;
ALTER TABLE IF EXISTS "aws_ecs_cluster_container_instances"
    ADD COLUMN IF NOT EXISTS "health_status_overall_status" text;
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_container_instance_health_status_details"
(
    "cq_id"                            uuid NOT NULL,
    "cq_meta"                          jsonb,
    "cluster_container_instance_cq_id" uuid,
    "last_status_change"               timestamp without time zone,
    "last_updated"                     timestamp without time zone,
    "status"                           text,
    "type"                             text,
    CONSTRAINT aws_ecs_cluster_container_instance_health_status_details_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (cluster_container_instance_cq_id) REFERENCES aws_ecs_cluster_container_instances (cq_id) ON DELETE CASCADE
);