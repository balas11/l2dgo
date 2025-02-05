package main

import (
	"context"
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func makeRandomNumbers() []int {
	s := make([]int, 0, 50)
	for i := 0; i < 50; i++ {
		s = append(s, i)
	}

	rand.Shuffle(len(s), func(i, j int) {
		s[i], s[j] = s[j], s[i]
	})
	return s
}

func generate(ctx context.Context, rnc chan int) {
	ns := makeRandomNumbers()
	i := 0
	for {
		select {
		case <-ctx.Done():
			close(rnc)
			return
		default:
			rnc <- ns[i]
			i++
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func main() {
	rnc := make(chan int)
	// ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	// ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(1*time.Second))
	ctx, cancel := context.WithTimeoutCause(context.Background(), 2*time.Second, errors.New("Timed out in 2 seconds"))
	defer cancel()
	go generate(ctx, rnc)

	for {
		cancelled := false
		select {
		case n := <-rnc:
			fmt.Println(n)
		case <-ctx.Done():
			fmt.Println("done")
			cancelled = true
		}

		if cancelled {
			break
		}
	}
	fmt.Println(context.Cause(ctx).Error())
	fmt.Println(ctx.Err().Error())
}
