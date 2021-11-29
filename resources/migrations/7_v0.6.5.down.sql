ALTER TABLE IF EXISTS "aws_iam_password_policies" DROP COLUMN policy_exists;

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "gateway_cq_id" TO "directconnect_gateway_cq_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_associations" RENAME COLUMN "gateway_id" TO "directconnect_gateway_id";

ALTER TABLE IF EXISTS "aws_directconnect_gateway_attachments" RENAME COLUMN "gateway_id" TO "directconnect_gateway_id";

ALTER TABLE IF EXISTS "aws_elbv2_listeners" DROP COLUMN "load_balancer_cq_id";
