package main

import (
	"flag"
	"fmt"
	"strings"
)

// func main() {
// 	len := len(os.Args)
// 	if len > 1 {
// 		text := strings.Join(os.Args[1:], " ")
// 		fmt.Print(text)
// 	}
// 	fmt.Println()
// }

func main() {
	shoutFlag := flag.Bool("s", false, "shout")
	flag.Parse()
	fmt.Println("flag:", *shoutFlag)
	fmt.Println("flag args:", flag.Args())
	len := len(flag.Args())
	if len >= 1 {
		text := strings.Join(flag.Args(), " ")
		if *shoutFlag {
			text = strings.ToUpper(text)
		}
		fmt.Print(text)
	}
	fmt.Println()
}
