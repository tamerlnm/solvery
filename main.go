package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"log"
	"math"
	"strconv"
	"strings"
)

type Point struct {
	x, y float64
}

type Polygon struct {
	vertices []Point
}

func (p *Point) getDistance(pSecond Point) float64 {
	return math.Sqrt(math.Pow(p.x-pSecond.x, 2) + math.Pow(p.y-pSecond.y, 2))
}

func (p *Point) PointInRadius(pSecond Point, radius float64) bool {
	distance := p.getDistance(pSecond)
	return distance <= radius
}

func (pg *Polygon) PointInPolygon(p Point) bool {
	n := len(pg.vertices)
	crossing := false
	for i := 0; i < n; i++ {
		v1 := pg.vertices[i]
		v2 := pg.vertices[(i+1)%n]

		if p.y > math.Min(v1.y, v2.y) && p.y <= math.Max(v1.y, v2.y) {
			if p.x <= math.Max(v1.x, v2.x) {
				if v1.y != v2.y {
					xt := (p.y-v1.y)*(v2.x-v1.x)/(v2.y-v1.y) + v1.x
					if v1.x == v2.x || p.x <= xt {
						crossing = !crossing
					}
				}
			}
		}
	}

	return crossing
}

func main() {

	points := pflag.StringArray("point", nil, "Set points")
	distance := pflag.Bool("distance", false, "Calculate distances")
	radius := pflag.String("radius", "", "Calculate that point in radius N")
	polygonFlag := pflag.String("polygon", "", "Set polygon points")
	pflag.Parse()

	firstPoint, secondPoint, err := setPoints(*points)
	if err != nil {
		log.Println(err)
		return
	}

	if *distance {
		fmt.Println(firstPoint.getDistance(secondPoint))
	}

	if *radius != "" {
		n, err := parseRadiusFlag(*radius)
		if err != nil {
			log.Println(err)
			return
		}
		p := Point{0, 0}
		fmt.Printf("first point in radius=%.2f: %t\n", n, p.PointInRadius(firstPoint, n))
		fmt.Printf("second point in radius=%.2f: %t\n", n, p.PointInRadius(secondPoint, n))
	}

	if *polygonFlag != "" {
		polygon, err := parsePolygon(*polygonFlag)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println("first point in Polygon:", polygon.PointInPolygon(firstPoint))
		fmt.Println("second point in Polygon:", polygon.PointInPolygon(secondPoint))
	}
}

func parsePolygon(str string) (Polygon, error) {
	slice := strings.Split(str, ",")
	if len(slice) <= 5 {
		return Polygon{}, fmt.Errorf("error, not enough values: %s", str)
	}
	if len(slice)%2 != 0 {
		return Polygon{}, fmt.Errorf("error, wrong quantity of values: %s", str)
	}
	var vertices []Point
	for i := 0; i < len(slice); i += 2 {
		x, err1 := strconv.ParseFloat(slice[i], 64)
		y, err2 := strconv.ParseFloat(slice[i+1], 64)

		if err1 != nil || err2 != nil {
			return Polygon{}, fmt.Errorf("Error converting coordinates '%s' and '%s' to float64: %v %v\n", slice[i], slice[i+1], err1, err2)
		}

		vertices = append(vertices, Point{x: x, y: y})
	}
	return Polygon{vertices}, nil
}

func parseRadiusFlag(radius string) (float64, error) {
	n, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		return 0, fmt.Errorf("error to parse point value: %s", radius)
	}
	return n, nil
}

func parseFlagPoint(slice []string) ([]float64, error) {

	var res []float64

	for i := 0; i < len(slice); i++ {
		subSlice := strings.Split(slice[i], ",")
		for j := 0; j < len(subSlice); j++ {
			el, err := strconv.ParseFloat(subSlice[j], 64)
			if err != nil {
				return nil, fmt.Errorf("error to parse point value: %s", slice[i])
			}
			res = append(res, el)
		}
	}
	return res, nil
}

func setPoints(slice []string) (Point, Point, error) {
	res, err := parseFlagPoint(slice)
	if err != nil {
		return Point{}, Point{}, err
	}

	return Point{res[0], res[1]}, Point{res[2], res[3]}, nil
}
