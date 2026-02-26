package go_archivematica

import (
	"fmt"
	"testing"
)

func TestSpaces(t *testing.T) {
	var spaceUUID string
	t.Run("TestGetSpaces", func(t *testing.T) {
		spaces, err := client.GetSpaces()
		if err != nil {
			t.Errorf("Error getting spaces: %v", err)
		}
		fmt.Println(spaces)
		spaceUUID = spaces.Objects[0].UUID.String()
	})

	t.Run("TestGetSpace", func(t *testing.T) {
		space, err := client.GetSpace(spaceUUID)
		if err != nil {
			t.Errorf("Error getting space: %v", err)
		}
		fmt.Println(space)
	})

	t.Run("TestBrowseSpace", func(t *testing.T) {
		spaceBrowser, err := client.BrowseSpace(spaceUUID)
		if err != nil {
			t.Errorf("Error browsing space: %v", err)
		}
		fmt.Println(spaceBrowser)
	})

	/*
		t.Run("TestCreateSpace", func(t *testing.T) {
			createSpaceRequest := CreateSpaceRequest{
				AccessProtocol: "FS",
				Path:           "/var/archivematica/testspace",
				StagingPath:    "/",
			}
			err := client.CreateSpace(createSpaceRequest)
			if err != nil {
				t.Errorf("Error creating space: %v", err)
			}
		})
	*/
}
