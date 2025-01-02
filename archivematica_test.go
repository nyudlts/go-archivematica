package go_archivematica

import (
	"testing"
)

func TestArchivematicaEndpoints(t *testing.T) {
	var amClient *AMClient
	var completedIngestId string

	amClient, err := NewAMClient(config, 20)
	if err != nil {
		t.Fatal(err.Error())
	}

	t.Run("Test Get Completed Ingests", func(t *testing.T) {
		_, err := amClient.GetCompletedIngests()
		if err != nil {
			t.Error(err)
		}
	})

	t.Run("Test Get Ingest Status", func(t *testing.T) {
		if completedIngestId != "" {
			_, err := amClient.GetIngestStatus(completedIngestId)
			if err != nil {
				t.Error(err)
			}
		} else {
			t.Skip("cannot run test.")
		}

	})
}
