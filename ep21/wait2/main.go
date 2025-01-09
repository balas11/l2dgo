package main

import (
	"fmt"
	"time"
)

func f(n int, done chan bool) {
	for i := 0; i < n; i++ {
		fmt.Println("f:", i)
		time.Sleep(2 * time.Millisecond)
	}
	done <- true
}

func main() {
	done := make(chan bool)

	go f(100, done)

	<-done
}
