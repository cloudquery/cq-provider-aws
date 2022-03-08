-- Resource: ec2.eips
TRUNCATE aws_ec2_eips CASCADE;

ALTER TABLE IF EXISTS aws_ec2_eips
    DROP CONSTRAINT aws_ec2_eips_pk;
ALTER TABLE IF EXISTS aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_pk PRIMARY KEY (account_id,allocation_id);

