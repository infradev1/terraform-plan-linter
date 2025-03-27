package main

import "terraform-plan-linter/cmd"

func main() {
	cmd.Execute()
}

// go mod init terraform-plan-linter
// go mod tidy
// cd terraform && terraform plan -out=tfplan.binary && terraform show -json tfplan.binary > tf-plan.json
