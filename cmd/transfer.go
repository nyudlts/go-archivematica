package cmd

import (
	"bufio"
	"fmt"
	"go_am"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

var (
	directoryName string
	writer        *bufio.Writer
)

const locationName = "amatica rws ingest point"

func init() {
	transferCmd.Flags().StringVar(&config, "config", "", "")
	transferCmd.Flags().StringVar(&directoryName, "directory-name", "", "")
	rootCmd.AddCommand(transferCmd)
}

var transferCmd = &cobra.Command{
	Use: "transfer",
	Run: func(cmd *cobra.Command, args []string) {

		//create a client
		var err error
		client, err = go_am.NewAMClient(config, 20)
		if err != nil {
			panic(err)
		}

		//create an output file
		of, err := os.Create("output.tsv")
		if err != nil {
			panic(err)
		}
		defer of.Close()
		writer = bufio.NewWriter(of)

		//process the directory
		xfrDirs, err := os.ReadDir(directoryName)
		if err != nil {
			panic(err)
		}

		for _, xferDir := range xfrDirs {
			if strings.Contains(xferDir.Name(), "fales") {
				packageName := filepath.Join("test-dir", xferDir.Name())
				fmt.Println(packageName)
				transferPackage(packageName)
			}
		}

	},
}

func transferPackage(packageName string) {
	location, err := client.GetLocationByName(locationName)
	if err != nil {
		panic(err)
	}

	startTransferResponse, err := client.StartTransfer(location.UUID, packageName)
	if err != nil {
		panic(err)
	}

	uuid, err := startTransferResponse.GetUUID()
	if err != nil {
		panic(err)
	}

	fmt.Println("transfer started: " + uuid + packageName)

	foundUnapproved := false

	for !foundUnapproved {
		foundUnapproved = findUnapprovedTransfer(uuid)
		if !foundUnapproved {
			fmt.Println("  * Waiting for approval process to complete")
			time.Sleep(5 * time.Second)
		}
	}

	//approve the transfer

	transfer, err := client.GetTransferStatus(uuid)
	if err != nil {
		panic(err)
	}

	if err := client.ApproveTransfer(transfer.Directory, "standard"); err != nil {
		panic(err)
	}

	approvedTransfer, err := client.GetTransferStatus(uuid)
	if err != nil {
		panic(err)
	}

	fmt.Println("Transfer approved:", approvedTransfer.UUID)

	foundCompleted := false
	for !foundCompleted {
		foundCompleted = findCompletedTransfer(uuid)
		if !foundCompleted {
			fmt.Println("  * Waiting for transfer process to complete")
			time.Sleep(5 * time.Second)
		}
	}

	completedTransfer, err := client.GetTransferStatus(uuid)
	if err != nil {
		panic(err)
	}

	sipUUID := completedTransfer.SIPUUID

	fmt.Println("Transfer completed: ", sipUUID)

	foundIngestCompleted := false
	for !foundIngestCompleted {
		foundIngestCompleted = findCompletedIngest(sipUUID)
		if !foundIngestCompleted {
			fmt.Println("  * Waiting for ingest process to complete")
			time.Sleep(5 * time.Second)
		}
	}

	fmt.Println("Ingest Completed:", sipUUID)
	writer.WriteString(fmt.Sprintf("%s\t%s\n", sipUUID, packageName))
	writer.Flush()
}

func findCompletedIngest(sipUuid string) bool {
	completedIngests, err := client.GetCompletedIngests()
	if err != nil {
		panic(err)
	}

	completedIngestsMap, err := client.GetCompletedIngestsMap(completedIngests)
	if err != nil {
		panic(err)
	}

	for k, _ := range completedIngestsMap {
		if k == sipUuid {
			return true
		}
	}

	return false
}

func findCompletedTransfer(uuid string) bool {
	completedTransfers, err := client.GetCompletedTransfers()
	if err != nil {
		panic(err)
	}

	completedTransfersMap, err := client.GetCompletedTransfersMap(completedTransfers)
	if err != nil {
		panic(err)

	}

	for k, _ := range completedTransfersMap {
		if k == uuid {
			return true
		}
	}

	return false
}

func findUnapprovedTransfer(uuid string) bool {
	unapprovedTransfers, err := client.GetUnapprovedTransfers()
	if err != nil {
		panic(err)

	}

	unapprovedTransfersMap, err := client.GetUnapprovedTransfersMap(unapprovedTransfers)
	if err != nil {
		panic(err)
	}

	//find the unapproved transfer
	for k, _ := range unapprovedTransfersMap {
		if k == uuid {
			return true
		}
	}

	return false
}
