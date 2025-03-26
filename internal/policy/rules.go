package policy

import (
	"terraform-plan-linter/internal/parser"
)

type Violation struct {
	Resource string
	Message  string
}

func CheckPublicS3(plan *parser.Plan) []Violation {
	var violations []Violation
	for _, r := range allResources(plan) {
		if r.Type == "aws_s3_bucket" {
			acl := r.Values["acl"]
			if acl == "public-read" || acl == "public-read-write" {
				violations = append(violations, Violation{
					Resource: r.Address,
					Message:  "S3 bucket allows public access via ACL",
				})
			}
		}
	}
	return violations
}

func CheckUntaggedBuckets(plan *parser.Plan) []Violation {
	var violations []Violation
	for _, r := range allResources(plan) {
		if r.Type == "aws_s3_bucket" {
			if _, ok := r.Values["tags"]; !ok {
				violations = append(violations, Violation{
					Resource: r.Address,
					Message:  "S3 bucket is missing tags",
				})
			}
		}
	}
	return violations
}

func CheckMissingPreventDestroy(plan *parser.Plan) []Violation {
	var violations []Violation
	for _, r := range allResources(plan) {
		if r.Type == "aws_db_instance" || r.Type == "aws_s3_bucket" {
			lifecycle, ok := r.Values["lifecycle"].(map[string]interface{})
			if !ok || lifecycle["prevent_destroy"] != true {
				violations = append(violations, Violation{
					Resource: r.Address,
					Message:  "Missing lifecycle.prevent_destroy on critical resource",
				})
			}
		}
	}
	return violations
}

func allResources(plan *parser.Plan) []parser.Resource {
	all := plan.PlannedValues.RootModule.Resources
	for _, child := range plan.PlannedValues.RootModule.ChildModules {
		all = append(all, child.Resources...)
	}
	return all
}
