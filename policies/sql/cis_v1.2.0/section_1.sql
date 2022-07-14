\echo  "Executing CIS V1.2.0 Section 1"
\set check_id "1.1"
\echo "Executing check 1.1"
\i sql/queries/iam/avoid_root_usage.sql
\set check_id "1.2"
\echo "Executing check 1.2"
\i sql/queries/iam/mfa_enabled_for_console_access.sql
\set check_id "1.3"
\echo "Executing check 1.3"
\i sql/queries/iam/unused_creds_disabled.sql
\set check_id "1.4"
\echo "Executing check 1.4"
\i sql/queries/iam/old_access_keys.sql
\set check_id "1.5"
\echo "Executing check 1.5"
\i sql/queries/iam/password_policy_min_uppercase.sql
\set check_id "1.6"
\echo "Executing check 1.6"
\i sql/queries/iam/password_policy_min_lowercase.sql
\set check_id "1.7"
\echo "Executing check 1.7"
\i sql/queries/iam/password_policy_min_one_symbol.sql
\set check_id "1.8"
\echo "Executing check 1.8"
\i sql/queries/iam/password_policy_min_number.sql
\set check_id "1.9"
\echo "Executing check 1.9"
\i sql/queries/iam/password_policy_min_length.sql
\set check_id "1.10"
\echo "Executing check 1.10"
\i sql/queries/iam/password_policy_prevent_reuse.sql
\set check_id "1.11"
\echo "Executing check 1.11"
\i sql/queries/iam/password_policy_expire_old_passwords.sql
\set check_id "1.12"
\echo "Executing check 1.12"
\i sql/queries/iam/root_user_no_access_keys.sql
\set check_id "1.13"
\echo "Executing check 1.13"
\i sql/queries/iam/mfa_enabled_for_root.sql
\set check_id "1.14"
\echo "Executing check 1.14"
\i sql/queries/iam/hardware_mfa_enabled_for_root.sql
\set check_id "1.16"
\echo "Executing check 1.16"
\i sql/queries/iam/policies_attached_to_groups_roles.sql
