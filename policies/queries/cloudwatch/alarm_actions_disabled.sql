insert into aws_policy_results
select :'execution_time'           as execution_time,
       :'framework'                as framework,
       :'check_id'                 as check_id,
       'Disabled Cloudwatch alarm' as title,
       account_id,
       arn                         as resource_id,
       'fail'                      as status
from aws_cloudwatch_alarms
where actions_enabled = false
   or array_length(actions, 1) = 0