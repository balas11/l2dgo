package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	resp, err := http.Get("https://gobyexample.com")
	// resp, err := http.Get("http://numbersapi.com")
	// resp, err := http.Get("http://numbersapi.com/11")
	// resp, err := http.Get("http://numbersapi.com/11?json")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	for k, v := range resp.Header {
		fmt.Println(k, ": ", v)
	}
	fmt.Println("Press return key")
	fmt.Scanln()
	content, _ := io.ReadAll(resp.Body)
	fmt.Println(string(content))

}

// func main() {
// 	client := &http.Client{Timeout: 10 * time.Second}
// 	req, err := http.NewRequest(http.MethodGet, "https://gobyexample.com", nil)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	resp, err := client.Do(req)
// 	if err != nil {
// 		log.Fatalln(err)
// 	}
// 	defer resp.Body.Close()

// 	content, _ := io.ReadAll(resp.Body)
// 	log.Println(string(content))
// }
