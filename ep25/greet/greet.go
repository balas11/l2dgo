package main

import (
	"flag"
	"strings"
)

func main() {
	nameFlag := flag.String("name", "World", "Name to greet")

	var shoutFlag bool
	flag.BoolVar(&shoutFlag, "s", false, "Shout")

	flag.Parse()

	if shoutFlag {
		fmt.Println("HELLO " + strings.ToUpper(*nameFlag) + "!")
	} else {
		fmt.Println("Hello " + *nameFlag + "!")
	}
}
