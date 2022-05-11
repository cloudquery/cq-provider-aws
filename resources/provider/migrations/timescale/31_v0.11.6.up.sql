-- Resource: sqs.queues
TRUNCATE TABLE aws_sqs_queues CASCADE;
ALTER TABLE IF EXISTS aws_sqs_queues ALTER COLUMN created_timestamp TYPE timestamp USING to_timestamp(created_timestamp);
ALTER TABLE IF EXISTS aws_sqs_queues ALTER COLUMN last_modified_timestamp TYPE timestamp USING to_timestamp(last_modified_timestamp);

-- Resource: ec2.images
TRUNCATE TABLE aws_ec2_images CASCADE;
ALTER TABLE IF EXISTS aws_ec2_images ALTER COLUMN creation_date TYPE timestamp USING to_timestamp(creation_date, 'YYYY-MM-DDTHH:MM:SS.MS');

-- Resource: ec2.instances
TRUNCATE TABLE aws_ec2_instance_elastic_gpu_associations CASCADE;
ALTER TABLE IF EXISTS aws_ec2_instance_elastic_gpu_associations ALTER COLUMN elastic_gpu_association_time TYPE timestamp USING to_timestamp(elastic_gpu_association_time, 'YYYY-MM-DDTHH:MM:SS.MS');

-- Resource: lambda.layers
TRUNCATE TABLE aws_lambda_layers CASCADE;
ALTER TABLE IF EXISTS aws_lambda_layers ALTER COLUMN latest_matching_version_created_date TYPE timestamp USING to_timestamp(latest_matching_version_created_date, 'YYYY-MM-DDTHH:MM:SS.MSZ');
TRUNCATE TABLE aws_lambda_layer_versions CASCADE;
ALTER TABLE IF EXISTS aws_lambda_layer_versions ALTER COLUMN created_date TYPE timestamp USING to_timestamp(created_date, 'YYYY-MM-DDTHH:MM:SS.MSZ');

-- Resource: lambda.functions
TRUNCATE TABLE aws_lambda_functions CASCADE;
ALTER TABLE IF EXISTS aws_lambda_functions ALTER COLUMN last_modified TYPE timestamp USING to_timestamp(last_modified, 'YYYY-MM-DDTHH:MM:SS.MSZ');
TRUNCATE TABLE aws_lambda_function_versions CASCADE;
ALTER TABLE IF EXISTS aws_lambda_function_versions ALTER COLUMN last_modified TYPE timestamp USING to_timestamp(last_modified, 'YYYY-MM-DDTHH:MM:SS.MSZ');
TRUNCATE TABLE aws_lambda_function_concurrency_configs CASCADE;
ALTER TABLE IF EXISTS aws_lambda_function_concurrency_configs ALTER COLUMN last_modified TYPE timestamp USING to_timestamp(last_modified, 'YYYY-MM-DDTHH:MM:SS.MSZ');