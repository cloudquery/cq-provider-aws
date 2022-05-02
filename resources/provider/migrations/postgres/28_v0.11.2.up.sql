ALTER TABLE IF EXISTS "aws_sns_topics" ADD COLUMN IF NOT EXISTS "tags" jsonb;

CREATE TABLE IF NOT EXISTS "aws_elasticbeanstalk_application_versions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "region" text,
    "application_name" text,
    "arn" text,
    "build_arn" text,
    "date_created" timestamp WITHOUT TIME ZONE,
    "date_updated" timestamp WITHOUT TIME ZONE,
    "description" text,
    "source_location" text,
    "source_repository" text,
    "source_type" text,
    "source_bundle_s3_bucket" text,
    "source_bundle_s3_key" text,
    "status" text,
    "version_label" text,
    CONSTRAINT aws_elasticbeanstalk_application_versions_pk PRIMARY KEY (arn),
    UNIQUE (cq_id)
);

-- Resource: shield.attacks
CREATE TABLE IF NOT EXISTS "aws_shield_attacks" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "attack_counters" jsonb,
    "id" text,
    "end_time" timestamp WITHOUT TIME ZONE,
    "mitigations" text[],
    "resource_arn" text,
    "start_time" timestamp WITHOUT TIME ZONE,
    CONSTRAINT aws_shield_attacks_pk PRIMARY KEY (id),
    UNIQUE (cq_id)
);
CREATE TABLE IF NOT EXISTS "aws_shield_attack_properties" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "attack_cq_id" uuid,
    "attack_layer" text,
    "attack_property_identifier" text,
    "top_contributors" jsonb,
    "total" bigint,
    "unit" text,
    CONSTRAINT aws_shield_attack_properties_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (attack_cq_id) REFERENCES aws_shield_attacks(cq_id) ON DELETE CASCADE
);
CREATE TABLE IF NOT EXISTS "aws_shield_attack_sub_resources" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "attack_cq_id" uuid,
    "attack_vectors" jsonb,
    "counters" jsonb,
    "id" text,
    "type" text,
    CONSTRAINT aws_shield_attack_sub_resources_pk PRIMARY KEY (cq_id),
    UNIQUE (cq_id),
    FOREIGN KEY (attack_cq_id) REFERENCES aws_shield_attacks(cq_id) ON DELETE CASCADE
);

-- Resource: shield.protections
CREATE TABLE IF NOT EXISTS "aws_shield_protections" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "application_automatic_response_configuration_status" text,
    "health_check_ids" text[],
    "id" text,
    "name" text,
    "arn" text,
    "resource_arn" text,
    CONSTRAINT aws_shield_protections_pk PRIMARY KEY (arn),
    UNIQUE (cq_id)
);

-- Resource: shield.protections_groups
CREATE TABLE IF NOT EXISTS "aws_shield_protection_groups" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "aggregation" text,
    "members" text[],
    "pattern" text,
    "id" text,
    "arn" text,
    "resource_type" text,
    CONSTRAINT aws_shield_protection_groups_pk PRIMARY KEY (arn),
    UNIQUE (cq_id)
);

-- Resource: shield.subscriptions
CREATE TABLE IF NOT EXISTS "aws_shield_subscriptions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "account_id" text,
    "protection_group_limits_max_protection_groups" bigint,
    "protection_group_limits_arbitrary_pattern_limits_max_members" bigint,
    "protected_resource_type_limits" jsonb,
    "auto_renew" text,
    "end_time" timestamp WITHOUT TIME ZONE,
    "limits" jsonb,
    "proactive_engagement_status" text,
    "start_time" timestamp WITHOUT TIME ZONE,
    "arn" text,
    "time_commitment_in_seconds" bigint,
    CONSTRAINT aws_shield_subscriptions_pk PRIMARY KEY (arn),
    UNIQUE (cq_id)
);

