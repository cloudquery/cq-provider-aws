resource "aws_s3_bucket" "aws_s3_bucket" {
  bucket        = "${var.prefix}-glue-target-bucket"
  force_destroy = true
}


resource "aws_iam_role" "crawler_role" {
  name = "${var.prefix}-crawler-role"

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  assume_role_policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "",
        "Effect" : "Allow",
        "Principal" : {
          "Service" : "sagemaker.amazonaws.com"
        },
        "Action" : "sts:AssumeRole"
      }
    ]
  })
}

resource "aws_glue_catalog_database" "aws_glue_catalog_database" {
  name = "${var.prefix}-glue-database"
}

resource "aws_glue_connection" "aws_glue_connection" {
  connection_properties = {
    JDBC_CONNECTION_URL = "jdbc:mysql://example.com/exampledatabase"
    PASSWORD            = "examplepassword"
    USERNAME            = "exampleusername"
  }

  name = "example"
}

resource "aws_glue_catalog_table" "aws_glue_catalog_table" {
  name          = "${var.prefix}-glue-table"
  database_name = aws_glue_catalog_database.aws_glue_catalog_database.name
}

resource "aws_glue_crawler" "aws_glue_crawler" {
  database_name = aws_glue_catalog_database.aws_glue_catalog_database.name
  name          = "${var.prefix}-crawler"
  role          = aws_iam_role.crawler_role.arn

  jdbc_target {
    connection_name = aws_glue_connection.aws_glue_connection.name
    path            = "database-name/%"
  }

  s3_target {
    path = "s3://${aws_s3_bucket.aws_s3_bucket.bucket}"
  }
  dynamodb_target {
    path = "table-name"
  }

  catalog_target {
    database_name = aws_glue_catalog_database.aws_glue_catalog_database.name
    tables        = [aws_glue_catalog_table.aws_glue_catalog_table.name]
  }

  schema_change_policy {
    delete_behavior = "LOG"
  }

  configuration = <<EOF
{
  "Version":1.0,
  "Grouping": {
    "TableGroupingPolicy": "CombineCompatibleSchemas"
  }
}
EOF
}