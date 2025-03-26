package cmd

import (
	"os/exec"
	"strings"
	"testing"
)

func TestCLI_ReportsViolations(t *testing.T) {
	cmd := exec.Command("go", "run", "../main.go", "--file", "testdata/sample-plan.json")
	out, err := cmd.CombinedOutput()
	if err != nil {
		t.Fatalf("CLI run failed: %v\nOutput: %s", err, out)
	}

	output := string(out)
	if !strings.Contains(output, "Violation") {
		t.Errorf("Expected violations, got: %s", output)
	}
}
