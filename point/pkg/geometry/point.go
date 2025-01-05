package geometry

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	X, Y float64
}

type Polygon struct {
	Vertices []Point
}

func NewPointFromStringCoordinates(slice []string, n int) ([]Point, error) {
	var res []float64

	for i := 0; i < len(slice); i++ {
		subSlice := strings.Split(slice[i], ",")
		for j := 0; j < len(subSlice); j++ {
			el, err := strconv.ParseFloat(subSlice[j], 64)
			if err != nil {
				return []Point{}, fmt.Errorf("error to parse point value: %s", slice[i])
			}
			res = append(res, el)
		}
	}

	if n == 1 && len(res) < 2 {
		return nil, fmt.Errorf("not enough coordinates for 1 point")
	}
	if n == 2 && len(res) < 4 {
		return nil, fmt.Errorf("not enough coordinates for 2 points")
	}

	points := make([]Point, n)
	for i := 0; i < n; i++ {
		points[i] = Point{res[i*2], res[i*2+1]}
	}

	return points, nil
}

func NewPolygonFromStringCoordinates(slice []string) (Polygon, error) {
	if len(slice) <= 2 {
		return Polygon{}, errors.New("error: not enough value")
	}
	var vertices []Point

	for _, v := range slice {
		subSlice := strings.Split(v, ",")

		x, err := strconv.ParseFloat(subSlice[0], 64)
		if err != nil {
			return Polygon{}, fmt.Errorf("error to parse polygon point=%s value=%f", subSlice, x)
		}

		y, err := strconv.ParseFloat(subSlice[1], 64)
		if err != nil {
			return Polygon{}, fmt.Errorf("error to parse polygon point=%s value=%f", subSlice, y)
		}
		vertices = append(vertices, Point{x, y})

	}

	return Polygon{vertices}, nil
}

func (p *Point) GetDistance(pSecond Point) float64 {
	return math.Sqrt(math.Pow(p.X-pSecond.X, 2) + math.Pow(p.Y-pSecond.Y, 2))
}

func (p *Point) PointInRadius(radius float64) bool {
	return math.Sqrt(p.X*p.X+p.Y*p.Y) <= radius
}

func (pg *Polygon) PointInPolygon(p Point) bool {
	n := len(pg.Vertices)
	crossing := false
	for i := 0; i < n; i++ {
		v1 := pg.Vertices[i]
		v2 := pg.Vertices[(i+1)%n]

		if p.Y <= math.Min(v1.Y, v2.Y) || p.Y > math.Max(v1.Y, v2.Y) {
			continue
		}
		if p.X > math.Max(v1.X, v2.X) {
			continue
		}

		if v1.Y == v2.Y {
			continue
		}
		xt := (p.Y-v1.Y)*(v2.X-v1.X)/(v2.Y-v1.Y) + v1.X
		if v1.X == v2.X || p.X <= xt {
			crossing = !crossing
		}

	}

	return crossing
}
