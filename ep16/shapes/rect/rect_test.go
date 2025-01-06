package rect_test

import (
	"fmt"
	"shapes/rect"
	"testing"
)

var cases = []struct {
	length   int
	width    int
	area     int
	perimter int
	isSquare bool
}{
	{10, 5, 50, 30, false},
	{10, 10, 100, 40, true},
	{5, 5, 25, 20, true},
	{5, 10, 50, 30, false},
}

func TestNew(t *testing.T) {
	r := rect.New(10, 5)

	if r.Length() != 10 || r.Width() != 5 {
		t.Errorf("Expected length 10 and width 5, got %v and %v", r.Length(), r.Width())
	} else {
		t.Log("Test New Passed")
	}
}

func TestLength(t *testing.T) {
	r := rect.New(10, 5)

	if r.Length() != 10 {
		t.Errorf("Expected length 10, got %v", r.Length())
	}
}

func TestWidth(t *testing.T) {
	r := rect.New(10, 5)

	if r.Width() != 5 {
		t.Errorf("Expected width 5, got %v", r.Width())
	}
}

func TestSetLength(t *testing.T) {
	r := rect.New(10, 5)

	r.SetLength(15)
	if r.Length() != 15 {
		t.Errorf("Expected length 15, got %v", r.Length())
	}
}
func TestSetWidth(t *testing.T) {
	r := rect.New(10, 5)

	r.SetWidth(15)
	if r.Width() != 15 {
		t.Errorf("Expected width 15, got %v", r.Width())
	}
}

func TestPermiter(t *testing.T) {
	for _, c := range cases {
		r := rect.New(c.length, c.width)
		if r.Perimeter() != c.perimter {
			t.Errorf("Expected perimeter %v, got %v", c.perimter, r.Perimeter())
		}
	}
}

func TestArea(t *testing.T) {
	for _, c := range cases {
		r := rect.New(c.length, c.width)
		if r.Area() != c.area {
			t.Errorf("Expected area %v, got %v", c.area, r.Area())
		}
	}
}

func TestIsSquare(t *testing.T) {
	for _, c := range cases {
		r := rect.New(c.length, c.width)
		if r.IsSquare() != c.isSquare {
			t.Errorf("Expected isSquare %v, got %v", c.isSquare, r.IsSquare())
		}
	}
}

func TestScale(t *testing.T) {
	l, w := 10, 5
	name := "l=10,w=5"
	factors := []int{2, 3, 4}
	lengths := []int{20, 30, 40}
	widths := []int{10, 15, 20}
	for i, factor := range factors {
		t.Run(fmt.Sprintf("%s_f=%v", name, factor), func(t *testing.T) {
			r := rect.New(l, w)
			r.Scale(factor)
			if r.Length() != lengths[i] || r.Width() != widths[i] {
				t.Errorf("Expected length %v and width %v, got %v and %v", lengths[i], widths[i], r.Length(), r.Width())
			}
		})
	}
}
