package go_am

import (
	"flag"
)

var config string

func init() {
	flag.StringVar(&config, "config", "", "")

}
