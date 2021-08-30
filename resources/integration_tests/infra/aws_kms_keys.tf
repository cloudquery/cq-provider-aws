resource "aws_kms_key" "aws_kms_keys_key" {
  description = "kms-key-${var.test_prefix}${var.test_suffix}"
  deletion_window_in_days = 10
//  policy = <<POLICY
//{
//    "Version": "2012-10-17",
//    "Statement": [
//        {
//            "Sid": "AWSCloudTrailAclCheck",
//            "Effect": "Allow",
//            "Principal": {
//              "Service": "cloudtrail.amazonaws.com"
//            },
//            "Action": "s3:GetBucketAcl",
//            "Resource": "arn:aws:s3:::tb-${var.test_prefix}-${var.test_suffix}"
//        },
//        {
//            "Sid": "AWSCloudTrailWrite",
//            "Effect": "Allow",
//            "Principal": {
//              "Service": "cloudtrail.amazonaws.com"
//            },
//            "Action": "s3:PutObject",
//            "Resource": "arn:aws:s3:::tb-${var.test_prefix}-${var.test_suffix}/*",
//            "Condition": {
//                "StringEquals": {
//                    "s3:x-amz-acl": "bucket-owner-full-control"
//                }
//            }
//        },
//        {
//            "Sid": "AWSGlueReadLogs",
//            "Effect": "Allow",
//            "Principal": {
//              "Service": "glue.amazonaws.com"
//            },
//            "Action": "s3:GetObject",
//            "Resource": "arn:aws:s3:::tb-${var.test_prefix}-${var.test_suffix}/*"
//        }
//    ]
//}
//POLICY
}


resource "time_sleep" "aws_directconnect_virtual_interfaces_wait_for_id" {
  depends_on = [
    aws_kms_key.aws_kms_keys_key]

  create_duration = "10m"
}