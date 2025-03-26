package policy_test

import (
	"testing"

	"terraform-plan-linter/internal/parser"
	"terraform-plan-linter/internal/policy"
)

func makePlan(resources []parser.Resource) *parser.Plan {
	return &parser.Plan{
		PlannedValues: parser.PlannedValues{
			RootModule: parser.RootModule{
				Resources: resources,
			},
		},
	}
}

func TestCheckPublicS3(t *testing.T) {
	plan := makePlan([]parser.Resource{
		{
			Address: "aws_s3_bucket.public_bucket",
			Type:    "aws_s3_bucket",
			Values:  map[string]interface{}{"acl": "public-read"},
		},
		{
			Address: "aws_s3_bucket.private_bucket",
			Type:    "aws_s3_bucket",
			Values:  map[string]interface{}{"acl": "private"},
		},
	})

	violations := policy.CheckPublicS3(plan)

	if len(violations) != 1 {
		t.Errorf("expected 1 violation, got %d", len(violations))
	}
	if violations[0].Resource != "aws_s3_bucket.public_bucket" {
		t.Errorf("wrong resource flagged: %s", violations[0].Resource)
	}
}

func TestCheckUntaggedBuckets(t *testing.T) {
	plan := makePlan([]parser.Resource{
		{
			Address: "aws_s3_bucket.untagged",
			Type:    "aws_s3_bucket",
			Values:  map[string]interface{}{}, // no tags
		},
		{
			Address: "aws_s3_bucket.tagged",
			Type:    "aws_s3_bucket",
			Values:  map[string]interface{}{"tags": map[string]interface{}{"env": "prod"}},
		},
	})

	violations := policy.CheckUntaggedBuckets(plan)

	if len(violations) != 1 {
		t.Errorf("expected 1 violation, got %d", len(violations))
	}
	if violations[0].Resource != "aws_s3_bucket.untagged" {
		t.Errorf("wrong resource flagged: %s", violations[0].Resource)
	}
}

func TestCheckMissingPreventDestroy(t *testing.T) {
	plan := makePlan([]parser.Resource{
		{
			Address: "aws_db_instance.no_lifecycle",
			Type:    "aws_db_instance",
			Values:  map[string]interface{}{}, // no lifecycle
		},
		{
			Address: "aws_s3_bucket.prevent_destroy_missing",
			Type:    "aws_s3_bucket",
			Values: map[string]interface{}{
				"lifecycle": map[string]interface{}{
					"prevent_destroy": false,
				},
			},
		},
		{
			Address: "aws_s3_bucket.prevent_destroy_ok",
			Type:    "aws_s3_bucket",
			Values: map[string]interface{}{
				"lifecycle": map[string]interface{}{
					"prevent_destroy": true,
				},
			},
		},
	})

	violations := policy.CheckMissingPreventDestroy(plan)

	if len(violations) != 2 {
		t.Errorf("expected 2 violations, got %d", len(violations))
	}

	expected := map[string]bool{
		"aws_db_instance.no_lifecycle":          true,
		"aws_s3_bucket.prevent_destroy_missing": true,
	}

	for _, v := range violations {
		if !expected[v.Resource] {
			t.Errorf("unexpected violation: %s", v.Resource)
		}
	}
}
