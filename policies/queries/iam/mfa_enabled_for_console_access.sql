insert into aws_policy_results
select
  :'execution_time' as execution_time,
  :'framework' as framework,
  :'check_id' as check_id,
  'Ensure MFA is enabled for all IAM users that have a console password (Scored)' as title,
  account_id,
  arn as resource_id,
  case when
    password_enabled and not mfa_active
    then 'fail'
    else 'pass'
  end as status
from aws_iam_users
