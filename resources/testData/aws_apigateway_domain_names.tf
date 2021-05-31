resource "aws_api_gateway_domain_name" "example" {
  certificate_arn = aws_acm_certificate_validation.example.certificate_arn
  domain_name = "api.${var.test_prefix}${var.test_suffix}.com"
}

resource "aws_route53_zone" "main" {
  name = "${var.test_prefix}${var.test_suffix}.com"
}

resource "aws_route53_record" "www" {
  zone_id = aws_route53_zone.main.zone_id
  name = "api.${var.test_prefix}${var.test_suffix}.com"
  type = "A"
  ttl = "300"
  records = [
    "192.168.1.1"]
}

resource "aws_acm_certificate" "example" {
  domain_name = "example.com"
  validation_method = "DNS"
}

resource "aws_acm_certificate_validation" "example" {
  certificate_arn = aws_acm_certificate.example.arn
  validation_record_fqdns = [for record in aws_route53_record.validation : record.fqdn]
}


resource "aws_route53_record" "validation" {
  for_each = toset(aws_acm_certificate.example.domain_validation_options)
//  for_each = {
//  for dvo in aws_acm_certificate.example.domain_validation_options : dvo.domain_name => {
//    name = dvo.resource_record_name
//    record = dvo.resource_record_value
//    type = dvo.resource_record_type
//  }
//  }


  allow_overwrite = true
  name = each.value.resource_record_name
  records = [
    each.value.resource_record_value]
  ttl = 60
  type = each.value.resource_record_type
  zone_id = aws_route53_zone.main.zone_id
}
