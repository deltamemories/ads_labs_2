package main

import (
	"fmt"
)

const INF = 1000000000

// TSP solves the Traveling Salesperson Problem using Dynamic Programming with bitmask.
func TSP(dist [][]int) int {
	n := len(dist)
	// dp[mask][i] is the shortest path to visit cities in mask, ending at city i
	dp := make([][]int, 1<<n)
	for i := range dp {
		dp[i] = make([]int, n)
		for j := range dp[i] {
			dp[i][j] = INF
		}
	}

	// Starting city is 0
	dp[1][0] = 0

	for mask := 1; mask < (1 << n); mask++ {
		for u := 0; u < n; u++ {
			if dp[mask][u] == INF {
				continue
			}
			// Try to go to city v
			for v := 0; v < n; v++ {
				if (mask & (1 << v)) == 0 {
					newMask := mask | (1 << v)
					if dp[newMask][v] > dp[mask][u]+dist[u][v] {
						dp[newMask][v] = dp[mask][u] + dist[u][v]
					}
				}
			}
		}
	}

	// Find the minimum cost to return to city 0
	res := INF
	fullMask := (1 << n) - 1
	for i := 1; i < n; i++ {
		if res > dp[fullMask][i]+dist[i][0] {
			res = dp[fullMask][i] + dist[i][0]
		}
	}

	return res
}

func main() {
	// Example connectivity matrix
	dist := [][]int{
		{0, 10, 15, 20},
		{10, 0, 35, 25},
		{15, 35, 0, 30},
		{20, 25, 30, 0},
	}

	minCost := TSP(dist)
	fmt.Println("Connectivity Matrix:")
	for _, row := range dist {
		fmt.Printf("%v\n", row)
	}
	fmt.Printf("The shortest possible route cost is: %d\n", minCost)
}
