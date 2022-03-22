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


