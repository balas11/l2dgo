package main

import (
	"fmt"
	"math/rand"
	"time"
)

func makeDeck() []string {
	suites := []string{"♠️", "♥️", "♦️", "♣️"}
	values := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
	deck := make([]string, 0)
	for _, suite := range suites {
		for _, value := range values {
			deck = append(deck, value+suite)
		}
	}
	return deck
}

func shuffle(deck []string) {
	rand.Shuffle(len(deck), func(i, j int) { deck[i], deck[j] = deck[j], deck[i] })
}

func pickCard(deck []string, deal chan string, stop chan bool) {
	i := 0
	for {
		select {
		case <-stop:
			close(deal)
			return
		// case deal <- deck[i]:
		// 	i++
		// 	time.Sleep(400 * time.Millisecond)
		default:
			deal <- deck[i]
			i++
			time.Sleep(400 * time.Millisecond)
		}
	}
}

func main() {
	deck := makeDeck()
	shuffle(deck)
	deal := make(chan string)
	stop := make(chan bool)
	go pickCard(deck, deal, stop)
	for c := range deal {
		fmt.Println("Received ", c)
		if c == "A♠️" {
			close(stop)
		}
	}
}
