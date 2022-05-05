-- aws_redshift_clusters: update ARN to correct values and change PK to arn field
ALTER TABLE IF EXISTS aws_redshift_clusters DROP CONSTRAINT aws_redshift_clusters_pk;
UPDATE aws_redshift_clusters SET arn = format('arn:aws:redshift:%s:%s:cluster:%s', region, account_id, id);
ALTER TABLE IF EXISTS aws_redshift_clusters ADD CONSTRAINT aws_redshift_clusters_pk PRIMARY KEY (arn);

-- aws_redshift_snapshots: add cluster_cq_id and a corresponding FK constraint
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD COLUMN IF NOT EXISTS cluster_cq_id uuid;

UPDATE aws_redshift_snapshots s
SET cluster_cq_id = cl.cq_id
FROM aws_redshift_clusters cl
WHERE cl.id = s.cluster_identifier AND cl.cluster_create_time = s.cluster_create_time;

ALTER TABLE IF EXISTS aws_redshift_snapshots
    ADD CONSTRAINT aws_redshift_snapshots_cluster_cq_id_fkey
    FOREIGN KEY (cluster_cq_id)
    REFERENCES aws_redshift_clusters(cq_id);

-- aws_redshift_snapshots: add ARN, change PK to ARN and fill values
ALTER TABLE IF EXISTS aws_redshift_snapshots DROP CONSTRAINT aws_redshift_snapshots_pk;
ALTER TABLE IF EXISTS aws_redshift_snapshots ADD COLUMN IF NOT EXISTS arn text;

UPDATE aws_redshift_snapshots s
SET arn = format('arn:aws:redshift:%s:%s:snapshot:%s/%s', cl.region, cl.account_id, cl.id, s.snapshot_identifier)
FROM aws_redshift_clusters cl
WHERE cl.cq_id = s.cluster_cq_id;

ALTER TABLE IF EXISTS aws_redshift_snapshots ADD CONSTRAINT aws_redshift_snapshots_pk PRIMARY KEY (arn);

-- aws_redshift_event_subscriptions: add ARN, change PK to ARN and fill values
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions DROP CONSTRAINT aws_redshift_event_subscriptions_pk;
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD COLUMN IF NOT EXISTS arn text;
UPDATE aws_redshift_event_subscriptions SET arn = format('arn:aws:redshift:%s:%s:eventsubscription:%s', region, account_id, id);
ALTER TABLE IF EXISTS aws_redshift_event_subscriptions ADD CONSTRAINT aws_redshift_event_subscriptions_pk PRIMARY KEY (arn);

-- aws_redshift_subnet_groups: change PK to ARN and update values
ALTER TABLE IF EXISTS aws_redshift_subnet_groups DROP CONSTRAINT aws_redshift_subnet_groups_pk;
UPDATE aws_redshift_subnet_groups SET arn = format('arn:aws:redshift:%s:%s:subnetgroup:%s', region, account_id, cluster_subnet_group_name);
ALTER TABLE IF EXISTS aws_redshift_subnet_groups ADD CONSTRAINT aws_redshift_subnet_groups_pk PRIMARY KEY (arn);
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
    "tags" jsonb,
    CONSTRAINT aws_shield_protection_groups_pk PRIMARY KEY (arn),
    UNIQUE (cq_id)
);

-- Resource: shield.subscriptions
CREATE TABLE IF NOT EXISTS "aws_shield_subscriptions" (
    "cq_id" uuid NOT NULL,
    "cq_meta" jsonb,
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
    CONSTRAINT aws_shield_subscriptions_pk PRIMARY KEY (arn),
    UNIQUE (cq_id)
);
