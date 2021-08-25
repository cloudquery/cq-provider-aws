//resource "aws_s3_bucket" "s3_rds_bucket" {
//  bucket = "bucket-rds-import-${var.test_prefix}${var.test_suffix}"
//  acl    = "private"
//}
//
//resource "aws_iam_role_policy" "rds_s3_role_policy" {
//  name = "aws_rds_role_policy${var.test_prefix}${var.test_suffix}"
//  role = aws_iam_role.aws_rds_clusters_s3_role.name
//  policy = jsonencode({
//    Version = "2012-10-17"
//    Statement = [
//      {
//        Action = [
//          "s3:PutObject",
//          "s3:GetObject",
//          "s3:AbortMultipartUpload",
//          "s3:ListBucket",
//          "s3:DeleteObject",
//          "s3:GetObjectVersion",
//          "s3:ListMultipartUploadParts"
//        ]
//        Effect   = "Allow"
//        Resource = [
//          "${aws_s3_bucket.s3_rds_bucket.arn}/*",
//          aws_s3_bucket.s3_rds_bucket.arn
//        ]
//      },
//    ]
//  })
//
//  depends_on = [aws_s3_bucket.s3_rds_bucket]
//}
//
//resource "aws_iam_role" "aws_rds_clusters_s3_role" {
//  name               = "rds_s3_role_${var.test_prefix}${var.test_suffix}"
//  path               = "/"
//  assume_role_policy = jsonencode({
//    Version = "2012-10-17"
//    Statement = [
//      {
//        Action = "sts:AssumeRole"
//        Effect = "Allow"
//        Sid    = ""
//        Principal = {
//          Service = "rds.amazonaws.com"
//        }
//      },
//    ]
//  })
//}
//
//resource "aws_rds_cluster_role_association" "rds_role_assoc" {
//  db_cluster_identifier = aws_rds_cluster.rds_cluster.id
//  feature_name          = "rds_s3_role${var.test_suffix}"
//  role_arn              = aws_iam_role.aws_rds_clusters_s3_role.arn
//}

resource "aws_rds_cluster_instance" "cluster_instances" {
  count              = 1
  identifier         = "rdsclusterdb${var.test_suffix}"
  cluster_identifier = aws_rds_cluster.rds_cluster.id
  instance_class     = "db.t3.small"
  engine             = aws_rds_cluster.rds_cluster.engine
  engine_version     = aws_rds_cluster.rds_cluster.engine_version
}

resource "aws_rds_cluster" "rds_cluster" {
  cluster_identifier      = "rdscluster${var.test_suffix}"
  database_name           = "rdsclusterdb${var.test_suffix}"
  master_username         = "foo"
  master_password         = "bar123foo456"
  backup_retention_period = 5
  preferred_backup_window = "07:00-09:00"
}