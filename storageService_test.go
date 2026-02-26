package go_archivematica

import (
	"fmt"
	"testing"
)

func TestStorageService(t *testing.T) {
	t.Run("TestGetEndpoints", func(t *testing.T) {
		endpoints, err := client.GetEndpoints()
		if err != nil {
			t.Errorf("Error getting API root: %v", err)
		}
		fmt.Printf("%+v\n", endpoints)
	})
}
