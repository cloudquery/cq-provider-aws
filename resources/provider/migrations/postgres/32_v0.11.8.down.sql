-- Resource: iam.groups
ALTER TABLE IF EXISTS "aws_iam_group_policies" ADD COLUMN IF NOT EXISTS "group_id" text;
DROP TABLE IF EXISTS aws_iam_group_accessed_detail_tracked_actions_last_accessed;
DROP TABLE IF EXISTS aws_iam_group_accessed_detail_entities;
DROP TABLE IF EXISTS aws_iam_group_accessed_details;
