package go_am

import "testing"

func TestAMClient(t *testing.T) {
	t.Run("Test AMClient Creation", func(t *testing.T) {
		amClient, err := NewAMClient("go-archivematica.yml", 20)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%s", amClient)

	})
}
