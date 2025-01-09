package main

import (
	"fmt"
)

func produce(ch chan int) {
	for i := 0; i < 10; i++ {
		fmt.Println("Writing ", i)
		ch <- i
	}
	close(ch)
	fmt.Println("Closed the channel")
}
func main() {
	// ch := make(chan int, 3)
	ch := make(chan int)
	go produce(ch)

	for v := range ch {
		fmt.Println("Reading ", v)
		// time.Sleep(1 * time.Second)
		// println(v)
	}

	//	for {
	//		v, ok := <-ch
	//		if !ok {
	//			break
	//		}
	//		fmt.Println("Reading ", v)
	//		println(v)
	//	}
	fmt.Println("Got all!")
}
