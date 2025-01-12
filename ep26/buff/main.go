package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// 	r := bufio.NewReader(os.Stdin)
// 	f, err := os.Create("test2.txt")
// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 		os.Exit(1)
// 	}
// 	defer f.Close()
// 	w := bufio.NewWriter(f)
// 	for {
// 		line, err := r.ReadString('\n')
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Printf("Error: %s\n", err)
// 		}
// 		w.WriteString(line)
// 	}
// 	w.Flush()

// }

// func main() {
// 	f, err := os.Open("uni.txt")

// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 		os.Exit(1)
// 	}
// 	defer f.Close()
// 	r := bufio.NewReader(f)
// 	for {
// 		r, size, err := r.ReadRune()
// 		if err == io.EOF {
// 			break
// 		}
// 		if err != nil {
// 			fmt.Printf("Error: %s\n", err)
// 		}
// 		fmt.Println(string(r), size)
// 	}
// }

// func main() {
// 	f, err := os.Open("test2.txt")

// 	if err != nil {
// 		fmt.Printf("Error: %s\n", err)
// 		os.Exit(1)
// 	}
// 	defer f.Close()
// 	scanner := bufio.NewScanner(f)

// 	for scanner.Scan() {
// 		fmt.Println(scanner.Text())
// 	}
// }

func main() {
	f, err := os.Open("test2.txt")
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
}
