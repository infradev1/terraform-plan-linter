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
}

resource "aws_s3_bucket" "missing_lifecycle" {
  bucket = "my-unprotected-bucket-${random_id.suffix.hex}"
  
  force_destroy = true

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
