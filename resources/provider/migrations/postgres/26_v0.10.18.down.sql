-- Resource: ecs.task_definitions
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "ephemeral_storage_size";
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "runtime_platform_cpu_architecture";
ALTER TABLE IF EXISTS "aws_ecs_task_definitions" DROP COLUMN IF EXISTS "runtime_platform_os_family";


-- Resource: shield.attacks
DROP TABLE IF EXISTS aws_shield_attack_properties;
DROP TABLE IF EXISTS aws_shield_attack_sub_resources;
DROP TABLE IF EXISTS aws_shield_attacks;

-- Resource: shield.protections
DROP TABLE IF EXISTS aws_shield_protections;

-- Resource: shield.protections_groups
DROP TABLE IF EXISTS aws_shield_protection_groups;

-- Resource: shield.subscriptions
DROP TABLE IF EXISTS aws_shield_subscriptions;