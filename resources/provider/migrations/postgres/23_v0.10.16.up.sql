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

