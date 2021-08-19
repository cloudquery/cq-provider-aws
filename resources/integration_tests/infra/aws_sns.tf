resource "aws_sns_topic" "sns_user_updates" {
  name = "sns-user-updates${var.test_prefix}-${var.test_suffix}"
  display_name = "${var.test_prefix}${var.test_suffix}"
  fifo_topic = true
  tags = {
    Environment = var.test_suffix
    Creator = var.test_prefix
  }
  delivery_policy = <<EOF
{
  "http": {
    "defaultHealthyRetryPolicy": {
      "minDelayTarget": 20,
      "maxDelayTarget": 20,
      "numRetries": 3,
      "numMaxDelayRetries": 0,
      "numNoDelayRetries": 0,
      "numMinDelayRetries": 0,
      "backoffFunction": "linear"
    },
    "disableSubscriptionOverrides": false,
    "defaultThrottlePolicy": {
      "maxReceivesPerSecond": 1
    }
  }
}
EOF
}

resource "aws_sqs_queue" "sns_test_queue" {
  name = "sns_test_queue${var.test_prefix}${var.test_suffix}"
}

resource "aws_sqs_queue_policy" "sns_test_policy" {
  queue_url = aws_sqs_queue.sns_test_queue.id

  policy = <<POLICY
{
  "Version": "2012-10-17",
  "Id": "sqspolicy",
  "Statement": [
    {
      "Sid": "First",
      "Effect": "Allow",
      "Principal": "*",
      "Action": "sqs:SendMessage",
      "Resource": "${aws_sqs_queue.sns_test_queue.arn}",
      "Condition": {
        "ArnEquals": {
          "aws:SourceArn": "${aws_sns_topic.sns_user_updates.arn}"
        }
      }
    }
  ]
}
POLICY
}

resource "aws_sns_topic_subscription" "user_updates_sqs_target" {
  topic_arn = aws_sns_topic.sns_user_updates.arn
  protocol = "sqs"
  endpoint = aws_sqs_queue.sns_test_queue.arn
}