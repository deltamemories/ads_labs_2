package main

import (
	"fmt"
)

// MaxSubArray finds the contiguous subarray within a one-dimensional array of numbers which has the largest sum.
func MaxSubArray(nums []int) (int, []int) {
	if len(nums) == 0 {
		return 0, nil
	}

	maxSoFar := nums[0]
	currentMax := nums[0]
	start, end, tempStart := 0, 0, 0

	for i := 1; i < len(nums); i++ {
		if nums[i] > currentMax+nums[i] {
			currentMax = nums[i]
			tempStart = i
		} else {
			currentMax = currentMax + nums[i]
		}

		if currentMax > maxSoFar {
			maxSoFar = currentMax
			start = tempStart
			end = i
		}
	}

	return maxSoFar, nums[start : end+1]
}

func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	maxSum, subArray := MaxSubArray(nums)

	fmt.Printf("Array: %v\n", nums)
	fmt.Printf("Maximum contiguous sum is %d\n", maxSum)
	fmt.Printf("Subarray: %v\n", subArray)
}
