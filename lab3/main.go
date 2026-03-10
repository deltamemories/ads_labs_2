package main

import (
	"fmt"
)

const AlphabetSize = 256

// computeTransitionTable builds the transition table for the finite automaton.
func computeTransitionTable(pattern string) [][]int {
	m := len(pattern)

	tf := make([][]int, m+1)
	for i := range tf {
		tf[i] = make([]int, AlphabetSize)
	}

	if m == 0 {
		return tf
	}

	// Base case for state 0.
	tf[0][pattern[0]] = 1

	// lps tracks the state to fall back to on mismatch.
	lps := 0

	for i := 1; i <= m; i++ {
		// Copy transitions from the fallback state.
		for x := 0; x < AlphabetSize; x++ {
			tf[i][x] = tf[lps][x]
		}

		if i < m {
			// Update transition for the correct next character in pattern.
			tf[i][pattern[i]] = i + 1

			// Update the fallback state for the next step.
			lps = tf[lps][pattern[i]]
		}
	}

	return tf
}

// SearchPatternFA searches for all occurrences of a pattern using a finite automaton.
func SearchPatternFA(text string, pattern string) []int {
	m := len(pattern)
	n := len(text)
	var result []int

	if m == 0 || m > n {
		return result
	}

	tf := computeTransitionTable(pattern)

	state := 0
	for i := 0; i < n; i++ {
		state = tf[state][text[i]]

		if state == m {
			result = append(result, i-m+1)
		}
	}

	return result
}

func main() {
	text := "AABAACAADAABAABA"
	pattern := "AABA"

	fmt.Printf("Text:    %s\n", text)
	fmt.Printf("Pattern: %s\n", pattern)

	indices := SearchPatternFA(text, pattern)

	if len(indices) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", indices)
	} else {
		fmt.Println("Pattern not found.")
	}
}
