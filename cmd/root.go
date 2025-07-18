package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
    Use:   "ghkit",
    Short: "ghkit - A CLI tool to scaffold GitHub repository essentials",
}

func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println("Error:", err)
    }
}
