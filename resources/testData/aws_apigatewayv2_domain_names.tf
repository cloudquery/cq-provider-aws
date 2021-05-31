// domain names
//**********************************
resource "aws_apigatewayv2_domain_name" "dn" {
  domain_name = "${var.test_prefix}${var.test_suffix}.com"

  domain_name_configuration {
    certificate_arn = aws_acm_certificate.example.arn
    endpoint_type = "REGIONAL"
    security_policy = "TLS_1_2"
  }

  depends_on = [
    aws_acm_certificate.example]
}

resource "tls_self_signed_cert" "example" {
  key_algorithm = "RSA"
  private_key_pem = file("private_key2.pem")

  subject {
    common_name = "*.example.com"
    organization = "ACME Examples, Inc"
  }

  validity_period_hours = 12

  allowed_uses = [
    "key_encipherment",
    "digital_signature",
    "server_auth",
  ]
}

resource "aws_acm_certificate" "example" {
  private_key = tls_self_signed_cert.example.private_key_pem
  certificate_body = tls_self_signed_cert.example.cert_pem
}

resource "aws_apigatewayv2_api_mapping" "dn" {
  api_id = aws_apigatewayv2_api.dn.id
  domain_name = aws_apigatewayv2_domain_name.dn.id
  stage = aws_apigatewayv2_stage.dn.id
}

resource "aws_apigatewayv2_api" "dn" {
  name = "v2dn${var.test_prefix}${var.test_suffix}"
  protocol_type = "HTTP"
  route_selection_expression = "$request.body.action"
}


resource "aws_apigatewayv2_stage" "dn" {
  api_id = aws_apigatewayv2_api.dn.id
  name = "v2stage${var.test_prefix}${var.test_suffix}"
}