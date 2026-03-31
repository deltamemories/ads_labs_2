package main

import (
	"fmt"
)

const AlphabetSize = 256

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// badCharHeuristic fills the bad character array.
func badCharHeuristic(pattern string) [AlphabetSize]int {
	var badChar [AlphabetSize]int
	for i := 0; i < AlphabetSize; i++ {
		badChar[i] = -1
	}
	for i := 0; i < len(pattern); i++ {
		badChar[pattern[i]] = i
	}
	return badChar
}

// SearchBoyerMoore searches for all occurrences of pattern in text using Boyer-Moore (Bad Character Heuristic).
func SearchBoyerMoore(text, pattern string) []int {
	m := len(pattern)
	n := len(text)
	var result []int

	if m == 0 || m > n {
		return result
	}

	badChar := badCharHeuristic(pattern)
	s := 0 // s is shift of the pattern with respect to text

	for s <= (n - m) {
		j := m - 1

		// Keep reducing index j of pattern while characters of pattern and text are matching at this shift s
		for j >= 0 && pattern[j] == text[s+j] {
			j--
		}

		if j < 0 {
			result = append(result, s)
			// Shift the pattern so that the next character in text aligns with the last occurrence of it in pattern.
			if s+m < n {
				s += m - badChar[text[s+m]]
			} else {
				s += 1
			}
		} else {
			// Shift the pattern so that the bad character in text aligns with the last occurrence of it in pattern.
			s += max(1, j-badChar[text[s+j]])
		}
	}
	return result
}

func main() {
	text := "ABAAABCDABC"
	pattern := "ABC"

	fmt.Printf("Text:    %s\n", text)
	fmt.Printf("Pattern: %s\n", pattern)

	indices := SearchBoyerMoore(text, pattern)

	if len(indices) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", indices)
	} else {
		fmt.Println("Pattern not found.")
	}
}
