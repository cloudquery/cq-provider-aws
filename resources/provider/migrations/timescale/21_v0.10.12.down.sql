-- Autogenerated by migration tool on 2022-03-14 15:54:35

-- Resource: ec2.eips
ALTER TABLE IF EXISTS aws_ec2_eips
DROP CONSTRAINT aws_ec2_eips_pk;
ALTER TABLE IF EXISTS aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_pk PRIMARY KEY (cq_fetch_date,account_id,public_ip);

-- Resource: cloudformation.stacks
DROP TABLE IF EXISTS aws_cloudformation_stack_outputs;
DROP TABLE IF EXISTS aws_cloudformation_stack_resources;
DROP TABLE IF EXISTS aws_cloudformation_stacks;
