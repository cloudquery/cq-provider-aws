resource "aws_api_gateway_api_key" "aws_apigateway_api_keys_key" {
  name    = "apigw-key-integration-test"
  enabled = true
  value   = "test-key-test-key-test-key-test-key-test-key-test-key"
  tags = {
    test = "test"
  }
}