package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	client := http.Client{}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Millisecond)
	defer cancel()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, "https://gobyexample.com", nil)
	if err != nil {
		log.Println("New request - ", err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Do request - ", err)
	} else {
		defer resp.Body.Close()
		fmt.Println("Status code - ", resp.StatusCode)
		content, _ := io.ReadAll(resp.Body)
		fmt.Println(string(content))
	}
}
