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
	t.Run("triangles perimeter", func(t *testing.T) {
		triangle := Triangle{12.0, 6}
		want := 28.97056274847714
		checkPerimeter(t, triangle, want)
	})
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{12, 6}, 72.0},
		{Circle{10.0}, 314.1592653589793},
		{Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("Got %g want %g", got, tt.want)
		}
	}
}
