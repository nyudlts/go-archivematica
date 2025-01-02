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
		completedIngests, err := amClient.GetCompletedIngests()
		if err != nil {
			t.Error(err)
		}

		if len(completedIngests.Results) > 0 {
			completedIngestId = completedIngests.Results[0]
		}
		t.Log(completedIngestId)

	})

	t.Run("Test Get Ingest Status", func(t *testing.T) {
		if completedIngestId != "" {
			completedIngest, err := amClient.GetIngestStatus(completedIngestId)
			if err != nil {
				t.Error(err)
			}
			t.Log(completedIngest)
		} else {
			t.Skip("cannot run test.")
		}

	})
}
