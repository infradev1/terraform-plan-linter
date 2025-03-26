package parser

import (
	"encoding/json"
	"fmt"
	"os"
)

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
	Address string                 `json:"address"`
	Type    string                 `json:"type"`
	Name    string                 `json:"name"`
	Values  map[string]interface{} `json:"values"`
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
