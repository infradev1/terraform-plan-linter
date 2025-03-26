package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var planFile string

var rootCmd = &cobra.Command{
	Use:   "tflint",
	Short: "Lint Terraform plans for anti-patterns and best practices",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("📄 Scanning plan file: %s\n", planFile)
		// TODO: parse and evaluate plan
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
