data "aws_ami" "elbv1-ubuntu" {
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

//todo might be moved to some general file
resource "aws_vpc" "elbv1-vpc" {
  cidr_block = "172.16.0.0/16"

  tags = {
    Name = "tf-example"
  }
}

resource "aws_internet_gateway" "elbv1-gw" {
  vpc_id = aws_vpc.aws_emr_clusters_vpc.id

}

resource "aws_subnet" "elbv1-subnet" {
  vpc_id = aws_vpc.aws_emr_clusters_vpc.id
  cidr_block = "172.16.10.0/24"
  availability_zone = "us-east-1a"

  tags = {
    Name = "tf-example"
  }
}



resource "aws_network_interface" "elbv1-int1" {
  subnet_id = aws_subnet.emr_clusters_subnet.id
  private_ips = [
    "172.16.10.101"]

  tags = {
    Name = "primary_network_interface"
  }
  depends_on = [
    aws_internet_gateway.elbv1-gw]
}


resource "aws_instance" "elbv1-instance-1" {
  ami = data.aws_ami.elbv1-ubuntu.id
  instance_type = "t2.nano"

  network_interface {
    network_interface_id = aws_network_interface.elbv1-int1.id
    device_index = 0
  }

  tags = {
    Name = "HelloWorld"
  }
}

resource "aws_s3_bucket" "elbv1-bucket" {
  bucket = "${var.test_prefix}${var.test_suffix}"
  acl = "private"
}


# Create a new load balancer
resource "aws_elb" "elbv1-loadbalancer" {
  name = "elbv1${var.test_suffix}"

  listener {
    instance_port = 8000
    instance_protocol = "http"
    lb_port = 80
    lb_protocol = "http"
  }

  health_check {
    healthy_threshold = 2
    unhealthy_threshold = 2
    timeout = 3
    target = "HTTP:8000/"
    interval = 30
  }

  subnets = [
    aws_subnet.emr_clusters_subnet.id]
  instances = [
    aws_instance.elbv1-instance-1.id]
  cross_zone_load_balancing = true
  idle_timeout = 400
  connection_draining = true
  connection_draining_timeout = 400



  tags = {
    Name = "foobar-terraform-elb"
  }
}

resource "aws_load_balancer_policy" "elbv1-policy-ssl" {
  load_balancer_name = aws_elb.elbv1-loadbalancer.name
  policy_name = "elbv1-policy${var.test_prefix}${var.test_suffix}"
  policy_type_name = "SSLNegotiationPolicyType"

  policy_attribute {
    name = "ECDHE-ECDSA-AES128-GCM-SHA256"
    value = "true"
  }

  policy_attribute {
    name = "Protocol-TLSv1.2"
    value = "true"
  }
}