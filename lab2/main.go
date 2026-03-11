package main

import (
	"fmt"
	"math"
)

type Point struct {
	X, Y float64
}

func sign(p1, p2, p3 Point) float64 {
	return (p1.X-p3.X)*(p2.Y-p3.Y) - (p2.X-p3.X)*(p1.Y-p3.Y)
}

func isPointStrictlyInside(pt, v1, v2, v3 Point) bool {
	d1 := sign(pt, v1, v2)
	d2 := sign(pt, v2, v3)
	d3 := sign(pt, v3, v1)

	// Проверка ориентации: все знаки должны быть строго положительными или строго отрицательными
	hasNeg := (d1 < -1e-9) || (d2 < -1e-9) || (d3 < -1e-9)
	hasPos := (d1 > 1e-9) || (d2 > 1e-9) || (d3 > 1e-9)
	
	// Если точка на границе (один из d равен 0), hasNeg и hasPos не дадут полной картины,
	// поэтому проверяем близость к нулю для строгой вложенности.
	if math.Abs(d1) <= 1e-9 || math.Abs(d2) <= 1e-9 || math.Abs(d3) <= 1e-9 {
		return false
	}

	return !(hasNeg && hasPos)
}

func isCollinear(p1, p2, p3 Point) bool {
	area := p1.X*(p2.Y-p3.Y) + p2.X*(p3.Y-p1.Y) + p3.X*(p1.Y-p2.Y)
	return math.Abs(area) < 1e-9
}

func HasNestedTrianglesOptimized(points []Point) bool {
	n := len(points)
	if n < 6 {
		return false
	}

	// 1. Выбираем внешний треугольник (три точки) - O(n^3)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			for k := j + 1; k < n; k++ {
				p1, p2, p3 := points[i], points[j], points[k]
				
				if isCollinear(p1, p2, p3) {
					continue
				}

				// 2. Ищем точки, которые лежат СТРОГО внутри этого треугольника
				var pointsInside []Point
				for m := 0; m < n; m++ {
					if m == i || m == j || m == k {
						continue
					}
					if isPointStrictlyInside(points[m], p1, p2, p3) {
						pointsInside = append(pointsInside, points[m])
					}
				}

				// 3. Если внутри достаточно точек (минимум 3), проверяем, 
				// не лежат ли они на одной прямой.
				if len(pointsInside) >= 3 {
					for a := 0; a < len(pointsInside); a++ {
						for b := a + 1; b < len(pointsInside); b++ {
							for c := b + 1; c < len(pointsInside); c++ {
								if !isCollinear(pointsInside[a], pointsInside[b], pointsInside[c]) {
									return true 
								}
							}
						}
					}
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
		{1, 1}, {9, 1},
	}

	if HasNestedTrianglesOptimized(points) {
		fmt.Println("Result: The set contains nested triangles.")
	} else {
		fmt.Println("Result: No nested triangles found.")
	}
}