package go_archivematica

import (
	"flag"
	"fmt"
	"testing"
)

var (
	config      string
	environment string
	client      *AMClient
)

func init() {
	flag.StringVar(&config, "config", "", "")
	flag.StringVar(&environment, "environment", "", "")
}

func TestAMClient(t *testing.T) {
	flag.Parse()
	t.Run("Test AMClient Creation", func(t *testing.T) {
		var err error
		client, err = NewAMClient(config, environment, 20)
		if err != nil {
			t.Error(err)
		}

		fmt.Println(client)

	})
}
