// for local static analysis; for real provisioning, use 'aws configure'
provider "aws" {
  access_key = "MOCK"
  secret_key = "MOCK"
  region     = "us-west-2"
  skip_credentials_validation = true
  skip_metadata_api_check     = true
  skip_requesting_account_id  = true
}

resource "aws_s3_bucket" "public_bucket" {
  bucket = "my-public-bucket-${random_id.suffix.hex}"
  force_destroy = true
}

resource "aws_s3_bucket_acl" "public" {
  bucket = aws_s3_bucket.public_bucket.id
  acl    = "public-read"
}

resource "aws_s3_bucket" "untagged" {
  bucket = "my-untagged-bucket-${random_id.suffix.hex}"
  # no tags — to trigger violation
}

resource "aws_s3_bucket" "missing_lifecycle" {
  bucket = "my-unprotected-bucket-${random_id.suffix.hex}"

  lifecycle {
    prevent_destroy = false
  }

  tags = {
    Env = "dev"
  }
}

resource "random_id" "suffix" {
  byte_length = 4
}
