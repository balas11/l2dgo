package main

import "fmt"

func main() {
	var name string
	fmt.Print("Enter a name : ")
	fmt.Scanf("%s", &name)

	if len(name) <= 3 {
		fmt.Println(name, " is a short name!")
	} else if len(name) <= 6 {
		fmt.Println(name, " is a medium name!")
	} else {
		fmt.Println(name, " is a long name!")
	}
}
