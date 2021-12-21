ALTER TABLE IF EXISTS "aws_sagemaker_model_containers"
    RENAME COLUMN "image_config_repository_auth_config_repository_credentials_provider_arn" TO "image_config_repository_auth_config_repo_creds_provider_arn";


ALTER TABLE IF EXISTS "aws_ec2_security_group_ip_permissions" ADD COLUMN "permission_type" string;



ALTER TABLE public.aws_ec2_security_group_ip_permission_ip_ranges DROP CONSTRAINT aws_ec2_security_group_ip_permission_ip_ranges_pkey


ALTER TABLE IF EXISTS "aws_ec2_security_group_ip_permission_ip_ranges" RENAME COLUMN "cidr_ip" TO "cidr";


ALTER TABLE IF EXISTS "aws_ec2_security_group_ip_permission_ip_ranges" ADD COLUMN "cidr_type" string;


DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permission_ipv6_ranges;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egresses;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_ip_ranges;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_ipv6_ranges;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_prefix_list_ids;
DROP TABLE IF EXISTS public.aws_ec2_security_group_ip_permissions_egress_user_group_pairs;
