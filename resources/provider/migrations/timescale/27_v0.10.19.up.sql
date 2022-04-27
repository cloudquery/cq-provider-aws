
-- Resource: athena.data_catalogs
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalogs" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"account_id" text,
	"region" text,
	"tags" jsonb,
	"name" text,
	"type" text,
	"description" text,
	"parameters" jsonb,
	CONSTRAINT aws_athena_data_catalogs_pk PRIMARY KEY(cq_fetch_date,account_id,region,name),
	UNIQUE(cq_fetch_date,cq_id)
);
SELECT setup_tsdb_parent('aws_athena_data_catalogs');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_databases" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"data_catalog_cq_id" uuid,
	"name" text,
	"description" text,
	"parameters" jsonb,
	CONSTRAINT aws_athena_data_catalog_databases_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_databases (cq_fetch_date, data_catalog_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_databases', 'data_catalog_cq_id', 'aws_athena_data_catalogs', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_tables" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"data_catalog_database_cq_id" uuid,
	"name" text,
	"create_time" timestamp without time zone,
	"last_access_time" timestamp without time zone,
	"parameters" jsonb,
	"table_type" text,
	CONSTRAINT aws_athena_data_catalog_database_tables_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_database_tables (cq_fetch_date, data_catalog_database_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_database_tables', 'data_catalog_database_cq_id', 'aws_athena_data_catalog_databases', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_table_columns" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"data_catalog_database_table_cq_id" uuid,
	"name" text,
	"comment" text,
	"type" text,
	CONSTRAINT aws_athena_data_catalog_database_table_columns_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_database_table_columns (cq_fetch_date, data_catalog_database_table_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_database_table_columns', 'data_catalog_database_table_cq_id', 'aws_athena_data_catalog_database_tables', 'cq_id');
CREATE TABLE IF NOT EXISTS "aws_athena_data_catalog_database_table_partition_keys" (
	"cq_id" uuid NOT NULL,
	"cq_meta" jsonb,
	"cq_fetch_date" timestamp without time zone NOT NULL,
	"data_catalog_database_table_cq_id" uuid,
	"name" text,
	"comment" text,
	"type" text,
	CONSTRAINT aws_athena_data_catalog_database_table_partition_keys_pk PRIMARY KEY(cq_fetch_date,cq_id),
	UNIQUE(cq_fetch_date,cq_id)
);
CREATE INDEX ON aws_athena_data_catalog_database_table_partition_keys (cq_fetch_date, data_catalog_database_table_cq_id);
SELECT setup_tsdb_child('aws_athena_data_catalog_database_table_partition_keys', 'data_catalog_database_table_cq_id', 'aws_athena_data_catalog_database_tables', 'cq_id');
