package cmd

import (
	"fmt"
	"go_am"

	"github.com/spf13/cobra"
)

var (
	ingests   bool
	transfers bool
	all       bool
)

func init() {
	clrCommand.Flags().StringVar(&config, "config", "", "")
	clrCommand.Flags().BoolVar(&ingests, "ingests", false, "")
	clrCommand.Flags().BoolVar(&transfers, "transfers", false, "")
	rootCmd.AddCommand(clrCommand)
}

var clrCommand = &cobra.Command{
	Use: "clear",
	Run: func(cmd *cobra.Command, args []string) {

		if ingests || transfers || all {
			var err error
			client, err = go_am.NewAMClient(config, 20)
			if err != nil {
				panic(err)
			}
		}

		if transfers {
			fmt.Println("Clearing completed transfers")
			completedTransfers, err := client.GetCompletedTransfers()
			if err != nil {
				panic(err)
			}

			completedTransfersMap, err := client.GetCompletedTransfersMap(completedTransfers)
			if err != nil {
				panic(err)
			}

			for k, v := range completedTransfersMap {
				fmt.Printf("clearing %s: %s\n", k, v.Name)
				if err := client.DeleteTransfer(v.UUID); err != nil {
					panic(err)
				}
				fmt.Printf("%s: %s cleared\n", k, v.Name)
			}
		}

		if ingests {
			completedIngests, err := client.GetCompletedIngests()
			if err != nil {
				panic(err)
			}

			completedIngestsMap, err := client.GetCompletedIngestsMap(completedIngests)
			if err != nil {
				panic(err)
			}

			for k, v := range completedIngestsMap {
				fmt.Printf("clearing %s: %s\n", k, v.Name)
				if err := client.DeleteIngest(v.UUID); err != nil {
					panic(err)
				}
				fmt.Printf("%s: %s cleared\n", k, v.Name)
			}
		}
	},
}
