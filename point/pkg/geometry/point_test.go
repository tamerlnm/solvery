package geometry

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	point1 := Point{0, 0}
	point2 := Point{3, 4}

	var expectedDistance float64 = 5
	distance := point1.GetDistance(point2)

	if distance != expectedDistance {
		t.Errorf("Expected distance to be %.2f but got %.2f", expectedDistance, distance)
	}
}

func TestInRadius(t *testing.T) {
	point := Point{1, 1}
	radius := 2.5

	if !point.PointInRadius(radius) {
		t.Errorf("Expected point %v to be inside radius %v", point, radius)
	}
}

func TestPointInPolygon(t *testing.T) {
	polygon := Polygon{
		Vertices: []Point{
			{X: 0, Y: 0},
			{X: 4, Y: 0},
			{X: 4, Y: 4},
			{X: 0, Y: 4},
		},
	}

	point := Point{2, 2}

	if !polygon.PointInPolygon(point) {
		t.Errorf("Expected point %v to be inside the polygon", point)
	}
}
