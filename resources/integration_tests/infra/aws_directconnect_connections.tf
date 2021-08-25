resource "aws_dx_connection" "aws_directconnect_connections_connection" {
  name      = "tf-dx-connection${var.test_prefix}-${var.test_suffix}"
  bandwidth = "1Gbps"
  location  = "EqDC2"
}