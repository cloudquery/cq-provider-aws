resource "aws_dx_gateway" "aws_directconnect_virtual_gateways_gateway" {
  name            = "vg-g-${var.test_prefix}-${var.test_suffix}"
  amazon_side_asn = "64512"
}

resource "aws_vpc" "aws_directconnect_virtual_gateways_vpc" {
  cidr_block = "10.255.255.0/28"
}

resource "aws_vpn_gateway" "example" {
  vpc_id = aws_vpc.aws_directconnect_virtual_gateways_vpc.id
}

resource "aws_dx_gateway_association" "aws_directconnect_virtual_gateways_association" {
  dx_gateway_id         = aws_dx_gateway.aws_directconnect_virtual_gateways_gateway.id
  associated_gateway_id = aws_vpn_gateway.example.id

  allowed_prefixes = [
    "210.52.109.0/24",
    "175.45.176.0/22",
  ]

}