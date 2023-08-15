package structsmethodinterfaces

import "testing"

func TestPerimeter(t *testing.T) {
	checkPerimeter := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		perimeter := shape.Perimeter()
		if perimeter != want {
			t.Errorf("Got %g, want %g", perimeter, want)
		}
	}
	t.Run("rectangles perimeter", func(t *testing.T) {
		rectangle := Rectangle{20.0, 10.0}
		want := 60.0
		checkPerimeter(t, rectangle, want)
	})
	t.Run("circles perimeter", func(t *testing.T) {
		circle := Circle{10.0}
		want := 62.83185307179586
		checkPerimeter(t, circle, want)
	})
}

func TestArea(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper() // Specify this is a helper function
		area := shape.Area()
		if want != area {
			t.Errorf("Got %g, want %g", area, want)
		}
	}
	t.Run("rectangles area", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		checkArea(t, rectangle, 72.0)
	})
	t.Run("circles area", func(t *testing.T) {
		circle := Circle{10.0}
		checkArea(t, circle, 314.1592653589793)
	})
}
