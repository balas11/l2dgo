package main

import (
	"fmt"
	"sync"
	"time"
)

func produce1(c chan string) {
	for i := 0; i < 5; i++ {
		job := fmt.Sprintf("Prod 1 : job %d", i+1)
		fmt.Println("Sending job: ", job)
		c <- job
		time.Sleep(2 * time.Second)
	}
}

func produce2(c chan string) {
	for i := 0; i < 5; i++ {
		job := fmt.Sprintf("Prod 2 : job %d", i+1)
		fmt.Println("Sending job: ", job)
		c <- job
		time.Sleep(2 * time.Second)
	}
}

func consume(c1, c2 chan string) {
	for {
		select {
		case job := <-c1:
			fmt.Println("Received job: ", job)
			time.Sleep(3 * time.Second)
		case job := <-c2:
			fmt.Println("Received job: ", job)
			time.Sleep(2 * time.Second)
		case <-time.After(1 * time.Second):
			fmt.Println("Timeout")
			return
		}
	}
}

func main() {
	var wg sync.WaitGroup
	c1 := make(chan string, 2)
	c2 := make(chan string, 2)

	wg.Add(3)
	go func() { defer wg.Done(); produce1(c1) }()
	go func() { defer wg.Done(); produce2(c2) }()
	go func() { defer wg.Done(); consume(c1, c2) }()

	wg.Wait()
}
