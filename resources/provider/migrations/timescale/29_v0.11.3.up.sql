-- Resource: shield.attacks
CREATE TABLE IF NOT EXISTS "aws_shield_attacks" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "account_id" text,
    "attack_counters" jsonb,
    "id" text,
    "end_time" timestamp WITHOUT TIME ZONE,
    "mitigations" text[],
    "resource_arn" text,
    "start_time" timestamp WITHOUT TIME ZONE,
    CONSTRAINT aws_shield_attacks_pk PRIMARY KEY (cq_fetch_date, id),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_attacks');
CREATE TABLE IF NOT EXISTS "aws_shield_attack_properties" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "attack_cq_id" uuid,
    "attack_layer" text,
    "attack_property_identifier" text,
    "top_contributors" jsonb,
    "total" bigint,
    "unit" text,
    CONSTRAINT aws_shield_attack_properties_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_shield_attack_properties(cq_fetch_date, attack_cq_id);
SELECT setup_tsdb_child('aws_shield_attack_properties', 'attack_cq_id', 'aws_shield_attacks', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_shield_attack_sub_resources" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "attack_cq_id" uuid,
    "attack_vectors" jsonb,
    "counters" jsonb,
    "id" text,
    "type" text,
    CONSTRAINT aws_shield_attack_sub_resources_pk PRIMARY KEY (cq_fetch_date, cq_id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_shield_attack_sub_resources(cq_fetch_date, attack_cq_id);
SELECT setup_tsdb_child('aws_shield_attack_sub_resources', 'attack_cq_id', 'aws_shield_attacks', 'cq_id');

-- Resource: shield.protections
CREATE TABLE IF NOT EXISTS "aws_shield_protections" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "account_id" text,
    "region" text,
    "tags" jsonb,
    "application_automatic_response_configuration_status" text,
    "health_check_ids" text[],
    "id" text,
    "name" text,
    "arn" text,
    "resource_arn" text,
    CONSTRAINT aws_shield_protections_pk PRIMARY KEY (cq_fetch_date, arn),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_protections');

-- Resource: shield.protections_groups
CREATE TABLE IF NOT EXISTS "aws_shield_protection_groups" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "account_id" text,
    "aggregation" text,
    "members" text[],
    "pattern" text,
    "id" text,
    "arn" text,
    "resource_type" text,
    CONSTRAINT aws_shield_protection_groups_pk PRIMARY KEY (cq_fetch_date, arn),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_protection_groups');

-- Resource: shield.subscriptions
CREATE TABLE IF NOT EXISTS "aws_shield_subscriptions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
    "cq_fetch_date" timestamp WITHOUT TIME ZONE NOT NULL,
    "account_id" text,
    "protection_group_limits_max_protection_groups" int,
    "protection_group_limits_arbitrary_pattern_limits_max_members" int,
    "protected_resource_type_limits" jsonb,
    "auto_renew" text,
    "end_time" timestamp WITHOUT TIME ZONE,
    "limits" jsonb,
    "proactive_engagement_status" text,
    "start_time" timestamp WITHOUT TIME ZONE,
    "arn" text,
    "time_commitment_in_seconds" int,
    CONSTRAINT aws_shield_subscriptions_pk PRIMARY KEY (cq_fetch_date, arn),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_shield_subscriptions');
