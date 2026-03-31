package main

import (
	"fmt"
)

// computeLPSArray computes the Longest Proper Prefix which is also Suffix array.
func computeLPSArray(pattern string) []int {
	m := len(pattern)
	lps := make([]int, m)
	length := 0
	i := 1

	for i < m {
		if pattern[i] == pattern[length] {
			length++
			lps[i] = length
			i++
		} else {
			if length != 0 {
				length = lps[length-1]
			} else {
				lps[i] = 0
				i++
			}
		}
	}
	return lps
}

// SearchKMP searches for all occurrences of pattern in text using KMP algorithm.
func SearchKMP(text, pattern string) []int {
	n := len(text)
	m := len(pattern)
	var result []int

	if m == 0 {
		return result
	}

	lps := computeLPSArray(pattern)
	i := 0 // index for text
	j := 0 // index for pattern

	for i < n {
		if pattern[j] == text[i] {
			i++
			j++
		}

		if j == m {
			result = append(result, i-j)
			j = lps[j-1]
		} else if i < n && pattern[j] != text[i] {
			if j != 0 {
				j = lps[j-1]
			} else {
				i++
			}
		}
	}
	return result
}

func main() {
	text := "ABABDABACDABABCABAB"
	pattern := "ABABCABAB"

	fmt.Printf("Text:    %s\n", text)
	fmt.Printf("Pattern: %s\n", pattern)

	indices := SearchKMP(text, pattern)

	if len(indices) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", indices)
	} else {
		fmt.Println("Pattern not found.")
	}
}
