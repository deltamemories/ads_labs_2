package main

import (
	"fmt"
	"math"
)

// Point represents a 2D coordinate.
type Point struct {
	X, Y float64
}

// Triangle represents a triangle defined by three points.
type Triangle struct {
	A, B, C Point
}

// isCollinear checks if three points are collinear.
func isCollinear(p1, p2, p3 Point) bool {
	area := p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y)
	return math.Abs(area) < 1e-9
}

// sign calculates the cross product for point-line orientation.
func sign(p1, p2, p3 Point) float64 {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}

// isPointStrictlyInside checks if a point is strictly inside a triangle.
func isPointStrictlyInside(pt, v1, v2, v3 Point) bool {
	d1 := sign(pt, v1, v2)
	d2 := sign(pt, v2, v3)
	d3 := sign(pt, v3, v1)

	hasNeg := (d1 < -1e-9) || (d2 < -1e-9) || (d3 < -1e-9)
	hasPos := (d1 > 1e-9) || (d2 > 1e-9) || (d3 > 1e-9)

	// Points on the boundary are not considered strictly inside.
	isZero := math.Abs(d1) <= 1e-9 || math.Abs(d2) <= 1e-9 || math.Abs(d3) <= 1e-9
	if isZero {
		return false
	}

	return !(hasNeg && hasPos)
}

// isTriangleNested checks if one triangle is strictly nested within another.
func isTriangleNested(inner, outer Triangle) bool {
	return isPointStrictlyInside(inner.A, outer.A, outer.B, outer.C) &&
		isPointStrictlyInside(inner.B, outer.A, outer.B, outer.C) &&
		isPointStrictlyInside(inner.C, outer.A, outer.B, outer.C)
}

// HasNestedTriangles checks if any triangle in the set is nested within another.
func HasNestedTriangles(points []Point) bool {
	n := len(points)
	if n < 6 {
		return false
	}

	var triangles []Triangle

	// Find all valid triangles.
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				if !isCollinear(points[i], points[j], points[k]) {
					triangles = append(triangles, Triangle{points[i], points[j], points[k]})
				}
			}
		}
	}

	// Check for nested pairs.
	for i := 0; i < len(triangles); i++ {
		for j := 0; j < len(triangles); j++ {
			if i == j {
				continue
			}
			if isTriangleNested(triangles[i], triangles[j]) {
				return true
			}
		}
	}

	return false
}

func main() {
	points := []Point{
		{0, 0}, {10, 0}, {5, 10}, // Outer triangle
		{4, 2}, {6, 2}, {5, 4},   // Inner triangle
		{1, 1}, {9, 1},           // Noise points
	}

	if HasNestedTriangles(points) {
		fmt.Println("Result: The set contains nested triangles.")
	} else {
		fmt.Println("Result: No nested triangles found.")
	}
}
