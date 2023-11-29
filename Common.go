package go_am

import (
	"flag"
	"regexp"
)

var config string

func init() {
	flag.StringVar(&config, "config", "", "")

}

var uuidPtn = regexp.MustCompile("[a-fA-F0-9]{8}-[a-fA-F0-9]{4}-4[a-fA-F0-9]{3}-[8|9|aA|bB][a-fA-F0-9]{3}-[a-fA-F0-9]{12}")
