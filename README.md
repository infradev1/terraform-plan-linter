# terraform-plan-linter

Reads Terraform .plan JSON output and flags anti-patterns (e.g. untagged AWS resources, public S3 buckets, no lifecycle policy).

- Real-world infrastructure problem
- IaC scanning via static analysis
- Policy-as-Code
- Go as a backend tool
- Unit tests & CI