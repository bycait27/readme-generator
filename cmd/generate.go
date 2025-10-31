package cmd

import (
	"fmt"

	"github.com/bycait27/readme-generator/internal/generator"
	"github.com/bycait27/readme-generator/internal/prompts"
	"github.com/spf13/cobra"
)

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a README file",
	Long:  "Generate a professional README file based on your project details",
	Run: func(cmd *cobra.Command, args []string) {
		template, err := cmd.Flags().GetString("template")
		if err != nil {
			fmt.Printf("❌ Error: failed to get template flag: %v\n", err)
		}

		output, err := cmd.Flags().GetString("output")
		if err != nil {
			fmt.Printf("❌ Error: failed to get output flag: %v\n", err)
		}

		// validate template exists
		if err := generator.ValidateTemplate(template); err != nil {
			fmt.Printf("❌ Error: %v\n", err)
			return
		}

		// get project information through prompts
		fmt.Println("🚀 Let's create your README!")
		baseInfo, err := prompts.PromptBaseInfo()
		if err != nil {
			fmt.Printf("❌ Error collecting project info: %v\n", err)
			return
		}

		// generate README
		var genErr error
		if output == "README.md" {
			genErr = generator.GenerateREADME(template, baseInfo)
		} else {
			genErr = generator.GenerateREADMEToFile(template, baseInfo, output)
		}

		if genErr != nil {
			fmt.Printf("❌ Error generating README: %v\n", genErr)
			return
		}

		fmt.Printf("✅ README generated successfully: %s\n", output)
		fmt.Println("📝 Don't forget to customize the installation and usage sections!")
	},
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringP("template", "t", "basic", "Template to use (basic)")
	generateCmd.Flags().StringP("output", "o", "README.md", "Output file name")
}
