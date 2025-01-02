package go_archivematica

import "flag"

var (
	config     string
	testClient *AMClient
)

func init() {
	flag.StringVar(&config, "config", "", "")
}
