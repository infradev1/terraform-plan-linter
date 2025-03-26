# 🧹 Terraform Plan Linter

[![CI](https://github.com/your-username/terraform-plan-linter/actions/workflows/ci.yml/badge.svg)](https://github.com/your-username/terraform-plan-linter/actions)

A fast, pluggable CLI to lint Terraform plans for dangerous anti-patterns and missing best practices. Reads Terraform .plan JSON output and flags anti-patterns (e.g. untagged AWS resources, public S3 buckets, no lifecycle policy).

## ✨ Features

- IaC scanning via static analysis
- Policy-as-Code
- Detects public S3 buckets
- Flags untagged AWS resources
- Warns on missing `prevent_destroy` lifecycle rules
- JSON and human-readable output
- Unit-tested and CI-integrated

## 🚀 Usage

```bash
go run main.go --file testdata/sample-plan.json