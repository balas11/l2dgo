package rect

type Rectangle struct {
	length int
	width  int
}

func New(length int, width int) *Rectangle {
	return &Rectangle{length, width}
}

func (r Rectangle) Length() int {
	return r.length
}

func (r Rectangle) Width() int {
	return r.width
}

func (r *Rectangle) SetLength(length int) {
	r.length = length
}

func (r *Rectangle) SetWidth(width int) {
	r.width = width
}

func (r Rectangle) Area() int {
	return r.length * r.width
}

func (r Rectangle) Perimeter() int {
	return 2 * (r.length + r.width)
}

func (r Rectangle) IsSquare() bool {
	return r.length == r.width
}

func (r *Rectangle) Scale(factor int) {
	r.length *= factor
	r.width *= factor
}
