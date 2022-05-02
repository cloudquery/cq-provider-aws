DROP TABLE IF EXISTS aws_elasticbeanstalk_application_versions;

ALTER TABLE IF EXISTS "aws_sns_topics" DROP COLUMN IF EXISTS "tags";

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