package go_archivematica

import (
	"fmt"
	"testing"
)

func TestDashboard(t *testing.T) {
	t.Run("Test Get Completed Ingests", func(t *testing.T) {
		uuidList, err := client.GetCompletedIngests()
		if err != nil {
			t.Error(err)
		}

		fmt.Println(uuidList)
	})
}
