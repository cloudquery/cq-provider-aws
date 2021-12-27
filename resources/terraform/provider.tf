provider "aws" {
  region = "us-east-1"

  default_tags {
    tags = {
      Type   = "integration_test"
    }
  }
}

data "aws_region" "current" {}