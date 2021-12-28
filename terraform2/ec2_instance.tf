
// This is based on this module
// https://github.com/terraform-aws-modules/terraform-aws-ec2-instance/blob/master/examples/complete/main.tf

module "ec2_vpc" {
  source  = "terraform-aws-modules/vpc/aws"
  version = "~> 3.0"

  name = "cq-provider-aws-ec2-instance-vpc"
  cidr = "10.99.0.0/18"

  azs              = ["eu-central-1a", "eu-central-1b", "eu-central-1c"]
  private_subnets  = ["10.99.3.0/24", "10.99.4.0/24", "10.99.5.0/24"]
}

module "ec2_security_group" {
  source  = "terraform-aws-modules/security-group/aws"
  version = "~> 4.0"

  name        = "cq-provider-aws-ec2-security-group"
  description = "Security group for for cq-provider-aws-ec2-instance"
  vpc_id      = module.ec2_vpc.vpc_id

//   ingress_cidr_blocks = ["0.0.0.0/0"]
//   ingress_rules       = ["http-80-tcp", "all-icmp"]
//   egress_rules        = ["all-all"]
}

// resource "aws_placement_group" "web" {
//   name     = "cq-provider-aws-placement-group"
//   strategy = "cluster"
// }

resource "aws_kms_key" "ec2_kms_key" {
}

module "ec2_instance" {
  source = "terraform-aws-modules/ec2-instance/aws"

  name = "cq-provider-aws-ec2-instance"
  // create_spot_instance = true
  // spot_price           = "0.60"
  // spot_type            = "persistent"

  ami                         = "ami-05d34d340fb1d89e5"
  instance_type               = "t2.micro"
  availability_zone           = element(module.ec2_vpc.azs, 0)
  subnet_id                   = element(module.ec2_vpc.private_subnets, 0)
  vpc_security_group_ids      = [module.ec2_security_group.security_group_id]
//   placement_group             = aws_placement_group.web.id
  associate_public_ip_address = true

  # only one of these can be enabled at a time
  hibernation = true
  # enclave_options_enabled = true

  //   user_data_base64 = base64encode(local.user_data)

  cpu_core_count       = 2 # default 4
  cpu_threads_per_core = 1 # default 2

  capacity_reservation_specification = {
    capacity_reservation_preference = "open"
  }

  enable_volume_tags = false
  root_block_device = [
    {
      encrypted   = true
      volume_type = "gp3"
      throughput  = 200
      volume_size = 50
      tags = {
        Name = "my-root-block"
      }
    },
  ]

  // ebs_block_device = [
  //   {
  //     device_name = "/dev/sdf"
  //     volume_type = "gp3"
  //     volume_size = 5
  //     throughput  = 200
  //     encrypted   = true
  //     kms_key_id  = aws_kms_key.ec2_kms_key.arn
  //   }
  // ]

}