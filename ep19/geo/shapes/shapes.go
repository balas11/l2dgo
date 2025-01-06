package shapes

import "fmt"

type Rectangle struct {
	Length int
	Width  int
}

func (r Rectangle) Area() int {
	return r.Length * r.Width
}

func (r Rectangle) String() string {
	return fmt.Sprintf("Length: %d, Width: %d, Area: %d", r.Length, r.Width, r.Area())
}

type Rectangles []Rectangle

func (r Rectangles) Len() int {
	return len(r)
}

func (r Rectangles) Less(i, j int) bool {
	return r[i].Area() < r[j].Area()
}

func (r Rectangles) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}
