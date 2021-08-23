resource "aws_emr_cluster" "aws_emr_clusters_cluster" {
  name = "emr-cluster-${var.test_prefix}${var.test_suffix}"
  release_label = "emr-5.12.0"
  applications = [
    "Spark"]

  log_uri = "s3://${aws_s3_bucket.aws_emr_cluster_logs.id}/"


  additional_info = <<EOF
{
  "instanceAwsClientConfiguration": {
    "proxyPort": 8099,
    "proxyHost": "myproxy.example.com"
  }
}
EOF

  termination_protection = false
  keep_job_flow_alive_when_no_steps = true

  ec2_attributes {
    subnet_id = aws_subnet.emr_clusters_subnet.id
    emr_managed_master_security_group = aws_security_group.aws_emr_clusters_security_group.id
    emr_managed_slave_security_group = aws_security_group.aws_emr_clusters_security_group.id
    instance_profile = aws_iam_instance_profile.aws_emr_clusters_instance_profile.arn
  }

  master_instance_group {
    instance_type = "m1.small"
  }

  core_instance_group {
    instance_type = "m1.small"
    instance_count = 1

    ebs_config {
      size = "40"
      type = "gp2"
      volumes_per_instance = 1
    }

    bid_price = "0.30"




    autoscaling_policy = <<EOF
{
"Constraints": {
  "MinCapacity": 1,
  "MaxCapacity": 2
},
"Rules": [
  {
    "Name": "ScaleOutMemoryPercentage",
    "Description": "Scale out if YARNMemoryAvailablePercentage is less than 15",
    "Action": {
      "SimpleScalingPolicyConfiguration": {
        "AdjustmentType": "CHANGE_IN_CAPACITY",
        "ScalingAdjustment": 1,
        "CoolDown": 300
      }
    },
    "Trigger": {
      "CloudWatchAlarmDefinition": {
        "ComparisonOperator": "LESS_THAN",
        "EvaluationPeriods": 1,
        "MetricName": "YARNMemoryAvailablePercentage",
        "Namespace": "AWS/ElasticMapReduce",
        "Period": 300,
        "Statistic": "AVERAGE",
        "Threshold": 15.0,
        "Unit": "PERCENT"
      }
    }
  }
]
}
EOF
  }

  ebs_root_volume_size = 100

  tags = {
    role = "rolename"
    env = "env"
  }

  bootstrap_action {
    path = "s3://elasticmapreduce/bootstrap-actions/run-if"
    name = "runif"
    args = [
      "instance.isMaster=true",
      "echo running on master node"]
  }

  configurations_json = <<EOF
  [
    {
      "Classification": "hadoop-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    },
    {
      "Classification": "spark-env",
      "Configurations": [
        {
          "Classification": "export",
          "Properties": {
            "JAVA_HOME": "/usr/lib/jvm/java-1.8.0"
          }
        }
      ],
      "Properties": {}
    }
  ]
EOF
  autoscaling_role = aws_iam_role.emr_clusters_autoscaling_role.arn
  //  autoscaling_role = "EMR_AutoScaling_DefaultRole"
  service_role = aws_iam_role.emr_clusters_service_role.arn

}

resource "aws_iam_instance_profile" "aws_emr_clusters_instance_profile" {
  name = "instance_profile_${var.test_prefix}${var.test_suffix}"
  path = "/"
  role = aws_iam_role.emr_clusters_instance_profile_role.id
}


resource "aws_iam_role" "emr_clusters_instance_profile_role" {
  name = "instance_${var.test_prefix}${var.test_suffix}"

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
          Service = [
            "ec2.amazonaws.com"]
        }
      }
    ]
  })

  inline_policy {
    name = "inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement: [
        {
          Action: [
            "cloudwatch:*",
            "dynamodb:*",
            "ec2:Describe*",
            "elasticmapreduce:Describe*",
            "elasticmapreduce:ListBootstrapActions",
            "elasticmapreduce:ListClusters",
            "elasticmapreduce:ListInstanceGroups",
            "elasticmapreduce:ListInstances",
            "elasticmapreduce:ListSteps",
            "kinesis:CreateStream",
            "kinesis:DeleteStream",
            "kinesis:DescribeStream",
            "kinesis:GetRecords",
            "kinesis:GetShardIterator",
            "kinesis:MergeShards",
            "kinesis:PutRecord",
            "kinesis:SplitShard",
            "rds:Describe*",
            "s3:*",
            "sdb:*",
            "sns:*",
            "sqs:*",
            "iam:CreateServiceLinkedRole",
            "glue:CreateDatabase",
            "glue:UpdateDatabase",
            "glue:DeleteDatabase",
            "glue:GetDatabase",
            "glue:GetDatabases",
            "glue:CreateTable",
            "glue:UpdateTable",
            "glue:DeleteTable",
            "glue:GetTable",
            "glue:GetTables",
            "glue:GetTableVersions",
            "glue:CreatePartition",
            "glue:BatchCreatePartition",
            "glue:UpdatePartition",
            "glue:DeletePartition",
            "glue:BatchDeletePartition",
            "glue:GetPartition",
            "glue:GetPartitions",
            "glue:BatchGetPartition",
            "glue:CreateUserDefinedFunction",
            "glue:UpdateUserDefinedFunction",
            "glue:DeleteUserDefinedFunction",
            "glue:GetUserDefinedFunction",
            "glue:GetUserDefinedFunctions"
          ],
          Effect: "Allow",
          Resource: "*"
        }
      ]
    })
  }

}


resource "aws_iam_role" "emr_clusters_autoscaling_role" {
  name = "autoscaling_${var.test_prefix}${var.test_suffix}"

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
          Service = [
            "ec2.amazonaws.com",
            "elasticmapreduce.amazonaws.com"]
        }
      }
    ]
  })

  inline_policy {
    name = "inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement: [
        {
          Action: [
            "cloudwatch:DescribeAlarms",
            "elasticmapreduce:ListInstanceGroups",
            "elasticmapreduce:ModifyInstanceGroups",
            "iam:CreateServiceLinkedRole",
          ],
          Effect: "Allow",
          Resource: "*"
        }
      ]
    })
  }
}

resource "aws_iam_role" "emr_clusters_service_role" {
  name = "service_${var.test_prefix}${var.test_suffix}"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    Version = "2008-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid = ""
        Principal = {
          Service = [
            "elasticmapreduce.amazonaws.com"]
        }
      }
    ]
  })

  inline_policy {
    name = "inline-${var.test_prefix}${var.test_suffix}"

    policy = jsonencode({
      Version = "2012-10-17"
      Statement: [
        {
          Action: [
            "ec2:AuthorizeSecurityGroupEgress",
            "ec2:AuthorizeSecurityGroupIngress",
            "ec2:CancelSpotInstanceRequests",
            "ec2:CreateFleet",
            "ec2:CreateLaunchTemplate",
            "ec2:CreateNetworkInterface",
            "ec2:CreateSecurityGroup",
            "ec2:CreateTags",
            "ec2:DeleteLaunchTemplate",
            "ec2:DeleteNetworkInterface",
            "ec2:DeleteSecurityGroup",
            "ec2:DeleteTags",
            "ec2:DescribeAvailabilityZones",
            "ec2:DescribeAccountAttributes",
            "ec2:DescribeDhcpOptions",
            "ec2:DescribeImages",
            "ec2:DescribeInstanceStatus",
            "ec2:DescribeInstances",
            "ec2:DescribeKeyPairs",
            "ec2:DescribeLaunchTemplates",
            "ec2:DescribeNetworkAcls",
            "ec2:DescribeNetworkInterfaces",
            "ec2:DescribePrefixLists",
            "ec2:DescribeRouteTables",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeSpotInstanceRequests",
            "ec2:DescribeSpotPriceHistory",
            "ec2:DescribeSubnets",
            "ec2:DescribeTags",
            "ec2:DescribeVpcAttribute",
            "ec2:DescribeVpcEndpoints",
            "ec2:DescribeVpcEndpointServices",
            "ec2:DescribeVpcs",
            "ec2:DetachNetworkInterface",
            "ec2:ModifyImageAttribute",
            "ec2:ModifyInstanceAttribute",
            "ec2:RequestSpotInstances",
            "ec2:RevokeSecurityGroupEgress",
            "ec2:RunInstances",
            "ec2:TerminateInstances",
            "ec2:DeleteVolume",
            "ec2:DescribeVolumeStatus",
            "ec2:DescribeInstanceAttribute",
            "ec2:DescribeVolumes",
            "ec2:DetachVolume",
            "iam:GetRole",
            "iam:CreateServiceLinkedRole",
            "iam:GetRolePolicy",
            "iam:ListInstanceProfiles",
            "iam:ListRolePolicies",
            "iam:PassRole",
            "s3:CreateBucket",
            "s3:Get*",
            "s3:List*",
            "sdb:BatchPutAttributes",
            "sdb:Select",
            "sqs:CreateQueue",
            "sqs:Delete*",
            "sqs:GetQueue*",
            "sqs:PurgeQueue",
            "sqs:ReceiveMessage",
            "cloudwatch:PutMetricAlarm",
            "cloudwatch:DescribeAlarms",
            "cloudwatch:DeleteAlarms",
            "application-autoscaling:RegisterScalableTarget",
            "application-autoscaling:DeregisterScalableTarget",
            "application-autoscaling:PutScalingPolicy",
            "application-autoscaling:DeleteScalingPolicy",
            "application-autoscaling:Describe*"
          ],
          Effect: "Allow",
          Resource: "*"
        }
      ]
    })
  }

}

resource "aws_s3_bucket" "aws_emr_cluster_logs" {
  bucket = "${var.test_prefix}${var.test_suffix}"
  acl = "private"

  tags = {
    Name = "${var.test_prefix}${var.test_suffix}"
    Environment = "Dev"
  }
}


//todo might be moved to some general file
resource "aws_vpc" "aws_emr_clusters_vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "tf-example"
  }
}


resource "aws_route_table" "aws_emr_clusters_rt" {
  vpc_id = aws_vpc.aws_emr_clusters_vpc.id


}

resource "aws_route" "public_internet_gateway" {
  route_table_id         = aws_route_table.aws_emr_clusters_rt.id
  destination_cidr_block = "0.0.0.0/0"
  gateway_id             = aws_internet_gateway.aws_emr_clusters_gw.id

  timeouts {
    create = "5m"
  }
}

resource "aws_route_table_association" "aws_ecs_clusters_a" {
  subnet_id = aws_subnet.emr_clusters_subnet.id
  route_table_id = aws_route_table.aws_emr_clusters_rt.id
}

resource "aws_internet_gateway" "aws_emr_clusters_gw" {
  vpc_id = aws_vpc.aws_emr_clusters_vpc.id
}


resource "aws_subnet" "emr_clusters_subnet" {
  vpc_id = aws_vpc.aws_emr_clusters_vpc.id
  cidr_block = "172.16.10.0/24"
  availability_zone = "us-east-1a"

  tags = {
    Name = "tf-example"
  }
}


resource "aws_security_group_rule" "allow_tcp_from_master_to_service" {
  type = "ingress"
  from_port = 9443
  to_port = 9443
  protocol = "tcp"
  security_group_id = aws_security_group.aws_emr_clusters_security_group.id
  source_security_group_id = aws_security_group.aws_emr_clusters_security_group.id
}

resource "aws_security_group" "aws_emr_clusters_security_group" {
  name = "aws_ecs_clusters_security_group${var.test_prefix}${var.test_suffix}"

  vpc_id = aws_vpc.aws_emr_clusters_vpc.id

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
    ipv6_cidr_blocks = [
      "::/0"]
  }
}
