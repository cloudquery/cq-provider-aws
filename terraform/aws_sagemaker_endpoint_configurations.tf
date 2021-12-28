resource "aws_sagemaker_endpoint_configuration" "sagemaker_endpoint_configuration" {
  name = "sagemaker-endpoint-configuration-test"

  production_variants {
    variant_name           = "variant-test"
    model_name             = aws_sagemaker_model.sagemaker_model.name
    initial_instance_count = 1
    instance_type          = "ml.t2.medium"
  }

  tags = {
    Name = "sagemaker-endpoint-configuration-test"
  }
}