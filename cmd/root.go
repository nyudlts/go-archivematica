package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) { fmt.Println("Hi") },
}

func Execute() {
	rootCmd.Execute()
}
