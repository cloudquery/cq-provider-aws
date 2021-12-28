resource "aws_dx_gateway" "integ-aws-dx-gateway" {
  name            = "dx-gatewaytest"
  amazon_side_asn = "64512"
}