resource "aws_route53_health_check" "route53_health_check" {
  fqdn              = "test.com"
  port              = 80
  type              = "HTTP"
  resource_path     = "/"
  failure_threshold = "5"
  request_interval  = "10"

  tags = {
    Name = "health-checktest"
  }
}
