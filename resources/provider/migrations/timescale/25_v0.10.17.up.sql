-- Autogenerated by migration tool on 2022-04-12 15:10:48
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: ec2.network_interfaces
CREATE TABLE IF NOT EXISTS "aws_ec2_network_interfaces" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"arn" text,
	"tags" jsonb,
	"association_allocation_id" text,
	"association_id" text,
	"association_carrier_ip" text,
	"association_customer_owned_ip" text,
	"association_ip_owner_id" text,
	"association_public_dns_name" text,
	"association_public_ip" text,
	"attachment_attach_time" timestamp without time zone,
	"attachment_id" text,
	"attachment_delete_on_termination" boolean,
	"attachment_device_index" integer,
	"attachment_instance_id" text,
	"attachment_instance_owner_id" text,
	"attachment_network_card_index" integer,
	"attachment_status" text,
	"availability_zone" text,
	"deny_all_igw_traffic" boolean,
	"description" text,
	"groups" jsonb,
	"interface_type" text,
	"ipv4_prefixes" text[],
	"ipv6_address" text,
	"ipv6_addresses" text[],
	"ipv6_native" boolean,
	"ipv6_prefixes" text[],
	"mac_address" text,
	"id" text,
	"outpost_arn" text,
	"owner_id" text,
	"private_dns_name" text,
	"private_ip_address" text,
	"requester_id" text,
	"requester_managed" boolean,
	"source_dest_check" boolean,
	"status" text,
	"subnet_id" text,
	"vpc_id" text,
	CONSTRAINT aws_ec2_network_interfaces_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_ec2_network_interfaces');
CREATE TABLE IF NOT EXISTS "aws_ec2_network_interface_private_ip_addresses" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"network_interface_cq_id" uuid,
	"association_allocation_id" text,
	"association_id" text,
	"association_carrier_ip" text,
	"association_customer_owned_ip" text,
	"association_ip_owner_id" text,
	"association_public_dns_name" text,
	"association_public_ip" text,
	"primary" boolean,
	"private_dns_name" text,
	"private_ip_address" text,
	CONSTRAINT aws_ec2_network_interface_private_ip_addresses_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_ec2_network_interface_private_ip_addresses (cq_fetch_date, network_interface_cq_id);
SELECT setup_tsdb_child('aws_ec2_network_interface_private_ip_addresses', 'network_interface_cq_id', 'aws_ec2_network_interfaces', 'cq_id');

-- Resource: wafv2.ipsets
CREATE TABLE IF NOT EXISTS "aws_wafv2_ipsets" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"scope" text,
	"arn" text,
	"addresses" cidr[],
	"ip_address_version" text,
	"id" text,
	"name" text,
	"description" text,
	"tags" jsonb,
	CONSTRAINT aws_wafv2_ipsets_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_wafv2_ipsets');

-- Resource: wafv2.regex_pattern_sets
CREATE TABLE IF NOT EXISTS "aws_wafv2_regex_pattern_sets" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"scope" text,
	"arn" text,
	"description" text,
	"id" text,
	"name" text,
	"regular_expression_list" text[],
	"tags" jsonb,
	CONSTRAINT aws_wafv2_regex_pattern_sets_pk PRIMARY KEY(cq_fetch_date,arn),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_wafv2_regex_pattern_sets');
