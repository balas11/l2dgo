package main

import "fmt"

func main() {

	// s := "Hello"
	// s += " world!"
	// fmt.Println(s)

	// fmt.Println("alpha" > "Beta")

	// s1 := "Hello, world!"
	// fmt.Println(s1)

	// fmt.Printf("%c\n", s1[2])

	// fmt.Printf("%d\n", len(s1))
	// fmt.Println(s1[1:5])
	// fmt.Println(s1[:5])
	// fmt.Println(s1[7:])

	s2 := "abc⨷def"
	fmt.Printf("%d\n", len(s2))
	fmt.Println(s2[:3])
	fmt.Println(s2[6:])
	for i := 0; i < len(s2); i++ {
		fmt.Printf("%c\n", s2[i])
	}

	fmt.Println(s2)
	fmt.Printf("%s\n", s2)
	for idx, ch := range s2 {
		fmt.Printf("(%d, %v) ", idx, ch)
	}
	fmt.Println()
	for _, ch := range s2 {
		fmt.Printf("%c ", ch)
	}
	fmt.Println()

}
