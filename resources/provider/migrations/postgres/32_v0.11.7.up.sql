-- Resource: elbv2.target_groups
CREATE TABLE IF NOT EXISTS "aws_elbv2_target_group_target_health_descriptions" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"target_group_cq_id" uuid,
	"health_check_port" text,
	"target_id" text,
	"target_availability_zone" text,
	"target_port" integer,
	"target_health_description" text,
	"target_health_reason" text,
	"target_health_state" text,
	CONSTRAINT aws_elbv2_target_group_target_health_descriptions_pk PRIMARY KEY(cq_id),
	UNIQUE(cq_id),
	FOREIGN KEY (target_group_cq_id) REFERENCES aws_elbv2_target_groups(cq_id) ON DELETE CASCADE
);
