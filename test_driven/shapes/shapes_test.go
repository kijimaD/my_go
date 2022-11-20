package shapes

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

// shapeはinterfaceで、異なる実装を入れられる。これによってテストしやすくなる
func TestArea(t *testing.T) {
	areaTests := []struct {
		name    string // テストケースの名前として使用する
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0}, // フィールド名をつけるとわかりやすい
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
	}

	for _, tt := range areaTests {
		// t.Runで名前をつけることで、失敗したときにどこで失敗しているのかわかりやすい
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
