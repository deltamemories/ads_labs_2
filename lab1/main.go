package main

import (
	"fmt"
	"sort"
)

type Point struct {
	X, Y float64
}

func CrossProduct(o, a, b Point) float64 {
	return (a.X-o.X)*(b.Y-o.Y) - (a.Y-o.Y)*(b.X-o.X)
}

func GetConvexHull(points []Point) []Point {
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

	hull := make([]Point, 0, 2*n)

	for _, p := range points {
		for len(hull) >= 2 && CrossProduct(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	lowerLen := len(hull)
	for i := n - 2; i >= 0; i-- {
		p := points[i]
		for len(hull) > lowerLen && CrossProduct(hull[len(hull)-2], hull[len(hull)-1], p) <= 0 {
			hull = hull[:len(hull)-1]
		}
		hull = append(hull, p)
	}

	return hull[:len(hull)-1]
}

func main() {
	var n int
	fmt.Print("Enter points count N: ")
	fmt.Scan(&n)

	if n < 3 {
		fmt.Println("For convex shell minimum 3 points required")
		return
	}

	points := make([]Point, n)
	fmt.Println("Enter x, y coords for each point:")
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].X, &points[i].Y)
	}

	hull := GetConvexHull(points)

	if len(hull) < 3 {
		fmt.Println("Convex shell does not exists (points is collinear)")
	} else {
		fmt.Println("Points of convex shell:")
		for _, p := range hull {
			fmt.Printf("(%.2f, %.2f)\n", p.X, p.Y)
		}
	}
}
