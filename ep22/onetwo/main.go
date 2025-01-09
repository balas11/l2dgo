package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

func produce(c chan string) {
	for i := 0; i < 5; i++ {
		job := "job" + strconv.Itoa(i+1)
		fmt.Println("created : ", job)
		c <- job
		time.Sleep(1 * time.Second)
	}
	close(c)
}

func consume1(c chan string) {
	for job := range c {
		fmt.Println("Consumer 1 processing : ", job)
		time.Sleep(2 * time.Second)
	}
}

func consume2(c chan string) {
	for job := range c {
		fmt.Println("Consumer 2processing : ", job)
		time.Sleep(3 * time.Second)
	}
}

func main() {
	var wg sync.WaitGroup
	ch := make(chan string, 2)

	wg.Add(3)
	go func() { defer wg.Done(); produce(ch) }()
	go func() { defer wg.Done(); consume1(ch) }()
	go func() { defer wg.Done(); consume2(ch) }()
	wg.Wait()
}
