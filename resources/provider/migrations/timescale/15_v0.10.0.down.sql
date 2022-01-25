-- Autogenerated by migration tool on 2022-01-25 17:59:57

-- Resource: accessanalyzer.analyzers
DROP TABLE IF EXISTS aws_access_analyzer_analyzer_finding_sources;
DROP TABLE IF EXISTS aws_access_analyzer_analyzer_findings;
DROP TABLE IF EXISTS aws_access_analyzer_analyzers;

-- Resource: acm.certificates
DROP TABLE IF EXISTS aws_acm_certificates;

-- Resource: apigateway.api_keys
DROP TABLE IF EXISTS aws_apigateway_api_keys;

-- Resource: apigateway.client_certificates
DROP TABLE IF EXISTS aws_apigateway_client_certificates;

-- Resource: apigateway.domain_names
DROP TABLE IF EXISTS aws_apigateway_domain_name_base_path_mappings;
DROP TABLE IF EXISTS aws_apigateway_domain_names;

-- Resource: apigateway.rest_apis
DROP TABLE IF EXISTS aws_apigateway_rest_api_authorizers;
DROP TABLE IF EXISTS aws_apigateway_rest_api_deployments;
DROP TABLE IF EXISTS aws_apigateway_rest_api_documentation_parts;
DROP TABLE IF EXISTS aws_apigateway_rest_api_documentation_versions;
DROP TABLE IF EXISTS aws_apigateway_rest_api_gateway_responses;
DROP TABLE IF EXISTS aws_apigateway_rest_api_models;
DROP TABLE IF EXISTS aws_apigateway_rest_api_request_validators;
DROP TABLE IF EXISTS aws_apigateway_rest_api_resources;
DROP TABLE IF EXISTS aws_apigateway_rest_api_stages;
DROP TABLE IF EXISTS aws_apigateway_rest_apis;

-- Resource: apigateway.usage_plans
DROP TABLE IF EXISTS aws_apigateway_usage_plan_api_stages;
DROP TABLE IF EXISTS aws_apigateway_usage_plan_keys;
DROP TABLE IF EXISTS aws_apigateway_usage_plans;

-- Resource: apigateway.vpc_links
DROP TABLE IF EXISTS aws_apigateway_vpc_links;

-- Resource: apigatewayv2.apis
DROP TABLE IF EXISTS aws_apigatewayv2_api_authorizers;
DROP TABLE IF EXISTS aws_apigatewayv2_api_deployments;
DROP TABLE IF EXISTS aws_apigatewayv2_api_integration_responses;
DROP TABLE IF EXISTS aws_apigatewayv2_api_integrations;
DROP TABLE IF EXISTS aws_apigatewayv2_api_models;
DROP TABLE IF EXISTS aws_apigatewayv2_api_route_responses;
DROP TABLE IF EXISTS aws_apigatewayv2_api_routes;
DROP TABLE IF EXISTS aws_apigatewayv2_api_stages;
DROP TABLE IF EXISTS aws_apigatewayv2_apis;

-- Resource: apigatewayv2.domain_names
DROP TABLE IF EXISTS aws_apigatewayv2_domain_name_configurations;
DROP TABLE IF EXISTS aws_apigatewayv2_domain_name_rest_api_mappings;
DROP TABLE IF EXISTS aws_apigatewayv2_domain_names;

-- Resource: apigatewayv2.vpc_links
DROP TABLE IF EXISTS aws_apigatewayv2_vpc_links;

-- Resource: applicationautoscaling.policies
DROP TABLE IF EXISTS aws_applicationautoscaling_policies;

-- Resource: autoscaling.groups
DROP TABLE IF EXISTS aws_autoscaling_group_instances;
DROP TABLE IF EXISTS aws_autoscaling_group_tags;
DROP TABLE IF EXISTS aws_autoscaling_group_scaling_policies;
DROP TABLE IF EXISTS aws_autoscaling_group_lifecycle_hooks;
DROP TABLE IF EXISTS aws_autoscaling_groups;

-- Resource: autoscaling.launch_configurations
DROP TABLE IF EXISTS aws_autoscaling_launch_configuration_block_device_mappings;
DROP TABLE IF EXISTS aws_autoscaling_launch_configurations;

-- Resource: aws.regions
DROP TABLE IF EXISTS aws_regions;

-- Resource: cloudfront.cache_policies
DROP TABLE IF EXISTS aws_cloudfront_cache_policies;

-- Resource: cloudfront.distributions
DROP TABLE IF EXISTS aws_cloudfront_distribution_default_cache_behavior_functions;
DROP TABLE IF EXISTS aws_cloudfront_distribution_origins;
DROP TABLE IF EXISTS aws_cloudfront_distribution_cache_behavior_lambda_functions;
DROP TABLE IF EXISTS aws_cloudfront_distribution_cache_behaviors;
DROP TABLE IF EXISTS aws_cloudfront_distribution_custom_error_responses;
DROP TABLE IF EXISTS aws_cloudfront_distribution_origin_groups;
DROP TABLE IF EXISTS aws_cloudfront_distributions;

-- Resource: cloudtrail.trails
DROP TABLE IF EXISTS aws_cloudtrail_trail_event_selectors;
DROP TABLE IF EXISTS aws_cloudtrail_trails;

-- Resource: cloudwatch.alarms
DROP TABLE IF EXISTS aws_cloudwatch_alarm_metrics;
DROP TABLE IF EXISTS aws_cloudwatch_alarms;

-- Resource: cloudwatchlogs.filters
DROP TABLE IF EXISTS aws_cloudwatchlogs_filter_metric_transformations;
DROP TABLE IF EXISTS aws_cloudwatchlogs_filters;

-- Resource: codebuild.projects
DROP TABLE IF EXISTS aws_codebuild_project_environment_variables;
DROP TABLE IF EXISTS aws_codebuild_project_file_system_locations;
DROP TABLE IF EXISTS aws_codebuild_project_secondary_artifacts;
DROP TABLE IF EXISTS aws_codebuild_project_secondary_sources;
DROP TABLE IF EXISTS aws_codebuild_projects;

-- Resource: cognito.identity_pools
DROP TABLE IF EXISTS aws_cognito_identity_pool_cognito_identity_providers;
DROP TABLE IF EXISTS aws_cognito_identity_pools;

-- Resource: cognito.user_pools
DROP TABLE IF EXISTS aws_cognito_user_pool_schema_attributes;
DROP TABLE IF EXISTS aws_cognito_user_pool_identity_providers;
DROP TABLE IF EXISTS aws_cognito_user_pools;

-- Resource: config.configuration_recorders
DROP TABLE IF EXISTS aws_config_configuration_recorders;

-- Resource: config.conformance_packs
DROP TABLE IF EXISTS aws_config_conformance_packs;

-- Resource: dax.clusters
DROP TABLE IF EXISTS aws_dax_cluster_nodes;
DROP TABLE IF EXISTS aws_dax_clusters;

-- Resource: directconnect.connections
DROP TABLE IF EXISTS aws_directconnect_connection_mac_sec_keys;
DROP TABLE IF EXISTS aws_directconnect_connections;

-- Resource: directconnect.gateways
DROP TABLE IF EXISTS aws_directconnect_gateway_associations;
DROP TABLE IF EXISTS aws_directconnect_gateway_attachments;
DROP TABLE IF EXISTS aws_directconnect_gateways;

-- Resource: directconnect.lags
DROP TABLE IF EXISTS aws_directconnect_lag_mac_sec_keys;
DROP TABLE IF EXISTS aws_directconnect_lags;

-- Resource: directconnect.virtual_gateways
DROP TABLE IF EXISTS aws_directconnect_virtual_gateways;

-- Resource: directconnect.virtual_interfaces
DROP TABLE IF EXISTS aws_directconnect_virtual_interface_bgp_peers;
DROP TABLE IF EXISTS aws_directconnect_virtual_interfaces;

-- Resource: dms.replication_instances
DROP TABLE IF EXISTS aws_dms_replication_instance_replication_subnet_group_subnets;
DROP TABLE IF EXISTS aws_dms_replication_instance_vpc_security_groups;
DROP TABLE IF EXISTS aws_dms_replication_instances;

-- Resource: dynamodb.tables
DROP TABLE IF EXISTS aws_dynamodb_table_global_secondary_indexes;
DROP TABLE IF EXISTS aws_dynamodb_table_local_secondary_indexes;
DROP TABLE IF EXISTS aws_dynamodb_table_replicas;
DROP TABLE IF EXISTS aws_dynamodb_table_replica_auto_scalings;
DROP TABLE IF EXISTS aws_dynamodb_table_continuous_backups;
DROP TABLE IF EXISTS aws_dynamodb_tables;

-- Resource: ec2.byoip_cidrs
DROP TABLE IF EXISTS aws_ec2_byoip_cidrs;

-- Resource: ec2.customer_gateways
DROP TABLE IF EXISTS aws_ec2_customer_gateways;

-- Resource: ec2.ebs_snapshots
DROP TABLE IF EXISTS aws_ec2_ebs_snapshots;

-- Resource: ec2.ebs_volumes
DROP TABLE IF EXISTS aws_ec2_ebs_volume_attachments;
DROP TABLE IF EXISTS aws_ec2_ebs_volumes;

-- Resource: ec2.eips
DROP TABLE IF EXISTS aws_ec2_eips;

-- Resource: ec2.flow_logs
DROP TABLE IF EXISTS aws_ec2_flow_logs;

-- Resource: ec2.images
DROP TABLE IF EXISTS aws_ec2_image_block_device_mappings;
DROP TABLE IF EXISTS aws_ec2_images;

-- Resource: ec2.instances
DROP TABLE IF EXISTS aws_ec2_instance_block_device_mappings;
DROP TABLE IF EXISTS aws_ec2_instance_elastic_gpu_associations;
DROP TABLE IF EXISTS aws_ec2_instance_elastic_inference_accelerator_associations;
DROP TABLE IF EXISTS aws_ec2_instance_network_interface_groups;
DROP TABLE IF EXISTS aws_ec2_instance_network_interface_ipv6_addresses;
DROP TABLE IF EXISTS aws_ec2_instance_network_interface_private_ip_addresses;
DROP TABLE IF EXISTS aws_ec2_instance_network_interfaces;
DROP TABLE IF EXISTS aws_ec2_instance_product_codes;
DROP TABLE IF EXISTS aws_ec2_instance_security_groups;
DROP TABLE IF EXISTS aws_ec2_instances;

-- Resource: ec2.internet_gateways
DROP TABLE IF EXISTS aws_ec2_internet_gateway_attachments;
DROP TABLE IF EXISTS aws_ec2_internet_gateways;

-- Resource: ec2.nat_gateways
DROP TABLE IF EXISTS aws_ec2_nat_gateway_addresses;
DROP TABLE IF EXISTS aws_ec2_nat_gateways;

-- Resource: ec2.network_acls
DROP TABLE IF EXISTS aws_ec2_network_acl_associations;
DROP TABLE IF EXISTS aws_ec2_network_acl_entries;
DROP TABLE IF EXISTS aws_ec2_network_acls;

-- Resource: ec2.regional_config
DROP TABLE IF EXISTS aws_ec2_regional_config;

-- Resource: ec2.route_tables
DROP TABLE IF EXISTS aws_ec2_route_table_associations;
DROP TABLE IF EXISTS aws_ec2_route_table_propagating_vgws;
DROP TABLE IF EXISTS aws_ec2_route_table_routes;
DROP TABLE IF EXISTS aws_ec2_route_tables;

-- Resource: ec2.security_groups
DROP TABLE IF EXISTS aws_ec2_security_group_ip_permission_ip_ranges;
DROP TABLE IF EXISTS aws_ec2_security_group_ip_permission_prefix_list_ids;
DROP TABLE IF EXISTS aws_ec2_security_group_ip_permission_user_id_group_pairs;
DROP TABLE IF EXISTS aws_ec2_security_group_ip_permissions;
DROP TABLE IF EXISTS aws_ec2_security_groups;

-- Resource: ec2.subnets
DROP TABLE IF EXISTS aws_ec2_subnet_ipv6_cidr_block_association_sets;
DROP TABLE IF EXISTS aws_ec2_subnets;

-- Resource: ec2.transit_gateways
DROP TABLE IF EXISTS aws_ec2_transit_gateway_attachments;
DROP TABLE IF EXISTS aws_ec2_transit_gateway_route_tables;
DROP TABLE IF EXISTS aws_ec2_transit_gateway_vpc_attachments;
DROP TABLE IF EXISTS aws_ec2_transit_gateway_peering_attachments;
DROP TABLE IF EXISTS aws_ec2_transit_gateway_multicast_domains;
DROP TABLE IF EXISTS aws_ec2_transit_gateways;

-- Resource: ec2.vpc_endpoints
DROP TABLE IF EXISTS aws_ec2_vpc_endpoint_dns_entries;
DROP TABLE IF EXISTS aws_ec2_vpc_endpoint_groups;
DROP TABLE IF EXISTS aws_ec2_vpc_endpoints;

-- Resource: ec2.vpc_peering_connections
DROP TABLE IF EXISTS aws_ec2_vpc_peering_connections;

-- Resource: ec2.vpcs
DROP TABLE IF EXISTS aws_ec2_vpc_cidr_block_association_sets;
DROP TABLE IF EXISTS aws_ec2_vpc_ipv6_cidr_block_association_sets;
DROP TABLE IF EXISTS aws_ec2_vpcs;

-- Resource: ec2.vpn_gateways
DROP TABLE IF EXISTS aws_ec2_vpc_attachment;
DROP TABLE IF EXISTS aws_ec2_vpn_gateways;

-- Resource: ecr.repositories
DROP TABLE IF EXISTS aws_ecr_repository_images;
DROP TABLE IF EXISTS aws_ecr_repositories;

-- Resource: ecs.clusters
DROP TABLE IF EXISTS aws_ecs_cluster_attachments;
DROP TABLE IF EXISTS aws_ecs_cluster_service_deployments;
DROP TABLE IF EXISTS aws_ecs_cluster_service_events;
DROP TABLE IF EXISTS aws_ecs_cluster_service_load_balancers;
DROP TABLE IF EXISTS aws_ecs_cluster_service_service_registries;
DROP TABLE IF EXISTS aws_ecs_cluster_service_task_set_load_balancers;
DROP TABLE IF EXISTS aws_ecs_cluster_service_task_set_service_registries;
DROP TABLE IF EXISTS aws_ecs_cluster_service_task_sets;
DROP TABLE IF EXISTS aws_ecs_cluster_services;
DROP TABLE IF EXISTS aws_ecs_cluster_container_instance_attachments;
DROP TABLE IF EXISTS aws_ecs_cluster_container_instance_attributes;
DROP TABLE IF EXISTS aws_ecs_cluster_container_instance_registered_resources;
DROP TABLE IF EXISTS aws_ecs_cluster_container_instance_remaining_resources;
DROP TABLE IF EXISTS aws_ecs_cluster_container_instances;
DROP TABLE IF EXISTS aws_ecs_clusters;

-- Resource: ecs.task_definitions
DROP TABLE IF EXISTS aws_ecs_task_definition_container_definitions;
DROP TABLE IF EXISTS aws_ecs_task_definition_volumes;
DROP TABLE IF EXISTS aws_ecs_task_definitions;

-- Resource: efs.filesystems
DROP TABLE IF EXISTS aws_efs_filesystems;

-- Resource: eks.clusters
DROP TABLE IF EXISTS aws_eks_cluster_encryption_configs;
DROP TABLE IF EXISTS aws_eks_cluster_loggings;
DROP TABLE IF EXISTS aws_eks_clusters;

-- Resource: elasticbeanstalk.applications
DROP TABLE IF EXISTS aws_elasticbeanstalk_applications;

-- Resource: elasticbeanstalk.environments
DROP TABLE IF EXISTS aws_elasticbeanstalk_configuration_setting_options;
DROP TABLE IF EXISTS aws_elasticbeanstalk_configuration_settings;
DROP TABLE IF EXISTS aws_elasticbeanstalk_configuration_options;
DROP TABLE IF EXISTS aws_elasticbeanstalk_environment_links;
DROP TABLE IF EXISTS aws_elasticbeanstalk_environments;

-- Resource: elasticsearch.domains
DROP TABLE IF EXISTS aws_elasticsearch_domains;

-- Resource: elbv1.load_balancers
DROP TABLE IF EXISTS aws_elbv1_load_balancer_backend_server_descriptions;
DROP TABLE IF EXISTS aws_elbv1_load_balancer_listeners;
DROP TABLE IF EXISTS aws_elbv1_load_balancer_policies_app_cookie_stickiness;
DROP TABLE IF EXISTS aws_elbv1_load_balancer_policies_lb_cookie_stickiness;
DROP TABLE IF EXISTS aws_elbv1_load_balancer_policies;
DROP TABLE IF EXISTS aws_elbv1_load_balancers;

-- Resource: elbv2.load_balancers
DROP TABLE IF EXISTS aws_elbv2_listener_certificates;
DROP TABLE IF EXISTS aws_elbv2_listener_default_action_forward_config_target_groups;
DROP TABLE IF EXISTS aws_elbv2_listener_default_actions;
DROP TABLE IF EXISTS aws_elbv2_listeners;
DROP TABLE IF EXISTS aws_elbv2_load_balancer_availability_zone_addresses;
DROP TABLE IF EXISTS aws_elbv2_load_balancer_availability_zones;
DROP TABLE IF EXISTS aws_elbv2_load_balancer_attributes;
DROP TABLE IF EXISTS aws_elbv2_load_balancers;

-- Resource: elbv2.target_groups
DROP TABLE IF EXISTS aws_elbv2_target_groups;

-- Resource: emr.block_public_access_configs
DROP TABLE IF EXISTS aws_emr_block_public_access_config_port_ranges;
DROP TABLE IF EXISTS aws_emr_block_public_access_configs;

-- Resource: emr.clusters
DROP TABLE IF EXISTS aws_emr_clusters;

-- Resource: fsx.backups
DROP TABLE IF EXISTS aws_fsx_backups;

-- Resource: guardduty.detectors
DROP TABLE IF EXISTS aws_guardduty_detector_members;
DROP TABLE IF EXISTS aws_guardduty_detectors;

-- Resource: iam.accounts
DROP TABLE IF EXISTS aws_accounts;

-- Resource: iam.groups
DROP TABLE IF EXISTS aws_iam_group_policies;
DROP TABLE IF EXISTS aws_iam_groups;

-- Resource: iam.openid_connect_identity_providers
DROP TABLE IF EXISTS aws_iam_openid_connect_identity_providers;

-- Resource: iam.password_policies
DROP TABLE IF EXISTS aws_iam_password_policies;

-- Resource: iam.policies
DROP TABLE IF EXISTS aws_iam_policy_versions;
DROP TABLE IF EXISTS aws_iam_policies;

-- Resource: iam.roles
DROP TABLE IF EXISTS aws_iam_role_policies;
DROP TABLE IF EXISTS aws_iam_roles;

-- Resource: iam.saml_identity_providers
DROP TABLE IF EXISTS aws_iam_saml_identity_providers;

-- Resource: iam.server_certificates
DROP TABLE IF EXISTS aws_iam_server_certificates;

-- Resource: iam.users
DROP TABLE IF EXISTS aws_iam_user_access_keys;
DROP TABLE IF EXISTS aws_iam_user_groups;
DROP TABLE IF EXISTS aws_iam_user_attached_policies;
DROP TABLE IF EXISTS aws_iam_user_policies;
DROP TABLE IF EXISTS aws_iam_users;

-- Resource: iam.virtual_mfa_devices
DROP TABLE IF EXISTS aws_iam_virtual_mfa_devices;

-- Resource: iot.billing_groups
DROP TABLE IF EXISTS aws_iot_billing_groups;

-- Resource: iot.ca_certificates
DROP TABLE IF EXISTS aws_iot_ca_certificates;

-- Resource: iot.certificates
DROP TABLE IF EXISTS aws_iot_certificates;

-- Resource: iot.policies
DROP TABLE IF EXISTS aws_iot_policies;

-- Resource: iot.streams
DROP TABLE IF EXISTS aws_iot_stream_files;
DROP TABLE IF EXISTS aws_iot_streams;

-- Resource: iot.thing_groups
DROP TABLE IF EXISTS aws_iot_thing_groups;

-- Resource: iot.thing_types
DROP TABLE IF EXISTS aws_iot_thing_types;

-- Resource: iot.things
DROP TABLE IF EXISTS aws_iot_things;

-- Resource: iot.topic_rules
DROP TABLE IF EXISTS aws_iot_topic_rule_actions;
DROP TABLE IF EXISTS aws_iot_topic_rules;

-- Resource: kms.keys
DROP TABLE IF EXISTS aws_kms_keys;

-- Resource: lambda.functions
DROP TABLE IF EXISTS aws_lambda_function_file_system_configs;
DROP TABLE IF EXISTS aws_lambda_function_layers;
DROP TABLE IF EXISTS aws_lambda_function_aliases;
DROP TABLE IF EXISTS aws_lambda_function_event_invoke_configs;
DROP TABLE IF EXISTS aws_lambda_function_version_file_system_configs;
DROP TABLE IF EXISTS aws_lambda_function_version_layers;
DROP TABLE IF EXISTS aws_lambda_function_versions;
DROP TABLE IF EXISTS aws_lambda_function_concurrency_configs;
DROP TABLE IF EXISTS aws_lambda_function_event_source_mappings;
DROP TABLE IF EXISTS aws_lambda_functions;

-- Resource: lambda.layers
DROP TABLE IF EXISTS aws_lambda_layer_version_policies;
DROP TABLE IF EXISTS aws_lambda_layer_versions;
DROP TABLE IF EXISTS aws_lambda_layers;

-- Resource: lambda.runtimes
DROP TABLE IF EXISTS aws_lambda_runtimes;

-- Resource: mq.brokers
DROP TABLE IF EXISTS aws_mq_broker_configurations;
DROP TABLE IF EXISTS aws_mq_broker_users;
DROP TABLE IF EXISTS aws_mq_brokers;

-- Resource: organizations.accounts
DROP TABLE IF EXISTS aws_organizations_accounts;

-- Resource: rds.certificates
DROP TABLE IF EXISTS aws_rds_certificates;

-- Resource: rds.cluster_parameter_groups
DROP TABLE IF EXISTS aws_rds_cluster_parameters;
DROP TABLE IF EXISTS aws_rds_cluster_parameter_groups;

-- Resource: rds.cluster_snapshots
DROP TABLE IF EXISTS aws_rds_cluster_snapshots;

-- Resource: rds.clusters
DROP TABLE IF EXISTS aws_rds_cluster_associated_roles;
DROP TABLE IF EXISTS aws_rds_cluster_db_cluster_members;
DROP TABLE IF EXISTS aws_rds_cluster_domain_memberships;
DROP TABLE IF EXISTS aws_rds_cluster_vpc_security_groups;
DROP TABLE IF EXISTS aws_rds_clusters;

-- Resource: rds.db_parameter_groups
DROP TABLE IF EXISTS aws_rds_db_parameters;
DROP TABLE IF EXISTS aws_rds_db_parameter_groups;

-- Resource: rds.db_security_groups
DROP TABLE IF EXISTS aws_rds_db_security_groups;

-- Resource: rds.db_snapshots
DROP TABLE IF EXISTS aws_rds_db_snapshots;

-- Resource: rds.db_subnet_groups
DROP TABLE IF EXISTS aws_rds_subnet_group_subnets;
DROP TABLE IF EXISTS aws_rds_subnet_groups;

-- Resource: rds.event_subscriptions
DROP TABLE IF EXISTS aws_rds_event_subscriptions;

-- Resource: rds.instances
DROP TABLE IF EXISTS aws_rds_instance_associated_roles;
DROP TABLE IF EXISTS aws_rds_instance_db_instance_automated_backups_replications;
DROP TABLE IF EXISTS aws_rds_instance_db_parameter_groups;
DROP TABLE IF EXISTS aws_rds_instance_db_security_groups;
DROP TABLE IF EXISTS aws_rds_instance_db_subnet_group_subnets;
DROP TABLE IF EXISTS aws_rds_instance_domain_memberships;
DROP TABLE IF EXISTS aws_rds_instance_option_group_memberships;
DROP TABLE IF EXISTS aws_rds_instance_vpc_security_groups;
DROP TABLE IF EXISTS aws_rds_instances;

-- Resource: redshift.clusters
DROP TABLE IF EXISTS aws_redshift_cluster_nodes;
DROP TABLE IF EXISTS aws_redshift_cluster_parameters;
DROP TABLE IF EXISTS aws_redshift_cluster_parameter_group_status_lists;
DROP TABLE IF EXISTS aws_redshift_cluster_parameter_groups;
DROP TABLE IF EXISTS aws_redshift_cluster_security_groups;
DROP TABLE IF EXISTS aws_redshift_cluster_deferred_maintenance_windows;
DROP TABLE IF EXISTS aws_redshift_cluster_endpoint_vpc_endpoint_network_interfaces;
DROP TABLE IF EXISTS aws_redshift_cluster_endpoint_vpc_endpoints;
DROP TABLE IF EXISTS aws_redshift_cluster_iam_roles;
DROP TABLE IF EXISTS aws_redshift_cluster_vpc_security_groups;
DROP TABLE IF EXISTS aws_redshift_clusters;

-- Resource: redshift.subnet_groups
DROP TABLE IF EXISTS aws_redshift_subnet_group_subnets;
DROP TABLE IF EXISTS aws_redshift_subnet_groups;

-- Resource: route53.domains
DROP TABLE IF EXISTS aws_route53_domain_nameservers;
DROP TABLE IF EXISTS aws_route53_domains;

-- Resource: route53.health_checks
DROP TABLE IF EXISTS aws_route53_health_checks;

-- Resource: route53.hosted_zones
DROP TABLE IF EXISTS aws_route53_hosted_zone_query_logging_configs;
DROP TABLE IF EXISTS aws_route53_hosted_zone_resource_record_sets;
DROP TABLE IF EXISTS aws_route53_hosted_zone_traffic_policy_instances;
DROP TABLE IF EXISTS aws_route53_hosted_zone_vpc_association_authorizations;
DROP TABLE IF EXISTS aws_route53_hosted_zones;

-- Resource: route53.reusable_delegation_sets
DROP TABLE IF EXISTS aws_route53_reusable_delegation_sets;

-- Resource: route53.traffic_policies
DROP TABLE IF EXISTS aws_route53_traffic_policy_versions;
DROP TABLE IF EXISTS aws_route53_traffic_policies;

-- Resource: s3.accounts
DROP TABLE IF EXISTS aws_s3_account_config;

-- Resource: s3.buckets
DROP TABLE IF EXISTS aws_s3_bucket_grants;
DROP TABLE IF EXISTS aws_s3_bucket_cors_rules;
DROP TABLE IF EXISTS aws_s3_bucket_encryption_rules;
DROP TABLE IF EXISTS aws_s3_bucket_replication_rules;
DROP TABLE IF EXISTS aws_s3_bucket_lifecycles;
DROP TABLE IF EXISTS aws_s3_buckets;

-- Resource: sagemaker.endpoint_configurations
DROP TABLE IF EXISTS aws_sagemaker_endpoint_configuration_production_variants;
DROP TABLE IF EXISTS aws_sagemaker_endpoint_configurations;

-- Resource: sagemaker.models
DROP TABLE IF EXISTS aws_sagemaker_model_containers;
DROP TABLE IF EXISTS aws_sagemaker_model_vpc_config;
DROP TABLE IF EXISTS aws_sagemaker_models;

-- Resource: sagemaker.notebook_instances
DROP TABLE IF EXISTS aws_sagemaker_notebook_instances;

-- Resource: sagemaker.training_jobs
DROP TABLE IF EXISTS aws_sagemaker_training_job_algorithm_specification;
DROP TABLE IF EXISTS aws_sagemaker_training_job_debug_hook_config;
DROP TABLE IF EXISTS aws_sagemaker_training_job_debug_rule_configurations;
DROP TABLE IF EXISTS aws_sagemaker_training_job_debug_rule_evaluation_statuses;
DROP TABLE IF EXISTS aws_sagemaker_training_job_input_data_config;
DROP TABLE IF EXISTS aws_sagemaker_training_job_profiler_rule_configurations;
DROP TABLE IF EXISTS aws_sagemaker_training_job_profiler_rule_evaluation_statuses;
DROP TABLE IF EXISTS aws_sagemaker_training_jobs;

-- Resource: secretsmanager.secrets
DROP TABLE IF EXISTS aws_secretsmanager_secrets;

-- Resource: sns.subscriptions
DROP TABLE IF EXISTS aws_sns_subscriptions;

-- Resource: sns.topics
DROP TABLE IF EXISTS aws_sns_topics;

-- Resource: sqs.queues
DROP TABLE IF EXISTS aws_sqs_queues;

-- Resource: ssm.documents
DROP TABLE IF EXISTS aws_ssm_documents;

-- Resource: ssm.instances
DROP TABLE IF EXISTS aws_ssm_instance_compliance_items;
DROP TABLE IF EXISTS aws_ssm_instances;

-- Resource: waf.rule_groups
DROP TABLE IF EXISTS aws_waf_rule_groups;

-- Resource: waf.rules
DROP TABLE IF EXISTS aws_waf_rule_predicates;
DROP TABLE IF EXISTS aws_waf_rules;

-- Resource: waf.subscribed_rule_groups
DROP TABLE IF EXISTS aws_waf_subscribed_rule_groups;

-- Resource: waf.web_acls
DROP TABLE IF EXISTS aws_waf_web_acl_rules;
DROP TABLE IF EXISTS aws_waf_web_acls;

-- Resource: wafv2.managed_rule_groups
DROP TABLE IF EXISTS aws_wafv2_managed_rule_groups;

-- Resource: wafv2.rule_groups
DROP TABLE IF EXISTS aws_wafv2_rule_groups;

-- Resource: wafv2.web_acls
DROP TABLE IF EXISTS aws_wafv2_web_acl_rules;
DROP TABLE IF EXISTS aws_wafv2_web_acl_post_process_firewall_manager_rule_groups;
DROP TABLE IF EXISTS aws_wafv2_web_acl_pre_process_firewall_manager_rule_groups;
DROP TABLE IF EXISTS aws_wafv2_web_acls;
