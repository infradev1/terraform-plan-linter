provider "aws" {
  region = "us-west-2"
}

resource "aws_s3_bucket" "public_bucket" {
  bucket = "my-public-bucket-${random_id.suffix.hex}"
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
