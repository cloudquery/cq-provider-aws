resource "aws_api_gateway_vpc_link" "apigateway_vpc_link" {
  name        = "apigw-vpc-link-${var.test_prefix}-${var.test_suffix}"
  description = "example description"
  target_arns = [
  aws_lb.apigateway_nlb.arn]
}

resource "aws_lb" "apigateway_nlb" {
  name               = "apigateway-nlb-${var.test_suffix}"
  internal           = false
  load_balancer_type = "network"
  subnets = [
    aws_subnet.aws_vpc_subnet2.id,
  aws_subnet.aws_vpc_subnet3.id]

  enable_deletion_protection = false

  tags = {
    Environment = "dev"
  }
}
//
//resource "aws_api_gateway_rest_api" "apigateway_api" {
//  body = jsonencode({
//    openapi = "3.0.1"
//    info = {
//      title = "example"
//      version = "1.0"
//    }
//    paths = {
//      "/path1" = {
//        get = {
//          x-amazon-apigateway-integration = {
//            httpMethod = "GET"
//            payloadFormatVersion = "1.0"
//            type = "HTTP_PROXY"
//            uri = "https://ip-ranges.amazonaws.com/ip-ranges.json"
//          }
//        }
//      }
//    }
//  })
//
//
//
//  name = "a-${var.test_prefix}-${var.test_suffix}"
//
//  endpoint_configuration {
//    types = [
//      "REGIONAL"]
//  }
//}
//
//resource "aws_api_gateway_deployment" "apigateway_deploymnet" {
//  rest_api_id = aws_api_gateway_rest_api.apigateway_api.id
//
//  triggers = {
//    redeployment = sha1(jsonencode(aws_api_gateway_rest_api.apigateway_api.body))
//  }
//
//  lifecycle {
//    create_before_destroy = true
//  }
//}
//
//resource "aws_api_gateway_stage" "apigateway_stage" {
//  deployment_id = aws_api_gateway_deployment.apigateway_deploymnet.id
//  rest_api_id = aws_api_gateway_rest_api.apigateway_api.id
//  stage_name = "test"
//}
//
//resource "aws_api_gateway_resource" "apigateway_resource" {
//  rest_api_id = aws_api_gateway_rest_api.apigateway_api.id
//  parent_id = aws_api_gateway_rest_api.apigateway_api.root_resource_id
//  path_part = "test"
//}
//
//resource "aws_api_gateway_method" "apigateway_method" {
//  rest_api_id = aws_api_gateway_rest_api.apigateway_api.id
//  resource_id = aws_api_gateway_resource.apigateway_resource.id
//  http_method = "GET"
//  authorization = "NONE"
//}
//
//resource "aws_api_gateway_integration" "apigateway_integration" {
//  rest_api_id = aws_api_gateway_rest_api.apigateway_api.id
//  resource_id = aws_api_gateway_resource.apigateway_resource.id
//  http_method = aws_api_gateway_method.apigateway_method.http_method
//  type = "MOCK"
//  // cache_key_parameters = ["method.request.path.param"]
//  cache_namespace = "test"
//  timeout_milliseconds = 29000
//
//  request_parameters = {
//    "integration.request.header.X-Authorization" = "'static'"
//  }
//
//  # Transforms the incoming XML request to JSON
//  request_templates = {
//    "application/xml" = <<EOF
//{
//   "body" : $input.json('$')
//}
//EOF
//  }
//}
//
//resource "aws_api_gateway_authorizer" "apigateway_authorizer" {
//  name = "a-${var.test_prefix}-${var.test_suffix}"
//  rest_api_id = aws_api_gateway_rest_api.apigateway_api.id
//  authorizer_uri = aws_lambda_function.apigateway_auth_function.invoke_arn
//  authorizer_credentials = aws_iam_role.invocation_role.arn
//}
//
//
//resource "aws_iam_role" "invocation_role" {
//  name = "ir-${var.test_prefix}-${var.test_suffix}"
//  path = "/"
//
//  assume_role_policy = <<EOF
//{
//  "Version": "2012-10-17",
//  "Statement": [
//    {
//      "Action": "sts:AssumeRole",
//      "Principal": {
//        "Service": "apigateway.amazonaws.com"
//      },
//      "Effect": "Allow",
//      "Sid": ""
//    }
//  ]
//}
//EOF
//}
//
//resource "aws_iam_role_policy" "invocation_policy" {
//  name = "ip-${var.test_prefix}-${var.test_suffix}"
//  role = aws_iam_role.invocation_role.id
//
//  policy = <<EOF
//{
//  "Version": "2012-10-17",
//  "Statement": [
//    {
//      "Action": "lambda:InvokeFunction",
//      "Effect": "Allow",
//      "Resource": "${aws_lambda_function.apigateway_auth_function.arn}"
//    }
//  ]
//}
//EOF
//}
//
//resource "aws_iam_role" "apigateway_lambda_role" {
//  name = "l-${var.test_prefix}-${var.test_suffix}"
//
//  assume_role_policy = <<EOF
//{
//  "Version": "2012-10-17",
//  "Statement": [
//    {
//      "Action": "sts:AssumeRole",
//      "Principal": {
//        "Service": "lambda.amazonaws.com"
//      },
//      "Effect": "Allow",
//      "Sid": ""
//    }
//  ]
//}
//EOF
//}
//
//resource "aws_lambda_function" "apigateway_auth_function" {
//  filename = data.archive_file.apigateway_zip_file.output_path
//  source_code_hash = data.archive_file.apigateway_zip_file.output_base64sha256
//  function_name = "function_${var.test_prefix}${var.test_suffix}"
//  role = aws_iam_role.apigateway_lambda_role.arn
//  handler = "exports.example"
//  runtime = "nodejs12.x"
//  publish = true
//
//  environment {
//    variables = {
//      foo = "bar"
//    }
//  }
//}
//
//
//data "archive_file" "apigateway_zip_file" {
//  type = "zip"
//  output_path = "./tmp/${var.test_prefix}${var.test_suffix}.zip"
//  source {
//    content = <<EOF
//module.exports.handler = async (event, context, callback) => {
//	const what = "world";
//	const response = `Hello $${what}!`;
//	callback(null, response);
//};
//EOF
//    filename = "main.js"
//  }
//}
