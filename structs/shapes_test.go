package structs

import (
	"math"
	"testing"
)

func TestPerimeter(t *testing.T) {
	rec := Rectangle{10.0, 10.0}
	got := Perimeter(rec)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape AreaCalculator, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}

	}
	t.Run("test Rectangle area", func(t *testing.T) {
		rec := Rectangle{5.0, 8.0}
		want := 40.0
		checkArea(t, rec, want)
	})
	t.Run("test circle area", func(t *testing.T) {
		c := Circle{10}
		want := math.Pi * 100
		checkArea(t, c, want)
	})
}

func TestArea2(t *testing.T) {
	type TestRow struct {
		name  string
		shape AreaCalculator
		want  float64
	}
	areaTests := []TestRow{
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, math.Pi * 10 * 10},
		{"Triangle", Triangle{12, 6}, 36.0},
	}

	for _, tt := range areaTests {
		got := tt.shape.Area()
		if got != tt.want {
			t.Errorf("%#v got %g want %g", tt.name, got, tt.want)
		}
	}
}
