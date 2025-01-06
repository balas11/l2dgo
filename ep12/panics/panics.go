package main

import "fmt"

func main() {
	defer func() {
		message := recover()
		fmt.Println(message)
	}()

	fmt.Println("Starting to execute...")
	fmt.Println("Something happend, going to panic...")
	panic("Something went wrong!")
	fmt.Println("You are not going to see this")
}
