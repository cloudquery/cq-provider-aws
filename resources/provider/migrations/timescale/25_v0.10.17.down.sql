-- Autogenerated by migration tool on 2022-04-12 15:10:48
-- CHANGEME: Verify or edit this file before proceeding

-- Resource: ec2.network_interfaces
DROP TABLE IF EXISTS aws_ec2_network_interface_private_ip_addresses;
DROP TABLE IF EXISTS aws_ec2_network_interfaces;

-- Resource: wafv2.ipsets
DROP TABLE IF EXISTS aws_wafv2_ipsets;

-- Resource: wafv2.regex_pattern_sets
DROP TABLE IF EXISTS aws_wafv2_regex_pattern_sets;

-- Resource: iam.virtual_mfa_devices
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices DROP CONSTRAINT aws_iam_virtual_mfa_devices_pk;
ALTER TABLE IF EXISTS aws_iam_virtual_mfa_devices ADD CONSTRAINT aws_iam_virtual_mfa_devices_pk PRIMARY KEY (cq_fetch_date,serial_number,enable_date);
