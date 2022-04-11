-- Resource: backup.plans
DROP TABLE IF EXISTS aws_backup_plan_rules;
DROP TABLE IF EXISTS aws_backup_plan_selections;
DROP TABLE IF EXISTS aws_backup_plans;

-- Resource: backup.vaults
DROP TABLE IF EXISTS aws_backup_vault_recovery_points;
DROP TABLE IF EXISTS aws_backup_vaults;

-- Resource: codepipeline.webhooks
DROP TABLE IF EXISTS aws_codepipeline_webhook_filters;
DROP TABLE IF EXISTS aws_codepipeline_webhooks;