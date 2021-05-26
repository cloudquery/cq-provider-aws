resource "aws_api_gateway_rest_api" "example" {
  body = jsonencode({
    openapi = "3.0.1"
    info = {
      title = "example"
      version = "1.0"
    }
    paths = {
      "/path1" = {
        get = {
          x-amazon-apigateway-integration = {
            httpMethod = "GET"
            payloadFormatVersion = "1.0"
            type = "HTTP_PROXY"
            uri = "https://ip-ranges.amazonaws.com/ip-ranges.json"
          }
        }
      }
    }
  })

  name = "${var.test_prefix}${var.test_suffix}"

  endpoint_configuration {
    types = [
      "REGIONAL"]
  }

  tags = {
    TestId = var.test_suffix
  }
}


resource "aws_api_gateway_deployment" "example" {
  rest_api_id = aws_api_gateway_rest_api.example.id
  variables = {
    test:"test"
  }
  description = "test description"

  triggers = {
    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.example.body))
  }

  lifecycle {
    create_before_destroy = true
  }
}
