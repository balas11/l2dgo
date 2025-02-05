package main

import (
	"context"
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
	ctx, cancel := context.WithCancel(context.Background())
	go generate(ctx, rnc)

	for n := range rnc {
		fmt.Println(n)
		if n == 25 {
			cancel()
		}
	}
}
