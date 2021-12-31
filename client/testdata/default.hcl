cloudquery {
  plugin_directory = "./cq/providers"
  policy_directory = "./cq/policies"

  provider "aws" {
    source = "cloudquery/cq-provider-aws"
    version = "v0.8.4"
  }

  connection {
    dsn = "host=localhost user=postgres password=pass database=postgres port=5432 sslmode=disable"
  }
}

provider "aws" {
  configuration {
    max_retries = 7
    max_backoff = 60
  }
  resources = [
    "*"]
}
