resource "aws_cloudwatch_log_metric_filter" "aws_cloudwatch_log_metric_filter" {
  name           = "aws_cloudwatch_log_metric_filter_"
  pattern        = ""
  log_group_name = aws_cloudwatch_log_group.aws_cloudwatch_log_metric_filter_group.name

  metric_transformation {
    name      = "aws_cloudwatch_log_metric_filter_name"
    namespace = "YourNamespace"
    value     = "1"
  }
}

resource "aws_cloudwatch_log_group" "aws_cloudwatch_log_metric_filter_group" {
  name = "MyApp/access.log"
}