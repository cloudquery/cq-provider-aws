resource "aws_glue_catalog_database" "aws_glue_catalog_database" {
  name = "${var.prefix}-glue-database"
}

resource "aws_glue_catalog_table" "aws_glue_catalog_table" {
  name          = "${var.prefix}-glue-table"
  database_name = aws_glue_catalog_database.aws_glue_catalog_database.name
}