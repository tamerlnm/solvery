package main

import (
	"testing"
)

func TestGetDistance(t *testing.T) {
	point1 := Point{0, 0}
	point2 := Point{3, 4}

	var expectedDistance float64 = 5
	distance := point1.getDistance(point2)

	if distance != expectedDistance {
		t.Errorf("Expected distance to be %.2f but got %.2f", expectedDistance, distance)
	}
}

func TestInRadius(t *testing.T) {
	point := Point{0, 0}
	firstPoint := Point{1, 1}
	secondPoint := Point{5, 5}
	radius := 2.5

	if !point.PointInRadius(firstPoint, radius) {
		t.Errorf("Expected point %v to be inside radius %v", firstPoint, radius)
	}

	if point.PointInRadius(secondPoint, radius) {
		t.Errorf("Expected point %v to be inside radius %v", secondPoint, radius)
	}
}

func TestPointInPolygon(t *testing.T) {
	polygon := Polygon{
		vertices: []Point{
			{x: 0, y: 0},
			{x: 4, y: 0},
			{x: 4, y: 4},
			{x: 0, y: 4},
		},
	}

	firstPoint := Point{2, 2}
	secondPoint := Point{5, 5}

	if !polygon.PointInPolygon(firstPoint) {
		t.Errorf("Expected point %v to be inside the polygon", firstPoint)
	}

	if polygon.PointInPolygon(secondPoint) {
		t.Errorf("Expected point %v to be inside the polygon", secondPoint)
	}
}
