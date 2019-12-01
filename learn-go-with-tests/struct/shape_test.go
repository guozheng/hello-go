package _struct

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	tests := []struct{
		name string
		shape Shape
		want float64
	}{
		{"Rectangle", Rectangle{12.0, 6.0}, 72.0},
		{"Circle", Circle{10.0}, 314.1592653589793},
		{"Triangle", Triangle{12.0, 6.0}, 36.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := test.shape.Area()
			if got != test.want {
				t.Errorf("%#v got %.2f want %.2f", test.shape, got, test.want)
			}
		})
	}
}