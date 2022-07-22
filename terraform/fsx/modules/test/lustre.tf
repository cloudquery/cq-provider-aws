module "vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "${var.prefix}-fsx"
  cidr = "10.0.0.0/16"

  azs                    = ["us-east-1a", "us-east-1b"]
  private_subnets        = ["10.0.1.0/24", "10.0.2.0/24"]
  enable_ipv6            = true
  enable_nat_gateway     = false
  create_egress_only_igw = true
}

module "security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "${var.prefix}-fsx"
  description = "${var.prefix}-fsx security group"
  vpc_id      = module.vpc.vpc_id

  # ingress
  ingress_with_cidr_blocks = [
    {
      from_port   = 988
      to_port     = 988
      protocol    = "tcp"
      description = "Lustre access from within VPC"
      cidr_blocks = module.vpc.vpc_cidr_block
    },
    {
      from_port   = 1021
      to_port     = 1023
      protocol    = "tcp"
      description = "Lustre access from within VPC"
      cidr_blocks = module.vpc.vpc_cidr_block
    },
  ]

  tags = var.tags

}

resource "aws_fsx_lustre_file_system" "default" {
  subnet_ids                        = element(module.vpc.private_subnets, 0)
  security_group_ids                = [module.security_group.id]
  storage_capacity                  = "1200"
  deployment_type                   = "SCRATCH_2"
  weekly_maintenance_start_time     = "m:00:00"
  automatic_backup_retention_days   = 7
  daily_automatic_backup_start_time = "03:00"
  copy_tags_to_backups              = true
  tags                              = var.tags
}