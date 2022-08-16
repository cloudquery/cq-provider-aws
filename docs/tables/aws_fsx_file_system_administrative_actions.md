
# Table: aws_fsx_file_system_administrative_actions
A list of administrative actions for the file system that are in process or waiting to be processed.
## Columns
| Name        | Type           | Description  |
| ------------- | ------------- | -----  |
|fsx_cq_id|uuid|Unique CloudQuery ID of aws_fsx_file_systems table (FK)|
|action_type|text|Describes the type of administrative action.|
|progress_percent|integer|The percentage-complete status of a STORAGE_OPTIMIZATION administrative action.|
|request_time|timestamp without time zone|The time that the administrative action request was received.|
|status|text|Describes the status of the administrative action.|
