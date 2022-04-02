CREATE TABLE IF NOT EXISTS "aws_access_analyzer_analyzer_archive_rules"
(
    "cq_id"          uuid                        NOT NULL,
    "cq_meta"        jsonb,
    "cq_fetch_date"  timestamp without time zone NOT NULL,
    "analyzer_cq_id" uuid,
    "created_at"     timestamp without time zone,
    "filter"         jsonb,
    "rule_name"      text,
    "updated_at"     timestamp without time zone,
    CONSTRAINT aws_access_analyzer_analyzer_archive_rules_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_access_analyzer_analyzer_archive_rules (cq_fetch_date, analyzer_cq_id);
SELECT setup_tsdb_child('aws_access_analyzer_analyzer_archive_rules', 'analyzer_cq_id', 'aws_access_analyzer_analyzers',
                        'cq_id');

-- Resource: ec2.images
ALTER TABLE IF EXISTS "aws_ec2_images" ADD COLUMN IF NOT EXISTS "last_launched_time" timestamp without time zone;

-- Resource: ec2.security_groups
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs DROP CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk;
ALTER TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs ADD CONSTRAINT aws_ec2_security_group_ip_permission_user_id_group_pairs_pk PRIMARY KEY (cq_fetch_date,cq_id);


-- Resource: ecs.clusters
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_tasks"
(
    "cq_id"                         uuid                        NOT NULL,
    "cq_meta"                       jsonb,
    "cq_fetch_date"                 timestamp without time zone NOT NULL,
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
    CONSTRAINT aws_ecs_cluster_tasks_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_tasks (cq_fetch_date, cluster_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_tasks', 'cluster_cq_id', 'aws_ecs_clusters', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_task_attachments"
(
    "cq_id"              uuid                        NOT NULL,
    "cq_meta"            jsonb,
    "cq_fetch_date"      timestamp without time zone NOT NULL,
    "cluster_task_cq_id" uuid,
    "details"            jsonb,
    "id"                 text,
    "status"             text,
    "type"               text,
    CONSTRAINT aws_ecs_cluster_task_attachments_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_task_attachments (cq_fetch_date, cluster_task_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_task_attachments', 'cluster_task_cq_id', 'aws_ecs_cluster_tasks', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_task_containers"
(
    "cq_id"              uuid                        NOT NULL,
    "cq_meta"            jsonb,
    "cq_fetch_date"      timestamp without time zone NOT NULL,
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
    CONSTRAINT aws_ecs_cluster_task_containers_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_task_containers (cq_fetch_date, cluster_task_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_task_containers', 'cluster_task_cq_id', 'aws_ecs_cluster_tasks', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_services"
(
    "cq_id"                                                        uuid                        NOT NULL,
    "cq_meta"                                                      jsonb,
    "cq_fetch_date"                                                timestamp without time zone NOT NULL,
    "cluster_cq_id"                                                uuid,
    "capacity_provider_strategy"                                   jsonb,
    "cluster_arn"                                                  text,
    "created_at"                                                   timestamp without time zone,
    "created_by"                                                   text,
    "deployment_configuration_deployment_circuit_breaker_enable"   boolean,
    "deployment_configuration_deployment_circuit_breaker_rollback" boolean,
    "deployment_configuration_maximum_percent"                     integer,
    "deployment_configuration_minimum_healthy_percent"             integer,
    "deployment_controller_type"                                   text,
    "desired_count"                                                integer,
    "enable_ecs_managed_tags"                                      boolean,
    "enable_execute_command"                                       boolean,
    "health_check_grace_period_seconds"                            integer,
    "launch_type"                                                  text,
    "network_configuration_awsvpc_configuration_subnets"           text[],
    "network_configuration_awsvpc_configuration_assign_public_ip"  text,
    "network_configuration_awsvpc_configuration_security_groups"   text[],
    "pending_count"                                                integer,
    "placement_constraints"                                        jsonb,
    "placement_strategy"                                           jsonb,
    "platform_family"                                              text,
    "platform_version"                                             text,
    "propagate_tags"                                               text,
    "role_arn"                                                     text,
    "running_count"                                                integer,
    "scheduling_strategy"                                          text,
    "arn"                                                          text,
    "name"                                                         text,
    "status"                                                       text,
    "tags"                                                         jsonb,
    "task_definition"                                              text,
    CONSTRAINT aws_ecs_cluster_services_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_services (cq_fetch_date, cluster_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_services', 'cluster_cq_id', 'aws_ecs_clusters', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_deployments"
(
    "cq_id"                                                       uuid                        NOT NULL,
    "cq_meta"                                                     jsonb,
    "cq_fetch_date"                                               timestamp without time zone NOT NULL,
    "cluster_service_cq_id"                                       uuid,
    "capacity_provider_strategy"                                  jsonb,
    "created_at"                                                  timestamp without time zone,
    "desired_count"                                               integer,
    "failed_tasks"                                                integer,
    "id"                                                          text,
    "launch_type"                                                 text,
    "network_configuration_awsvpc_configuration_subnets"          text[],
    "network_configuration_awsvpc_configuration_assign_public_ip" text,
    "network_configuration_awsvpc_configuration_security_groups"  text[],
    "pending_count"                                               integer,
    "platform_family"                                             text,
    "platform_version"                                            text,
    "rollout_state"                                               text,
    "rollout_state_reason"                                        text,
    "running_count"                                               integer,
    "status"                                                      text,
    "task_definition"                                             text,
    "updated_at"                                                  timestamp without time zone,
    CONSTRAINT aws_ecs_cluster_service_deployments_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_deployments (cq_fetch_date, cluster_service_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_deployments', 'cluster_service_cq_id', 'aws_ecs_cluster_services',
                        'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_events"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "cluster_service_cq_id" uuid,
    "created_at"            timestamp without time zone,
    "id"                    text,
    "message"               text,
    CONSTRAINT aws_ecs_cluster_service_events_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_events (cq_fetch_date, cluster_service_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_events', 'cluster_service_cq_id', 'aws_ecs_cluster_services', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_load_balancers"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "cluster_service_cq_id" uuid,
    "container_name"        text,
    "container_port"        integer,
    "load_balancer_name"    text,
    "target_group_arn"      text,
    CONSTRAINT aws_ecs_cluster_service_load_balancers_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_load_balancers (cq_fetch_date, cluster_service_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_load_balancers', 'cluster_service_cq_id', 'aws_ecs_cluster_services',
                        'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_service_registries"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "cluster_service_cq_id" uuid,
    "container_name"        text,
    "container_port"        integer,
    "port"                  integer,
    "registry_arn"          text,
    CONSTRAINT aws_ecs_cluster_service_service_registries_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_service_registries (cq_fetch_date, cluster_service_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_service_registries', 'cluster_service_cq_id',
                        'aws_ecs_cluster_services', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_task_sets"
(
    "cq_id"                                                       uuid                        NOT NULL,
    "cq_meta"                                                     jsonb,
    "cq_fetch_date"                                               timestamp without time zone NOT NULL,
    "cluster_service_cq_id"                                       uuid,
    "capacity_provider_strategy"                                  jsonb,
    "cluster_arn"                                                 text,
    "computed_desired_count"                                      integer,
    "created_at"                                                  timestamp without time zone,
    "external_id"                                                 text,
    "id"                                                          text,
    "launch_type"                                                 text,
    "network_configuration_awsvpc_configuration_subnets"          text[],
    "network_configuration_awsvpc_configuration_assign_public_ip" text,
    "network_configuration_awsvpc_configuration_security_groups"  text[],
    "pending_count"                                               integer,
    "platform_family"                                             text,
    "platform_version"                                            text,
    "running_count"                                               integer,
    "scale_unit"                                                  text,
    "scale_value"                                                 float,
    "service_arn"                                                 text,
    "stability_status"                                            text,
    "stability_status_at"                                         timestamp without time zone,
    "started_by"                                                  text,
    "status"                                                      text,
    "tags"                                                        jsonb,
    "task_definition"                                             text,
    "arn"                                                         text,
    "updated_at"                                                  timestamp without time zone,
    CONSTRAINT aws_ecs_cluster_service_task_sets_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_task_sets (cq_fetch_date, cluster_service_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_task_sets', 'cluster_service_cq_id', 'aws_ecs_cluster_services',
                        'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_task_set_load_balancers"
(
    "cq_id"                          uuid                        NOT NULL,
    "cq_meta"                        jsonb,
    "cq_fetch_date"                  timestamp without time zone NOT NULL,
    "cluster_service_task_set_cq_id" uuid,
    "container_name"                 text,
    "container_port"                 integer,
    "load_balancer_name"             text,
    "target_group_arn"               text,
    CONSTRAINT aws_ecs_cluster_service_task_set_load_balancers_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_task_set_load_balancers (cq_fetch_date, cluster_service_task_set_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_task_set_load_balancers', 'cluster_service_task_set_cq_id',
                        'aws_ecs_cluster_service_task_sets', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_ecs_cluster_service_task_set_service_registries"
(
    "cq_id"                          uuid                        NOT NULL,
    "cq_meta"                        jsonb,
    "cq_fetch_date"                  timestamp without time zone NOT NULL,
    "cluster_service_task_set_cq_id" uuid,
    "container_name"                 text,
    "container_port"                 integer,
    "port"                           integer,
    "arn"                            text,
    CONSTRAINT aws_ecs_cluster_service_task_set_service_registries_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_service_task_set_service_registries (cq_fetch_date, cluster_service_task_set_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_service_task_set_service_registries', 'cluster_service_task_set_cq_id',
                        'aws_ecs_cluster_service_task_sets', 'cq_id');
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
    "cq_id"                            uuid                        NOT NULL,
    "cq_meta"                          jsonb,
    "cq_fetch_date"                    timestamp without time zone NOT NULL,
    "cluster_container_instance_cq_id" uuid,
    "last_status_change"               timestamp without time zone,
    "last_updated"                     timestamp without time zone,
    "status"                           text,
    "type"                             text,
    CONSTRAINT aws_ecs_cluster_container_instance_health_status_details_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ecs_cluster_container_instance_health_status_details (cq_fetch_date, cluster_container_instance_cq_id);
SELECT setup_tsdb_child('aws_ecs_cluster_container_instance_health_status_details', 'cluster_container_instance_cq_id',
                        'aws_ecs_cluster_container_instances', 'cq_id');
