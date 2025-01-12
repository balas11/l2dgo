package main

import (
	"os"
)

// func main() {
// 	f, err := os.Create("test.txt")

// 	if err != nil {
// 		panic(err)
// 	}

// 	defer f.Close()

// 	f.WriteString("hello, file!\n")
// 	f.WriteString("Line 2 of this file\n")
// 	f.WriteString("Line 3 of this file\n")
// 	f.WriteString("End of the content\n")
// }

// func main() {
// 	content := `This is the first line
// This is the second line
// This is the third line
// `
// 	err := os.WriteFile("test1.txt", []byte(content), 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func main() {
// 	content, err := os.ReadFile("test.txt")

// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Print(string(content))
// 	// os.Stdout.Write(content)
// }

// func main() {
// 	f, err := os.Open("test.txt")

// 	if err != nil {
// 		panic(err)
// 	}
// 	defer f.Close()

// 	buf := make([]byte, 16)

// 	for {
// 		n, err := f.Read(buf)
// 		if err == io.EOF {
// 			break
// 		}
// 		fmt.Print(string(buf[:n]))
// 		// os.Stdout.Write(buf[:n])
// 	}
// }

func main() {
	f, err := os.OpenFile("test1.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

	if err != nil {
		panic(err)
	}
	defer f.Close()

	f.WriteString("Appended Line 1\n")
	f.WriteString("Appended Line 2\n")

}
