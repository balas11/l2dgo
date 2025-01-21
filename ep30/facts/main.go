package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

func getNumberFact(n int) string {
	url := "http://numbersapi.com/" + strconv.Itoa(n)
	resp, err := http.Get(url)
	if err != nil {
		return "Got error while fetching the fact"
	}
	defer resp.Body.Close()
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return "Got error while reading the fact"
	}
	return string(content)
}

func getNumberFactForRange(start int, end int, out chan string) {
	for i := start; i <= end; i++ {
		out <- getNumberFact(i)
	}
}

func getFactsSequential() {
	start := time.Now()
	for i := 0; i < 100; i++ {
		result := getNumberFact(i + 1)
		fmt.Println(result)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)
}

func getFactsParallel() {
	fc := make(chan string, 11)
	start := time.Now()

	for i := 0; i < 10; i++ {
		start := i*10 + 1
		end := start + 9
		go func() {
			getNumberFactForRange(start, end, fc)
		}()
	}
	for i := 0; i < 100; i++ {
		result := <-fc
		fmt.Println(result)
	}
//	close(fc)

	elapsed := time.Since(start)
	fmt.Println(elapsed)
}
func main() {
	// getFactsSequential()
	getFactsParallel()
}
