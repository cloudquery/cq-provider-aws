resource "aws_dx_lag" "aws_directconnect_lags_lag" {
  name                  = "tf-dx-${var.test_prefix}-${var.test_suffix}"
  connections_bandwidth = "1Gbps"
  location              = "EqDC2"
  force_destroy         = true
}