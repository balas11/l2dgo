package main

import "fmt"

//step 1
// func isEven(n int) bool {
// 	if n%2 == 0 {
// 		return true
// 	} else {
// 	return false
//  }
// }

// step 2 - drop else
// func isEven(n int) bool {
// 	if n%2 == 0 {
// 		return true
// 	}
// 	return false
// }

// step 3 - drop if
func isEven(n int) bool {
	return n%2 == 0
}

func main() {
	fmt.Println(isEven(10))
	fmt.Println(isEven(11))
}
