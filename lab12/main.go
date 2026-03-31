package main

import (
	"fmt"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Knapsack solves the 0/1 Knapsack problem using Dynamic Programming.
func Knapsack(W int, weights, values []int) int {
	n := len(values)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, W+1)
	}

	for i := 1; i <= n; i++ {
		for w := 1; w <= W; w++ {
			if weights[i-1] <= w {
				dp[i][w] = max(values[i-1]+dp[i-1][w-weights[i-1]], dp[i-1][w])
			} else {
				dp[i][w] = dp[i-1][w]
			}
		}
	}

	return dp[n][W]
}

func main() {
	values := []int{60, 100, 120}
	weights := []int{10, 20, 30}
	W := 50

	maxVal := Knapsack(W, weights, values)

	fmt.Printf("Values: %v\n", values)
	fmt.Printf("Weights: %v\n", weights)
	fmt.Printf("Knapsack Capacity: %d\n", W)
	fmt.Printf("Maximum value in Knapsack = %d\n", maxVal)
}
