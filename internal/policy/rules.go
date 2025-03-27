package policy

import (
	"regexp"
	"terraform-plan-linter/internal/parser"
)

type Violation struct {
	Resource string
	Message  string
}

func CheckPublicS3(plan *parser.Plan) []Violation {
	var violations []Violation
	for _, r := range parser.AllResources(plan) {
		if r.Type == parser.S3Bucket {
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
	for _, r := range parser.AllResources(plan) {
		if r.Type == parser.S3Bucket {
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

func CheckForceDestroy(plan *parser.Plan) []Violation {
	var violations []Violation
	for _, r := range parser.AllResources(plan) {
		if r.Type == "aws_db_instance" || r.Type == parser.S3Bucket {
			value, ok := r.Values["force_destroy"].(bool)
			if !ok {
				violations = append(violations, Violation{
					Resource: r.Address,
					Message:  "Critical resource does not support force_destroy",
				})
			} else if value {
				violations = append(violations, Violation{
					Resource: r.Address,
					Message:  "Critical resource with force_destroy: true",
				})
			}
		}
	}
	return violations
}

// CheckLeastPrivilegeAccess parses IAM policies and finds LPA violations
func CheckLeastPrivilegeAccess(plan *parser.Plan) []Violation {
	violations := make([]Violation, 0)
	for _, r := range parser.AllResources(plan) {
		if r.Type == parser.IAMPolicy {
			policy := r.Values["policy"].(string)
			// Compile the regular expression and find all matches
			if matches := regexp.MustCompile(`"\*"`).FindAllString(policy, -1); len(matches) > 0 {
				violations = append(violations, Violation{
					Resource: r.Address,
					Message:  "IAM policy is too permissive",
				})
			}
		}
	}
	return violations
}
