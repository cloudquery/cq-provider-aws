insert into aws_policy_results
select
  :execution_time,
  :'framework',
  :'check_id',
  'Ensure IAM password policy requires at least one symbol',
  account_id,
  account_id,
  case when
    require_symbols = false or policy_exists = false
    then 'fail'
    else 'pass'
  end as status
from
    aws_iam_password_policies
