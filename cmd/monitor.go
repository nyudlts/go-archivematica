package cmd

import (
	"flag"
	"go_am"

	"github.com/spf13/cobra"
)

func init() {
	monitorCommand.Flags().StringVar(&config, "config", "", "")
	rootCmd.AddCommand(monitorCommand)
}

var monitorCommand = &cobra.Command{
	Use: "monitor",
	Run: func(cmd *cobra.Command, args []string) {
		flag.Parse()
		client, err := go_am.NewAMClient(config, 20)
		if err != nil {
			panic(err)
		}

		client.Monitor()
	},
}
