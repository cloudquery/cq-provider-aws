-- Resource: ec2.eips

ALTER TABLE IF EXISTS aws_ec2_eips
    DROP CONSTRAINT aws_ec2_eips_pk;
ALTER TABLE IF EXISTS aws_ec2_eips
    ADD CONSTRAINT aws_ec2_eips_pk PRIMARY KEY (account_id,allocation_id);

CREATE TABLE IF NOT EXISTS "aws_access_analyzer_analyzer_archive_rules" (
                                                                            "cq_id" uuid NOT NULL,
                                                                            "cq_meta" jsonb,
                                                                            "analyzer_cq_id" uuid,
                                                                            "created_at" timestamp without time zone,
                                                                            "filter" jsonb,
                                                                            "rule_name" text,
                                                                            "updated_at" timestamp without time zone,
                                                                            CONSTRAINT aws_access_analyzer_analyzer_archive_rules_pk PRIMARY KEY(cq_id),
                                                                            UNIQUE(cq_id),
                                                                            FOREIGN KEY (analyzer_cq_id) REFERENCES aws_access_analyzer_analyzers(cq_id) ON DELETE CASCADE
);