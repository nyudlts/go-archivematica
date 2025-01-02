package go_archivematica

import (
	"flag"
	"testing"
)

var config string

func init() {
	flag.StringVar(&config, "config", "", "")
}

func TestAMClient(t *testing.T) {
	flag.Parse()
	t.Run("Test AMClient Creation", func(t *testing.T) {
		_, err := NewAMClient(config, 20)
		if err != nil {
			t.Error(err)
		}

	})
}
