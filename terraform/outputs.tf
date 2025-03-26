output "bucket_names" {
  value = [
    aws_s3_bucket.public_bucket.bucket,
    aws_s3_bucket.untagged.bucket,
    aws_s3_bucket.missing_lifecycle.bucket
  ]
}
