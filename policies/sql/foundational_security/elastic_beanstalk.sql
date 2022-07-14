\set check_id 'ElasticBeanstalk.1'
\echo "Executing check ElasticBeanstalk.1"
\i sql/queries/elasticbeanstalk/advanced_health_reporting_enabled.sql

\set check_id 'ElasticBeanstalk.2'
\echo "Executing check ElasticBeanstalk.2"
\i sql/queries/elasticbeanstalk/elastic_beanstalk_managed_updates_enabled.sql
