package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

const S3Bucket string = "aws_s3_bucket"

type Plan struct {
	PlannedValues PlannedValues `json:"planned_values"`
}

type PlannedValues struct {
	RootModule RootModule `json:"root_module"`
}

type RootModule struct {
	Resources    []Resource    `json:"resources"`
	ChildModules []ChildModule `json:"child_modules,omitempty"`
}

type ChildModule struct {
	Resources []Resource `json:"resources"`
}

type Resource struct {
	Address string         `json:"address"`
	Type    string         `json:"type"`
	Name    string         `json:"name"`
	Values  map[string]any `json:"values"`
}

// AllResources flattens a Terraform Plan into a slice of resources
func AllResources(plan *Plan) []Resource {
	all := plan.PlannedValues.RootModule.Resources
	for _, child := range plan.PlannedValues.RootModule.ChildModules {
		all = append(all, child.Resources...)
	}
	return all
}

func LoadPlan(filePath string) (*Plan, error) {
	data, err := os.ReadFile(filePath)
	if err != nil {
		return nil, fmt.Errorf("error reading plan: %w", err)
	}

	var plan Plan
	if err := json.Unmarshal(data, &plan); err != nil {
		return nil, fmt.Errorf("invalid plan format: %w", err)
	}
	return &plan, nil
}
