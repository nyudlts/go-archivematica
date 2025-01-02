package go_archivematica

import (
	"testing"
)

var client *AMClient

func TestCommon(t *testing.T) {
	var err error
	client, err = NewAMClient(config, 20)
	if err != nil {
		t.Error(err)
	}

	t.Run("test convert uuid string to am directory string", func(t *testing.T) {
		uid := "d2e28d83-28fc-4d3f-b9ac-2a2e946640dc"
		_, err := ConvertUUIDToAMDirectory(uid)
		if err != nil {
			t.Error(err)
		}

	})

	/*
		t.Run("test if uuid dir exists", func(t *testing.T) {
			sipID := "d31b1e0d-ee32-4910-977f-bf00d6f5f0a3"
			_, err := client.GetAIPLocation(sipID)
			if err != nil {
				t.Error(err)
			}
		})
	*/

}
