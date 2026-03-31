package main

import (
	"fmt"
)

// CountWaysToMakeChange finds the number of ways to make change for a given sum.
func CountWaysToMakeChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	dp[0] = 1

	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] += dp[i-coin]
		}
	}

	return dp[amount]
}

func main() {
	coins := []int{1, 2, 5}
	amount := 5
	ways := CountWaysToMakeChange(coins, amount)

	fmt.Printf("Coins: %v\n", coins)
	fmt.Printf("Amount: %d\n", amount)
	fmt.Printf("Number of ways to make change: %d\n", ways)
}
