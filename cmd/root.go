package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "readme-gen",
	Short: "Professional README generator",
	Long: `readme-gen is a CLI professional README generator that helps 
you create professional looking README files for your projects 
fast.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
