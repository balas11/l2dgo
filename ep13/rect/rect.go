package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func perimiter(rect Rectangle) int {
	return 2 * (rect.length + rect.width)
}

// func scale(rect Rectangle, factor int) {
// 	rect.length *= factor
// 	rect.width *= factor
// }

func scale(rect *Rectangle, factor int) {
	rect.length *= factor
	rect.width *= factor
}

func main() {
	var r1 Rectangle
	r1.length = 10
	r1.width = 5
	fmt.Println(r1.length, r1.width)

	r2 := Rectangle{length: 15, width: 10}
	fmt.Println(r2.length, r2.width)

	r3 := Rectangle{12, 10}
	fmt.Println(r3.length, r3.width)

	r4 := new(Rectangle)
	// (*r4).length = 20
	// (*r4).width = 15
	r4.length = 20
	r4.width = 15
	fmt.Println(r4.length, r4.width)

	p1 := perimiter(r3)
	fmt.Println(p1)
	p2 := perimiter(*r4)
	fmt.Println(p2)

	scale(&r3, 2)
	fmt.Println(r3.length, r3.width)
}
