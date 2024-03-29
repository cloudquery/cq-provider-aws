insert into aws_policy_results
select
    :'execution_time' as execution_time,
    :'framework' as framework,
    :'check_id' as check_id,
    'EMR clusters should not have public ip addresses' as title,
    aws_emr_clusters.account_id,
    aws_emr_clusters.arn as resource_id,
    case when aws_ec2_subnets.map_public_ip_on_launch and aws_emr_clusters.state in ('RUNNING', 'WAITING') then 'fail'
    else 'pass' end as status
from
    aws_emr_clusters
left outer join aws_ec2_subnets
    on aws_emr_clusters.ec2_instance_attribute_subnet_id = aws_ec2_subnets.id
