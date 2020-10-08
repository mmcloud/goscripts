package shapes

import "testing"

func TestArea(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10.0, 12.0}
		got := rectangle.Area()
		want := 120.0
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})
	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}

func TestAreaInterface(t *testing.T) {
	checkArea := func(t *testing.T, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{10, 12}
		checkArea(t, rectangle, 120.0)
	})
	t.Run("cirlces", func(t *testing.T) {
		cirlce := Circle{10}
		checkArea(t, cirlce, 314.1592653589793)
	})
}

func TestAreaInterfaceRefactored(t *testing.T) {
	areaTests := []struct {
		shape Shape
		want  float64
	}{
		{Rectangle{10, 12}, 120},
		{Circle{10}, 314.1592653589793},
	}
	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("got %g want %g", got, tt.want)
		}
	}
}
