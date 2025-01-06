package main

import "fmt"

// Step 1
// func main() {
// 	name := "John"

// 	l := len(name)

// 	if l <= 3 {
// 		fmt.Println(name, " is a short name")
// 	} else {
// 		fmt.Println(name, " is a long name")
// 	}
// }

// Step 2
// func main() {
// 	name := "John"

// 	if l := len(name); l <= 3 {
// 		fmt.Println(name, " is a short name")
// 	} else {
// 		fmt.Println(name, " is a long name")
// 	}
// }

// step 3
// func main() {
// 	name := "John"

// 	if len(name) <= 3 {
// 		fmt.Println(name, " is a short name")
// 	} else {
// 		fmt.Println(name, " is a long name")
// 	}
// }

func main() {
	name := "John"

	if len(name) <= 3 {
		fmt.Println(name, " is a short name")
	} else if len(name) <= 5 {
		fmt.Println(name, " is a medium name")
	} else {
		fmt.Println(name, " is a long name")
	}
}
