package structsmethodinterfaces

import "testing"

func TestRectanglePerimeter(t *testing.T) {
	got := RectanglePerimeter(10.0, 10.0)
	want := 40.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}

func TestRectangleArea(t *testing.T) {
	got := RectangleArea(20.0, 10.0)
	want := 200.0

	if got != want {
		t.Errorf("Got %.2f, want %.2f", got, want)
	}
}
