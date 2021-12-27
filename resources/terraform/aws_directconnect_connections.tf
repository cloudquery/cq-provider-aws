resource "aws_dx_connection" "aws_directconnect_connections_connection" {
  name      = "dx-connection"
  bandwidth = "1Gbps"
  location  = "EqDC2"
}