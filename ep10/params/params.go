package main

import "fmt"

// step 1
// func add(sum int, a int, b int) {
// 	sum = a + b
// 	fmt.Println("sum in add the function : ", sum) //5
// }
// func main() {
// 	sum := 0
// 	add(sum, 2, 3)
// 	fmt.Println("sum in the main function : ", sum) //0
// }

// step 2 & 3
// var sum int

// func add(a int, b int) {
// 	sum = a + b
// 	fmt.Println("sum in add the function : ", sum) //5
// }

// func main() {
// 	// sum := 0 //hides the global variable - comment this for step 3
// 	add(2, 3)
// 	fmt.Println("sum in the main function : ", sum) //prints the sum local to main 0
// }

// step 4

func add(sum *int, a int, b int) {
	*sum = a + b
	fmt.Println("sum in add the function : ", *sum)
}

func main() {
	sum := 0
	add(&sum, 2, 3)
	fmt.Println("sum in the main function : ", sum)
}
