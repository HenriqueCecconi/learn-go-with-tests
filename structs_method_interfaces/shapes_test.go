package structsmethodinterfaces

import "testing"

func TestShape(t *testing.T) {
	shapeTests := []struct {
		name         string
		shape        Shape
		hasArea      float64
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{12, 6}, hasArea: 72.0, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{10.0}, hasArea: 314.1592653589793, hasPerimeter: 62.83185307179586},
		{name: "Triangle", shape: Triangle{12, 6}, hasArea: 36.0, hasPerimeter: 28.97056274847714},
	}

	for _, tt := range shapeTests {
		t.Run(tt.name+" area", func(t *testing.T) {
			areaGot := tt.shape.Area()
			if areaGot != tt.hasArea {
				t.Errorf("%#v area got: %g, area wanted: %g", tt.shape, areaGot, tt.hasArea)
			}
		})
		t.Run(tt.name+" perimeter", func(t *testing.T) {
			perimeterGot := tt.shape.Perimeter()
			if perimeterGot != tt.hasPerimeter {
				t.Errorf("%#v perimeter got: %g, perimeter wanted: %g", tt.shape, perimeterGot, tt.hasPerimeter)
			}
		})
	}
}
