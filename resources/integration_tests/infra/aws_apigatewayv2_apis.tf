resource "aws_apigatewayv2_api" "v2_api_1" {
  name = "${var.test_prefix}v2api${var.test_suffix}"
  protocol_type = "WEBSOCKET"
  route_selection_expression = "$request.body.action"
}

resource "aws_apigatewayv2_integration" "v2_integration_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  integration_type = "HTTP_PROXY"

  integration_method = "ANY"
  integration_uri = "https://example.com/{proxy}"
}

resource "aws_apigatewayv2_model" "v2_model_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  content_type = "application/json"
  name = "${var.test_prefix}v2model${var.test_suffix}"

  schema = <<EOF
{
  "$schema": "http://json-schema.org/draft-04/schema#",
  "title": "ExampleModel",
  "type": "object",
  "properties": {
    "id": { "type": "string" }
  }
}
EOF
}

resource "aws_apigatewayv2_stage" "v2_stage_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  name = "${var.test_prefix}v2stage${var.test_suffix}"
}

resource "aws_apigatewayv2_route" "v2_route_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  route_key = "GET /example/v1/test"

  target = "integrations/${aws_apigatewayv2_integration.v2_integration_1.id}"
}

resource "aws_apigatewayv2_route_response" "v2_route_response" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  route_id = aws_apigatewayv2_route.v2_route_1.id
  route_response_key = "$default"
}

resource "aws_apigatewayv2_deployment" "v2_deployment_1" {
  api_id = aws_apigatewayv2_route.v2_route_1.api_id
  description = "Example deployment"

  lifecycle {
    create_before_destroy = true
  }
}

resource "aws_apigatewayv2_authorizer" "v2_authorizer_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  authorizer_type = "REQUEST"
  authorizer_uri = aws_lambda_function.authorizer_v2.invoke_arn
  identity_sources = [
    "route.request.header.Auth"]
  name = "example-authorizer"
}


resource "aws_iam_role" "v2_iam_role_1" {
  name = "apiv2${aws_apigatewayv2_integration.v2_integration_1.id}"
  path = "/"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "apigateway.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_iam_role_policy" "v2_role_policy" {
  name = "v2apipolicy${aws_apigatewayv2_integration.v2_integration_1.id}"
  role = aws_iam_role.v2_iam_role_1.id

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "lambda:InvokeFunction",
      "Effect": "Allow",
      "Resource": "${aws_lambda_function.authorizer_v2.arn}"
    }
  ]
}
EOF
}

resource "aws_iam_role" "v2_iam_role_2" {
  name = "v2api_lambda_role${aws_apigatewayv2_integration.v2_integration_1.id}"

  assume_role_policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": "sts:AssumeRole",
      "Principal": {
        "Service": "lambda.amazonaws.com"
      },
      "Effect": "Allow",
      "Sid": ""
    }
  ]
}
EOF
}

resource "aws_lambda_function" "authorizer_v2" {
  filename = data.archive_file.lambda_zip_inline.output_path
  source_code_hash = data.archive_file.lambda_zip_inline.output_base64sha256
  function_name = "v2authorizer${var.test_prefix}${var.test_suffix}"
  role = aws_iam_role.v2_iam_role_2.arn
  handler = "exports.example"
  runtime = "nodejs12.x"
}

resource "aws_apigatewayv2_integration_response" "v2_response_1" {
  api_id = aws_apigatewayv2_api.v2_api_1.id
  integration_id = aws_apigatewayv2_integration.v2_integration_1.id
  integration_response_key = "/200/"
}


data "archive_file" "lambda_zip_inline" {
  type = "zip"
  output_path = "./tmp/lambda_zip_inline.zip"
  source {
    content = <<EOF
module.exports.handler = async (event, context, callback) => {
	const what = "world";
	const response = `Hello $${what}!`;
	callback(null, response);
};
EOF
    filename = "main.js"
  }
}
