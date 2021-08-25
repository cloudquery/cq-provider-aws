// RDS certificate cannot be created without rds instance so we create one

resource "aws_db_instance" "rds_db_instance" {
  allocated_storage    = 10
  engine               = "mysql"
  engine_version       = "5.7"
  instance_class       = "db.t3.micro"
  name                 = "mydb"
  username             = "foo"
  password             = "foobarbaz"
  parameter_group_name = "default.mysql5.7"
  skip_final_snapshot  = true
  ca_cert_identifier = "ca-${var.test_prefix}${var.test_suffix}"

  tags = {
    Name = "rds-${var.test_prefix}${var.test_suffix}"
  }
}