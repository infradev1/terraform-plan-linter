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
			Values:  map[string]any{"acl": "public-read"},
		},
		{
			Address: "aws_s3_bucket.private_bucket",
			Type:    "aws_s3_bucket",
			Values:  map[string]any{"acl": "private"},
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
			Values:  map[string]any{}, // no tags
		},
		{
			Address: "aws_s3_bucket.tagged",
			Type:    "aws_s3_bucket",
			Values:  map[string]any{"tags": map[string]any{"env": "prod"}},
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

func TestCheckForceDestroy(t *testing.T) {
	plan := makePlan([]parser.Resource{
		{
			Address: "aws_db_instance.no_lifecycle",
			Type:    "aws_db_instance",
			Values:  map[string]any{}, // no lifecycle
		},
		{
			Address: "aws_s3_bucket.force_destroy",
			Type:    "aws_s3_bucket",
			Values: map[string]any{
				"force_destroy": true,
			},
		},
		{
			Address: "aws_s3_bucket.prevent_destroy",
			Type:    "aws_s3_bucket",
			Values: map[string]any{
				"force_destroy": false,
				"lifecycle": map[string]any{
					"prevent_destroy": true,
				},
			},
		},
	})

	violations := policy.CheckForceDestroy(plan)

	if len(violations) != 2 {
		t.Errorf("expected 2 violations, got %d", len(violations))
	}

	expected := map[string]bool{
		"aws_db_instance.no_lifecycle": true,
		"aws_s3_bucket.force_destroy":  true,
	}

	for _, v := range violations {
		if !expected[v.Resource] {
			t.Errorf("unexpected violation: %s", v.Resource)
		}
	}
}

func TestCheckLeastPrivilegeAccess(t *testing.T) {
	plan := makePlan([]parser.Resource{
		{
			Address: "aws_iam_role_policy.permissive_policy",
			Type:    parser.IAMPolicy,
			Values: map[string]any{
				"policy": "{\"Version\":\"2012-10-17\",\"Statement\":[{\"Action\":[\"ec2:Describe*\"],\"Effect\":\"Allow\",\"Resource\":\"*\"}]}",
			},
		},
	})

	violations := policy.CheckLeastPrivilegeAccess(plan)
	if len(violations) != 1 {
		t.Errorf("expected 1 violation, got %d", len(violations))
	}
	expected := map[string]bool{
		"aws_iam_role_policy.permissive_policy": true,
	}
	for _, v := range violations {
		if _, ok := expected[v.Resource]; !ok {
			t.Errorf("unexpected violation %s", v.Resource)
		}
	}
}
