package main

import (
	"fmt"

	go_archivematica "github.com/nyudlts/go-archivematica"
)

func main() {
	config := "go-archivematica.json"
	client, err := go_archivematica.NewAMClient(config, "806", 20)
	if err != nil {
		panic(err)
	}

	fmt.Println(client.AIPStoreLocation)

	/*
		spaces, err := client.GetSpaces()
		if err != nil {
			panic(err)
		}
		if len(spaces.Objects) > 0 {
			spaceUUID := spaces.Objects[0].UUID.String()
			fmt.Println(spaceUUID)
		} else {
			panic("No spaces returned")
		}
	*/
}
