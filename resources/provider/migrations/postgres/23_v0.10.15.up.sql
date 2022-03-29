CREATE TABLE IF NOT EXISTS "aws_access_analyzer_analyzer_archive_rules"
(
    "cq_id"          uuid NOT NULL,
    "cq_meta"        jsonb,
    "analyzer_cq_id" uuid,
    "created_at"     timestamp without time zone,
    "filter"         jsonb,
    "rule_name"      text,
    "updated_at"     timestamp without time zone,
    CONSTRAINT aws_access_analyzer_analyzer_archive_rules_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (analyzer_cq_id) REFERENCES aws_access_analyzer_analyzers (cq_id) ON DELETE CASCADE
);

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