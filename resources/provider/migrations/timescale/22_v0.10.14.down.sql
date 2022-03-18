-- Resource: ec2.instance_statuses
DROP TABLE IF EXISTS aws_ec2_instance_status_events;
DROP TABLE IF EXISTS aws_ec2_instance_statuses;

-- Resource: ec2.nat_gateways
ALTER TABLE aws_ec2_nat_gateway_addresses ALTER COLUMN allocation_id SET NOT NULL;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses DROP CONSTRAINT aws_ec2_nat_gateway_addresses_pk;
ALTER TABLE IF EXISTS aws_ec2_nat_gateway_addresses ADD CONSTRAINT aws_ec2_nat_gateway_addresses_pk PRIMARY KEY (cq_fetch_date,nat_gateway_cq_id,allocation_id,network_interface_id);

-- Resource: mq.brokers
ALTER TABLE IF EXISTS aws_mq_brokers DROP CONSTRAINT aws_mq_brokers_pk;
ALTER TABLE IF EXISTS aws_mq_brokers ADD CONSTRAINT aws_mq_brokers_pk PRIMARY KEY (account_id,broker_id);
ALTER TABLE IF EXISTS aws_mq_broker_configurations DROP CONSTRAINT aws_mq_broker_configurations_pk;
ALTER TABLE IF EXISTS aws_mq_broker_configurations ADD CONSTRAINT aws_mq_broker_configurations_pk PRIMARY KEY (broker_cq_id,id);
DROP TABLE IF EXISTS aws_mq_broker_configuration_revisions;
ALTER TABLE IF EXISTS "aws_mq_broker_configurations" ADD COLUMN IF NOT EXISTS "account_id" text;
ALTER TABLE IF EXISTS "aws_mq_broker_configurations" ADD COLUMN IF NOT EXISTS "region" text;
ALTER TABLE IF EXISTS aws_mq_broker_users DROP CONSTRAINT aws_mq_broker_users_pk;
ALTER TABLE IF EXISTS aws_mq_broker_users ADD CONSTRAINT aws_mq_broker_users_pk PRIMARY KEY (broker_cq_id,username);
ALTER TABLE IF EXISTS "aws_mq_broker_users" ADD COLUMN IF NOT EXISTS "account_id" text;
ALTER TABLE IF EXISTS "aws_mq_broker_users" ADD COLUMN IF NOT EXISTS "region" text;
ALTER TABLE IF EXISTS "aws_mq_brokers" RENAME COLUMN id TO broker_id;
