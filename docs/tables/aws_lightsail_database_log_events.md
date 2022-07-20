
# Table: aws_lightsail_database_log_events
Describes a database log event
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|database_cq_id|uuid|Unique CloudQuery ID of aws_lightsail_databases table (FK)|
|created_at|timestamp without time zone|The timestamp when the database log event was created|
|message|text|The message of the database log event|
