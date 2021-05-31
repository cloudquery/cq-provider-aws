resource "aws_api_gateway_domain_name" "adn" {
  certificate_arn = aws_acm_certificate_validation.adn.certificate_arn
  domain_name = "api.${var.test_prefix}${var.test_suffix}.com"
}

resource "aws_route53_zone" "adn" {
  name = "${var.test_prefix}${var.test_suffix}.com"
}

resource "aws_acm_certificate" "adn" {
  domain_name = "${var.test_prefix}${var.test_suffix}.com"
  validation_method = "DNS"
}

resource "aws_acm_certificate_validation" "adn" {
  certificate_arn = aws_acm_certificate.adn.arn
  validation_record_fqdns = [for record in aws_route53_record.adn : record.fqdn]
}

resource "aws_route53_record" "adn" {
  for_each = toset(aws_acm_certificate.adn.domain_validation_options)
  allow_overwrite = true
  name = each.value.resource_record_name
  records = [
    each.value.resource_record_value]
  ttl = 60
  type = each.value.resource_record_type
  zone_id = aws_route53_zone.adn.zone_id

  depends_on = [aws_acm_certificate.adn]
}
