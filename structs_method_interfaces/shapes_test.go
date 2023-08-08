package structsmethodinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	t.Run("rectangles perimeter", func(t *testing.T) {
		rectangle := Rectangle{20.0, 10.0}
		got := rectangle.Perimeter()
		want := 60.0

		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	})
	t.Run("circles perimeter", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Perimeter()
		want := 62.83185307179586

		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {
	t.Run("rectangles area", func(t *testing.T) {
		rectangle := Rectangle{20.0, 10.0}
		got := rectangle.Area()
		want := 200.0

		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	})
	t.Run("circles area", func(t *testing.T) {
		circle := Circle{10.0}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("Got %g, want %g", got, want)
		}
	})
}
