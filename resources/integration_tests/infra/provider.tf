

provider "aws" {
  region = "us-east-1"

  default_tags {
    tags = {
      TestId = var.test_suffix
      Type   = "integration_test"
    }
  }
}

data "aws_region" "current" {}
data aws_caller_identity "current" {}
data aws_ecr_authorization_token "token" {}

provider "docker" {
  registry_auth {
    address  = local.aws_ecr_url
    username = data.aws_ecr_authorization_token.token.user_name
    password = data.aws_ecr_authorization_token.token.password
  }
}