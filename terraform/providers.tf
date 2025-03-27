terraform {
  required_providers {
    aws = {
      source  = "hashicorp/aws"
      version = "~> 5.0"
    }
  }
}

// for local static analysis; for real provisioning, use 'aws configure'
provider "aws" {
  access_key = "MOCK"
  secret_key = "MOCK"
  region     = "us-west-2"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}