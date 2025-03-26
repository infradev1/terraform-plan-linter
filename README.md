# 🧹 Terraform Plan Linter

[![CI](https://github.com/CarlosLaraFP/terraform-plan-linter/actions/workflows/ci.yml/badge.svg)](https://github.com/CarlosLaraFP/terraform-plan-linter/actions)

Go CLI to lint Terraform .plan JSON output and flag anti-patterns & missing best practices (e.g. untagged AWS resources, public S3 buckets, no lifecycle policy).

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
make plan
make lint
make test
make clean