-- Autogenerated by migration tool on 2022-04-04 17:13:07
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: codepipeline.pipelines
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipelines" (
                                                            "cq_id" uuid NOT NULL,
                                                            "cq_meta" jsonb,
                                                            "cq_fetch_date" timestamp without time zone NOT NULL,
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
                                                            CONSTRAINT aws_codepipeline_pipelines_pk PRIMARY KEY(cq_fetch_date,arn),
                                                            UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_codepipeline_pipelines');
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipeline_stages" (
                                                                  "cq_id" uuid NOT NULL,
                                                                  "cq_meta" jsonb,
                                                                  "cq_fetch_date" timestamp without time zone NOT NULL,
                                                                  "pipeline_cq_id" uuid,
                                                                  "name" text,
                                                                  "blockers" jsonb,
                                                                  CONSTRAINT aws_codepipeline_pipeline_stages_pk PRIMARY KEY(cq_fetch_date,cq_id),
                                                                  UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_codepipeline_pipeline_stages (cq_fetch_date, pipeline_cq_id);
SELECT setup_tsdb_child('aws_codepipeline_pipeline_stages', 'pipeline_cq_id', 'aws_codepipeline_pipelines', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_codepipeline_pipeline_stage_actions" (
                                                                         "cq_id" uuid NOT NULL,
                                                                         "cq_meta" jsonb,
                                                                         "cq_fetch_date" timestamp without time zone NOT NULL,
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
                                                                         CONSTRAINT aws_codepipeline_pipeline_stage_actions_pk PRIMARY KEY(cq_fetch_date,cq_id),
                                                                         UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_codepipeline_pipeline_stage_actions (cq_fetch_date, pipeline_stage_cq_id);
SELECT setup_tsdb_child('aws_codepipeline_pipeline_stage_actions', 'pipeline_stage_cq_id', 'aws_codepipeline_pipeline_stages', 'cq_id');

-- Resource: sns.subscriptions
ALTER TABLE IF EXISTS aws_sns_subscriptions DROP CONSTRAINT aws_sns_subscriptions_pk;
ALTER TABLE IF EXISTS aws_sns_subscriptions ADD CONSTRAINT aws_sns_subscriptions_pk PRIMARY KEY (cq_fetch_date,endpoint,owner,protocol,arn,topic_arn);

-- Resource: mq.brokers
ALTER TABLE IF EXISTS aws_mq_brokers
    DROP CONSTRAINT aws_mq_brokers_pk;
ALTER TABLE IF EXISTS "aws_mq_brokers"
    RENAME COLUMN broker_id TO id;
ALTER TABLE IF EXISTS aws_mq_brokers
    ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (cq_fetch_date, account_id, id);
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    DROP CONSTRAINT aws_mq_broker_configurations_pk;
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    ADD CONSTRAINT aws_mq_broker_configurations_pk PRIMARY KEY (cq_fetch_date, cq_id);
CREATE TABLE IF NOT EXISTS "aws_mq_broker_configuration_revisions"
(
    "cq_id"                      uuid                        NOT NULL,
    "cq_meta"                    jsonb,
    "cq_fetch_date"              timestamp without time zone NOT NULL,
    "broker_configuration_cq_id" uuid,
    "configuration_id"           text,
    "created"                    timestamp without time zone,
    "data"                       jsonb,
    "description"                text,
    CONSTRAINT aws_mq_broker_configuration_revisions_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_mq_broker_configuration_revisions (cq_fetch_date, broker_configuration_cq_id);
SELECT setup_tsdb_child('aws_mq_broker_configuration_revisions', 'broker_configuration_cq_id',
                        'aws_mq_broker_configurations', 'cq_id');
ALTER TABLE IF EXISTS aws_mq_broker_users
    DROP CONSTRAINT aws_mq_broker_users_pk;
ALTER TABLE IF EXISTS aws_mq_broker_users
    ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (cq_fetch_date, cq_id);


