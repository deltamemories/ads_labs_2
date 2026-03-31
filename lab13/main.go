package main

import (
	"fmt"
	"sort"
)

// FirstFitDecreasing solves the Bin Packing problem using the FFD heuristic.
func FirstFitDecreasing(items []int, binCapacity int) int {
	// Sort items in decreasing order
	sort.Slice(items, func(i, j int) bool {
		return items[i] > items[j]
	})

	var bins []int

	for _, item := range items {
		placed := false
		for i := 0; i < len(bins); i++ {
			if bins[i] >= item {
				bins[i] -= item
				placed = true
				break
			}
		}

		if !placed {
			bins = append(bins, binCapacity-item)
		}
	}

	return len(bins)
}

func main() {
	items := []int{2, 5, 4, 7, 1, 3, 8}
	binCapacity := 10

	numBins := FirstFitDecreasing(items, binCapacity)

	fmt.Printf("Items: %v\n", items)
	fmt.Printf("Bin Capacity: %d\n", binCapacity)
	fmt.Printf("Number of bins required (FFD heuristic): %d\n", numBins)
}
