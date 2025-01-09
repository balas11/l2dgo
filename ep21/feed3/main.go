package main

import (
	"fmt"
	"sync"
	"time"
)

func feedMonkey(c chan string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1 * time.Millisecond)
		fmt.Println("Feeding Monkey with Banana #", i+1)
		c <- fmt.Sprintf("Banana #%d", i+1)
	}
}

func monkeyEat(c chan string) {
	for i := 0; i < 5; i++ {
		food := <-c
		fmt.Println("Monkey eating ", food)
		time.Sleep(3 * time.Second)
	}
}
func main() {
	c := make(chan string)
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		feedMonkey(c)
	}()
	go func() {
		defer wg.Done()
		monkeyEat(c)
	}()
	wg.Wait()
}
