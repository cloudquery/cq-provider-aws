data "aws_ami" "aws_ecs_clusters_ami_ubuntu" {
  most_recent = true

  filter {
    name = "name"
    values = [
      "ubuntu/images/hvm-ssd/ubuntu-focal-20.04-amd64-server-*"]
  }

  filter {
    name = "virtualization-type"
    values = [
      "hvm"]
  }

  owners = [
    "099720109477"]
  # Canonical
}

resource "aws_vpc" "aws_ecs_clusters_vpc" {
  cidr_block = "10.0.0.0/16"
}

# Declare the data source
data "aws_availability_zones" "aws_ecs_clusters_available" {
  state = "available"
}

resource "aws_subnet" "aws_ecs_clusters_subnet1" {
  vpc_id = aws_vpc.aws_ecs_clusters_vpc.id
  availability_zone = data.aws_availability_zones.aws_ecs_clusters_available.names[0]
  cidr_block = "10.0.1.0/24"
}

resource "aws_subnet" "aws_ecs_clusters_subnet2" {
  vpc_id = aws_vpc.aws_ecs_clusters_vpc.id
  availability_zone = data.aws_availability_zones.aws_ecs_clusters_available.names[1]
  cidr_block = "10.0.2.0/24"
}


resource "aws_security_group" "aws_ecs_clusters_security_group" {
  name = "aws_ecs_clusters_security_group${var.test_prefix}${var.test_suffix}"
  description = "controls access to the application LB"

  vpc_id = aws_vpc.aws_ecs_clusters_vpc.id

  ingress {
    protocol = "tcp"
    from_port = "80"
    to_port = "80"
    //    cidr_blocks = ["${split(",",var.lb_internal ? var.vpc_cidr : join(",",var.public_alb_whitelist))}"]
  }

  egress {
    from_port = 0
    to_port = 0
    protocol = "-1"

    cidr_blocks = [
      "0.0.0.0/0"]
  }
}

resource "aws_iam_role" "aws_ecs_clusters_ecs-instance-role" {
  name = "ecs-instance-role_${var.test_prefix}${var.test_suffix}"
  path = "/"
  assume_role_policy = data.aws_iam_policy_document.aws_ecs_clusters_ecs-instance-policy.json
}


data "aws_iam_policy_document" "aws_ecs_clusters_ecs-instance-policy" {
  statement {
    actions = [
      "sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "ec2.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "aws_ecs_clusters_ecs-instance-role-attachment" {
  role = aws_iam_role.aws_ecs_clusters_ecs-instance-role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceforEC2Role"
}

resource "aws_iam_instance_profile" "aws_ecs_clusters_instance-profile" {
  name = "ecs-instance-profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.aws_ecs_clusters_ecs-instance-role.id
  //  provisioner "local-exec" {
  //    command = "sleep 60"
  //  }
}

resource "aws_iam_role" "aws_ecs_clusters_service-role" {
  name = "ecs_service_role_${var.test_prefix}${var.test_suffix}"
  path = "/"
  assume_role_policy = data.aws_iam_policy_document.aws_ecs_clusters_service-policy.json
}

resource "aws_iam_role_policy_attachment" "aws_ecs_clusters_service-role-attachment" {
  role = aws_iam_role.aws_ecs_clusters_service-role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonEC2ContainerServiceRole"
}

data "aws_iam_policy_document" "aws_ecs_clusters_service-policy" {
  statement {
    actions = [
      "sts:AssumeRole"]
    principals {
      type = "Service"
      identifiers = [
        "ecs.amazonaws.com"]
    }
  }
}

##########################################################
# AWS ECS-CLUSTER
#########################################################

resource "aws_ecs_cluster" "aws_ecs_clusters_cluster" {
  name = "ecs_cluster_${var.test_prefix}${var.test_suffix}"
  setting {
    name = "containerInsights"
    value = "enabled"
  }
  tags = {
    name = "ecs_cluster_${var.test_prefix}${var.test_suffix}"
  }
}

###########################################################
# AWS ECS-EC2
###########################################################

resource "aws_iam_role" "aws_ecs_clusters_ec2_iam_role" {
  name = "ec2_iam_role_${var.test_prefix}${var.test_suffix}"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}


resource "aws_iam_instance_profile" "aws_ecs_clusters_ec2-instance-profile" {
  name = "ec2_instance_profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.aws_ecs_clusters_ec2_iam_role.id
  //  provisioner "local-exec" {
  //    command = "sleep 60"
  //  }
}


resource "aws_instance" "aws_ecs_clusters_ec2_instance" {
  ami = data.aws_ami.aws_ecs_clusters_ami_ubuntu.id
  subnet_id = aws_subnet.aws_ecs_clusters_subnet1.id
  instance_type = "t2.nano"
  iam_instance_profile = aws_iam_instance_profile.aws_ecs_clusters_ec2-instance-profile.name
  vpc_security_group_ids = [
    aws_security_group.aws_emr_clusters_security_group.id]
  //  key_name = "pnl-test"
  //  #CHANGE THIS
  ebs_optimized = "false"
  source_dest_check = "false"
  user_data = data.template_file.aws_ecs_clusters_user_data.rendered
  root_block_device {
    volume_type = "gp2"
    volume_size = "30"
    delete_on_termination = "true"
  }

  lifecycle {
    ignore_changes = [
      "ami",
      "user_data",
      "subnet_id",
      "key_name",
      "ebs_optimized",
      "private_ip"]
  }
}

data "template_file" "aws_ecs_clusters_user_data" {
  template = <<EOF
#!/bin/bash

# Update all packages

sudo yum update -y
sudo yum install -y ecs-init
sudo service docker start
sudo start ecs

#Adding cluster name in ecs config
echo ECS_CLUSTER=openapi-devl-cluster >> /etc/ecs/ecs.config
cat /etc/ecs/ecs.config | grep "ECS_CLUSTER"
EOF
}

############################################################
# AWS ECS-TASK
############################################################

resource "aws_ecs_task_definition" "aws_ecs_clusters_task_definition" {
  container_definitions = jsonencode([
    {
      "name": "web-server",
      "image": "nginx",
      # ecs registry image url
      "cpu": 10,
      "memory": 512,
      "logConfiguration": {
        "logDriver": "awslogs",
        "options": {
          "awslogs-group": "openapi-devl-cw",
          "awslogs-region": "eu-west-1",
          "awslogs-stream-prefix": "ecs"
        }
      },
      "links": [],
      "portMappings": [
        {
          "hostPort": 8080,
          "containerPort": 8080,
          "protocol": "tcp"
        }
      ],
      "essential": true,
      "entryPoint": [],
      "command": [],
      "environment": [],
      "mountPoints": [],
      "volumesFrom": []
    }
  ])
  # task defination json file location
  //  execution_role_arn = "EcsTaskExecutionRole"
  //  #CHANGE THIS                                                                      # role for executing task
  family = "openapi-task-defination"
  # task name
  network_mode = "awsvpc"
  # network mode awsvpc, brigde
  memory = "2048"
  cpu = "1024"
  requires_compatibilities = [
    "EC2"]
  //  # Fargate or EC2
  //  task_role_arn = "EcsTaskExecutionRole"
  //  #CHANGE THIS                                                                     # TASK running role
}


##############################################################
# AWS ECS-SERVICE
##############################################################

resource "aws_ecs_service" "aws_ecs_clusters_service" {
  cluster = aws_ecs_cluster.aws_ecs_clusters_cluster.id
  # ecs cluster id
  desired_count = 1
  # no of task running
  launch_type = "EC2"
  # Cluster type ECS OR FARGATE
  name = "ecs_service_${var.test_prefix}${var.test_suffix}"
  # Name of service
  task_definition = aws_ecs_task_definition.aws_ecs_clusters_task_definition.arn
  # Attaching Task to service
  load_balancer {
    container_name = "web-server"
    #"container_${var.component}_${var.environment}"
    container_port = "8080"
    target_group_arn = aws_lb_target_group.aws_ecs_clusters_lb_target_group.arn
    # attaching load_balancer target group to ecs
  }
  network_configuration {
    security_groups = [
      aws_security_group.aws_emr_clusters_security_group.id]
    subnets = [
      aws_subnet.aws_ecs_clusters_subnet1.id,
      aws_subnet.aws_ecs_clusters_subnet2.id]
    ## Enter the private subnet id
    assign_public_ip = "false"
  }
  depends_on = [
    "aws_lb_listener.aws_ecs_clusters_lb_listener"]
}

####################################################################
# AWS ECS-ALB
#####################################################################

resource "aws_lb" "ecs_clusters_load_balancer" {
  internal = true
  # internal = true else false
  name = "lb-${var.test_prefix}${var.test_suffix}"
  subnets = [
    aws_subnet.aws_ecs_clusters_subnet1.id,
    aws_subnet.aws_ecs_clusters_subnet2.id]
  # enter the private subnet
  security_groups = [
    aws_security_group.aws_emr_clusters_security_group.id]
}


resource "aws_lb_target_group" "aws_ecs_clusters_lb_target_group" {
  name = "tg-${var.test_prefix}${var.test_suffix}"
  port = "80"
  protocol = "HTTP"
  vpc_id = aws_vpc.aws_ecs_clusters_vpc.id
  target_type = "ip"


  #STEP 1 - ECS task Running
  health_check {
    healthy_threshold = "3"
    interval = "10"
    port = "8080"
    path = "/index.html"
    protocol = "HTTP"
    unhealthy_threshold = "3"
  }
}

resource "aws_lb_listener" "aws_ecs_clusters_lb_listener" {
  default_action {
    target_group_arn = aws_lb_target_group.aws_ecs_clusters_lb_target_group.id
    type = "forward"
  }

  #certificate_arn   = "arn:aws:acm:us-east-1:689019322137:certificate/9fcdad0a-7350-476c-b7bd-3a530cf03090"
  load_balancer_arn = aws_lb.ecs_clusters_load_balancer.arn
  port = "80"
  protocol = "HTTP"
}

####################################################################
# AWS CLOUDWATCH
#####################################################################

resource "aws_cloudwatch_log_group" "aws_ecs_clusters_log_group" {
  name = "log_group_${var.test_prefix}${var.test_suffix}"
  tags = {
    Environment = "production"
  }
}
