package main

import (
	"fmt"
	"time"
)

func doSomething(out chan string) {
	out <- "hello"
	fmt.Println("Sent something...")
}

func main() {
	ch := make(chan string)
	go doSomething(ch)
	fmt.Println("Waiting...")
	time.Sleep(3 * time.Second)
	fmt.Println("Going to read")
	fmt.Println(<-ch)
	time.Sleep(1 * time.Second)
}
