variable "test_suffix" {
  type = string
}

variable "test_prefix" {
  type = string
}

locals {
  aws_ecr_url = "${data.aws_caller_identity.current.account_id}.dkr.ecr.${data.aws_region.current.name}.amazonaws.com"
}