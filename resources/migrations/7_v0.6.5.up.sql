ALTER TABLE IF EXISTS "aws_iam_password_policies" ADD COLUMN policy_exists boolean;

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "directconnect_gateway_cq_id" TO "gateway_cq_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "directconnect_gateway_id" TO "gateway_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_attachments" RENAME COLUMN "directconnect_gateway_id" TO "gateway_id";

ALTER TABLE IF EXISTS "aws_elbv2_listeners" ADD COLUMN "load_balancer_cq_id" uuid;
