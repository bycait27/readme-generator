package cmd

import (
	"fmt"

	"github.com/bycait27/readme-generator/internal/generator"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List available templates",
	Long:  "List all available README templates",
	Run: func(cmd *cobra.Command, args []string) {
		templates, err := generator.ListAvailableTemplates()
		if err != nil {
			fmt.Printf("❌ Error listing templates: %v\n", err)
			return
		}

		fmt.Println("📋 Available templates:")
		for _, template := range templates {
			fmt.Printf("  • %s\n", template)
		}
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
