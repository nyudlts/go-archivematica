package go_am

import (
	"flag"
	"testing"
)

func TestAMClient(t *testing.T) {
	flag.Parse()
	t.Run("Test AMClient Creation", func(t *testing.T) {
		amClient, err := NewAMClient(config, 20)
		if err != nil {
			t.Error(err)
		}
		t.Logf("%s", amClient)

	})
}
