
# Table: aws_eks_clusters

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|account_id|text||
|region|text||
|arn|text||
|certificate_authority_data|text||
|client_request_token|text||
|created_at|timestamp without time zone||
|endpoint|text||
|identity_oidc_issuer|text||
|kubernetes_network_config_service_ipv4_cidr|text||
|name|text||
|platform_version|text||
|resources_vpc_config_cluster_security_group_id|text||
|resources_vpc_config_endpoint_private_access|boolean||
|resources_vpc_config_endpoint_public_access|boolean||
|resources_vpc_config_public_access_cidrs|text[]||
|resources_vpc_config_security_group_ids|text[]||
|resources_vpc_config_subnet_ids|text[]||
|resources_vpc_config_vpc_id|text||
|role_arn|text||
|status|text||
|tags|jsonb||
|version|text||
## Relations
## Table: aws_eks_cluster_encryption_configs

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|provider_key_arn|text||
|resources|text[]||
## Table: aws_eks_cluster_logging_cluster_loggings

## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|cluster_id|uuid||
|enabled|boolean||
|types|text[]||
