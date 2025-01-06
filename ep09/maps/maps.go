package main

import (
	"fmt"
)

func main() {
	var langs map[string]int

	langs = make(map[string]int)

	langs["Go"] = 1
	langs["Elixir"] = 2
	langs["Python"] = 3

	fmt.Println(langs)
	fmt.Println(len(langs))
	fmt.Println(langs["Go"])

	v := langs["Javascript"]
	fmt.Println(v)

	val, found := langs["Javascript"]
	if !found {
		fmt.Println("Javascript found ", found, val)
	}

	if _, found := langs["Go"]; found {
		fmt.Println("Go is present : ", found)
	} else {
		fmt.Println("Go is not present!")
	}

	delete(langs, "Python")
	fmt.Println(langs)
	fmt.Println(len(langs))

}
