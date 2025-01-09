package main

import (
	"fmt"
)

func f(n int) {
	for i := 0; i < n; i++ {
		fmt.Println(i)
	}
}
func main() {
	go f(5)
	fmt.Println("called a goroutine")
	// time.Sleep(1 * time.Second)
	// fmt.Println("Press enter key to terminate...")
	// fmt.Scanln()
}
