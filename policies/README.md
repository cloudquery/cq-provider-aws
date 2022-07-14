# CloudQuery Policies
CloudQuery SQL Policies for AWS

## Policies and Compliance Frameworks Available

- [AWS CIS V1.2.0](./sql/cis_v1.2.0/policy.sql)
- [AWS PCI DSS v3.2.1](./sql/pci_dss_v3.2.1/policy.sql)
- [AWS Foundational Security Best Practices](./sql/foundational_security/policy.sql)
- [AWS Public Egress](./sql/public_egress/policy.sql)
- [AWS Publicly Available](./sql/publicly_available/policy.sql)

## Running

You can execute policies with `psql`. For example:

```bash
# Execute the whole CIS Policy
psql -U postgres -f  ./sql/cis_v1.2.0/policy.sql
```

This will create all the results in `aws_policy_results` table which you can query directly, connect to any BI system (Grafana, Preset, AWS QuickSight, PowerBI, ...).

You can also output it into CSV or HTML with the following built-in psql commands:

```
# default tabular output
psql -U postgres -c "select * from aws_policy_results"
# CSV output
psql -U postgres -c "select * from aws_policy_results" --csv
# HTML output
psql -U postgres -c "select * from aws_policy_results" --html
```
