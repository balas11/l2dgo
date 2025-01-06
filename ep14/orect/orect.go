package main

import "fmt"

type Rectangle struct {
	length int
	width  int
}

func (r Rectangle) area() int {
	return r.length * r.width
}

func (r Rectangle) perimeter() int {
	return 2 * (r.length + r.width)
}

func (r *Rectangle) scale(factor int) {
	r.length = r.length * factor
	r.width = r.width * factor
}

func main() {

	r1 := Rectangle{length: 10, width: 5}
	fmt.Println("Dimension : ", r1.length, r1.width)
	fmt.Println("Area : ", r1.area())
	fmt.Println("Perimeter : ", r1.perimeter())
	r1.scale(2)
	fmt.Println("Dimension : ", r1.length, r1.width)
}
