package main

import "fmt"

func f1() {
	fmt.Println("Executing f1")
}

func f2() {
	fmt.Println("Executing f2")
}

func f3() {
	fmt.Println("Executing f3")
}

// func main() {
// 	defer f1()
// 	f2()
// 	f3()
// }

func main() {
	defer f1()
	defer f2()
	f3()
}

// func main() {
// 	defer fmt.Println("Done with main")
// 	defer f1()
// 	defer f2()
// 	f3()
// }
