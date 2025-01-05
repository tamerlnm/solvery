package main

import (
	"fmt"
	"log"
	"solvery/pkg/geometry"
	"strconv"

	"github.com/spf13/pflag"
)

func main() {

	pointsFlag := pflag.StringArray("point", nil, "Set points")
	distanceFlag := pflag.Bool("distance", false, "Calculate distances")
	radiusFlag := pflag.String("radius", "", "Calculate that point in radius N")
	polygonFlag := pflag.StringArray("polygon", nil, "Set polygon points")
	pflag.Parse()

	if *distanceFlag {
		points, err := geometry.NewPointFromStringCoordinates(*pointsFlag, 2)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(points[0].GetDistance(points[1]))
	} else if *radiusFlag != "" {
		points, err := geometry.NewPointFromStringCoordinates(*pointsFlag, 1)
		if err != nil {
			log.Println(err)
			return
		}
		n, err := parseRadiusFlag(*radiusFlag)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(points[0].PointInRadius(n))
	} else if *polygonFlag != nil {
		points, err := geometry.NewPointFromStringCoordinates(*pointsFlag, 1)
		if err != nil {
			log.Println(err)
			return
		}
		polygon, err := geometry.NewPolygonFromStringCoordinates(*polygonFlag)
		if err != nil {
			log.Println(err)
			return
		}
		fmt.Println(polygon.PointInPolygon(points[0]))
	}
}

func parseRadiusFlag(radius string) (float64, error) {
	n, err := strconv.ParseFloat(radius, 64)
	if err != nil {
		return 0, fmt.Errorf("error to radius value: %s", radius)
	}

	if n <= 0 {
		return 0, fmt.Errorf("error to use negative value %s for raduis", radius)
	}
	return n, nil
}
