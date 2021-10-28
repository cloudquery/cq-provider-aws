terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 3.55"
    }

    docker = {
      source  = "kreuzwerker/docker"
      version = "2.15.0"
    }
  }
}
