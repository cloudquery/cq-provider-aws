\set check_id 'EFS.1'
\echo "Executing check EFS.1"
\i sql/queries/efs/unencrypted_efs_filesystems.sql

\set check_id 'EFS.2'
\echo "Executing check EFS.2"
\i sql/queries/efs/efs_filesystems_with_disabled_backups.sql
