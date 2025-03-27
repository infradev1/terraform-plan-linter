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

resource "aws_iam_role" "permissive_role" {
  name = "permissive_role"

  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })

  tags = {
    env = var.account_id
  }
}

resource "aws_iam_role_policy" "permissive_policy" {
  name = "permissive_policy"
  role = aws_iam_role.permissive_role.id

  # Terraform's "jsonencode" function converts a
  # Terraform expression result to valid JSON syntax.
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "ec2:Describe*",
        ]
        Effect   = "Allow"
        Resource = "*"
      },
    ]
  })
}

resource "random_id" "suffix" {
  byte_length = 4
}
