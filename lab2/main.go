package main

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y float64
}

func crossProduct(a, b, c Point) float64 {
	return (b.X-a.X)*(c.Y-a.Y) - (b.Y-a.Y)*(c.X-a.X)
}

func getConvexHull(points []Point) []Point {
	n := len(points)
	if n <= 2 {
		return points
	}
	
	sort.Slice(points, func(i, j int) bool {
		if points[i].X == points[j].X {
			return points[i].Y < points[j].Y
		}
		return points[i].X < points[j].X
	})

	var hull []Point

	for _, p := range points {
		for len(hull) >= 2 && crossProduct(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	lowerLen := len(hull)
	for i := n - 2; i >= 0; i-- {
		p := points[i]
		for len(hull) > lowerLen && crossProduct(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	return hull[:len(hull)-1]
}

func HasNestedTrianglesFast(points []Point) bool {
	if len(points) < 6 {
		return false
	}

	hull := getConvexHull(points)
	
	if len(hull) < 3 {
		return false
	}

	hullMap := make(map[Point]bool)
	for _, p := range hull {
		hullMap[p] = true
	}

	var insidePoints []Point
	for _, p := range points {
		if !hullMap[p] {
			insidePoints = append(insidePoints, p)
		}
	}

	if len(insidePoints) < 3 {
		return false
	}

	for i := 0; i < len(insidePoints); i++ {
		for j := i + 1; j < len(insidePoints); j++ {
			for k := j + 1; k < len(insidePoints); k++ {
				if crossProduct(insidePoints[i], insidePoints[j], insidePoints[k]) != 0 {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	points := []Point{
		{0, 0}, {10, 0}, {5, 10},
		{4, 2}, {6, 2}, {5, 4},
	}

	if HasNestedTrianglesFast(points) {
		fmt.Println("Result: The set contains nested triangles.")
	} else {
		fmt.Println("Result: No nested triangles found.")
	}
}