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

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// EggDrop finds the minimum number of trials needed in the worst case with n eggs and k floors.
func EggDrop(n, k int) int {
	// dp[i][j] represents minimum trials with i eggs and j floors
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k+1)
	}

	// 1 trial for 1 floor, 0 trials for 0 floors
	for i := 1; i <= n; i++ {
		dp[i][1] = 1
		dp[i][0] = 0
	}

	// j trials for 1 egg and j floors
	for j := 1; j <= k; j++ {
		dp[1][j] = j
	}

	for i := 2; i <= n; i++ {
		for j := 2; j <= k; j++ {
			dp[i][j] = 1e9
			for x := 1; x <= j; x++ {
				res := 1 + max(dp[i-1][x-1], dp[i][j-x])
				dp[i][j] = min(dp[i][j], res)
			}
		}
	}

	return dp[n][k]
}

func main() {
	eggs := 2
	floors := 100
	fmt.Printf("Number of eggs: %d, Number of floors: %d\n", eggs, floors)
	fmt.Printf("Minimum number of trials in worst case is %d\n", EggDrop(eggs, floors))
}
