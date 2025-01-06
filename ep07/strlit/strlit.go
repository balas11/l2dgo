package main

import "fmt"

func main() {
	s1 := "hello"
	s2 := "This is\na multi line\ntext"
	s3 := "Strings are enclosed in \" and \""
	s4 := "include a \\n to break a line"
	s5 := `He said, "Hello!"`
	s6 := `include a \n to break a line`
	s7 := `Line 1
Line 2
Line 3
`

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
	fmt.Println(s4)
	fmt.Println(s5)
	fmt.Println(s6)
	fmt.Println(s7)
}
