resource "aws_accessanalyzer_analyzer" "example" {
  analyzer_name = "${var.prefix}-accessanalyzer-test"
  tags          = var.tags
}
