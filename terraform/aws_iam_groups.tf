resource "aws_iam_group" "group_developers" {
  name = "aws_iam_grouptest"
  path = "/users/"
}

resource "aws_iam_group_policy" "group_policy" {
  name  = "aws_iam_group_policytest"
  group = aws_iam_group.group_developers.name

  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Action": [
        "ec2:Describe*"
      ],
      "Effect": "Allow",
      "Resource": "*"
    }
  ]
}
EOF
}
