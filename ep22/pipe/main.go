package main

import (
	"fmt"
	"math/rand"
	"time"
)

func gen(done <-chan struct{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case out <- rand.Intn(10) + 1:
			case <-done:
				return
			}
		}
	}()
	return out
}

func triple(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for n := range in {
			select {
			case out <- n * 3:
			case <-done:
				return
			}
		}
	}()
	return out
}

func main() {

	done := make(chan struct{})
	defer close(done)

	out1 := gen(done)
	out2 := triple(done, out1)

	// for i := 0; i < 10; i++ {
	// 	fmt.Println(<-out2)
	// }

	timer1 := time.NewTimer(time.Second * 3)
	for {
		select {
		case n := <-out2:
			fmt.Println(n)
		case <-timer1.C:
			return
		}
	}
}
