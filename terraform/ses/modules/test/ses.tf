module "ses" {
  source  = "cloudposse/ses/aws"
  version = "0.22.3"
  # insert the 13 required variables here}


  namespace = "${var.prefix}-mq"
  stage     = "test"
  name      = "${var.prefix}-mq"


  enabled = true
  tags    = var.tags
}
