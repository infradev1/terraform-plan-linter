package cmd

import (
	"fmt"
	"os"

	"terraform-plan-linter/internal/parser"
	"terraform-plan-linter/internal/policy"

	"github.com/spf13/cobra"
)

var planFile string

var rootCmd = &cobra.Command{
	Use:   "tflint",
	Short: "Lint Terraform plans for anti-patterns and best practices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("📄 Scanning plan file: %s\n", planFile)

		plan, err := parser.LoadPlan(planFile)
		if err != nil {
			fmt.Fprintf(os.Stderr, "❌ Failed to parse plan: %v\n", err)
			os.Exit(1)
		}

		var violations []policy.Violation
		violations = append(violations, policy.CheckPublicS3(plan)...)
		violations = append(violations, policy.CheckUntaggedBuckets(plan)...)
		violations = append(violations, policy.CheckForceDestroy(plan)...)

		if len(violations) == 0 {
			fmt.Println("✅ No violations found.")
			return
		}

		for _, v := range violations {
			fmt.Printf("[Violation] %s: %s\n", v.Resource, v.Message)
		}
	},
}

func Execute() {
	rootCmd.PersistentFlags().StringVarP(&planFile, "file", "f", "", "Terraform plan JSON file to scan")
	_ = rootCmd.MarkPersistentFlagRequired("file")

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
