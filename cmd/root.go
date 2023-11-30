package cmd

import (
	"go_am"

	"github.com/spf13/cobra"
)

var (
	config string
	client *go_am.AMClient
)

var rootCmd = &cobra.Command{
	Use: "",
	Run: func(cmd *cobra.Command, args []string) {},
}

func Execute() {
	rootCmd.Execute()
}
