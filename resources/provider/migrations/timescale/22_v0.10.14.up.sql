-- Resource: ec2.instance_statuses
CREATE TABLE IF NOT EXISTS "aws_ec2_instance_statuses"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "account_id"            text,
    "region"                text,
    "arn"                   text,
    "availability_zone"     text,
    "instance_id"           text,
    "instance_state_code"   integer,
    "instance_state_name"   text,
    "details"               jsonb,
    "status"                text,
    "outpost_arn"           text,
    "system_status"         text,
    "system_status_details" jsonb,
    CONSTRAINT aws_ec2_instance_statuses_pk PRIMARY KEY (cq_fetch_date, arn),
    UNIQUE (cq_fetch_date, cq_id)
);
SELECT setup_tsdb_parent('aws_ec2_instance_statuses');
CREATE TABLE IF NOT EXISTS "aws_ec2_instance_status_events"
(
    "cq_id"                 uuid                        NOT NULL,
    "cq_meta"               jsonb,
    "cq_fetch_date"         timestamp without time zone NOT NULL,
    "instance_status_cq_id" uuid,
    "code"                  text,
    "description"           text,
    "id"                    text,
    "not_after"             timestamp without time zone,
    "not_before"            timestamp without time zone,
    "not_before_deadline"   timestamp without time zone,
    CONSTRAINT aws_ec2_instance_status_events_pk PRIMARY KEY (cq_fetch_date, instance_status_cq_id, id),
    UNIQUE (cq_fetch_date, cq_id)
);
CREATE INDEX ON aws_ec2_instance_status_events (cq_fetch_date, instance_status_cq_id);
SELECT setup_tsdb_child('aws_ec2_instance_status_events', 'instance_status_cq_id', 'aws_ec2_instance_statuses',
                        'cq_id');

-- Resource: ec2.nat_gateways
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses
    DROP CONSTRAINT aws_ec2_nat_gateway_addresses_pk;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses
    ADD CONSTRAINT aws_ec2_nat_gateway_addresses_pk PRIMARY KEY (cq_fetch_date, nat_gateway_cq_id, network_interface_id);
ALTER TABLE aws_ec2_nat_gateway_addresses
    ALTER COLUMN allocation_id DROP NOT NULL;

-- Resource: mq.brokers
ALTER TABLE IF EXISTS aws_mq_brokers
    DROP CONSTRAINT aws_mq_brokers_pk;
ALTER TABLE IF EXISTS "aws_mq_brokers"
    RENAME COLUMN broker_id TO id;
ALTER TABLE IF EXISTS aws_mq_brokers
    ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (cq_fetch_date, account_id, id);
ALTER TABLE IF EXISTS aws_mq_broker_configurations
    DROP CONSTRAINT aws_mq_broker_configurations_pk;
ALTER TABLE IF EXISTS "aws_mq_broker_configurations"
    DROP COLUMN IF EXISTS "account_id";
ALTER TABLE IF EXISTS "aws_mq_broker_configurations"
    DROP COLUMN IF EXISTS "region";
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
ALTER TABLE IF EXISTS "aws_mq_broker_users"
    DROP COLUMN IF EXISTS "account_id";
ALTER TABLE IF EXISTS "aws_mq_broker_users"
    DROP COLUMN IF EXISTS "region";
ALTER TABLE IF EXISTS aws_mq_broker_users
    ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (cq_fetch_date, cq_id);

