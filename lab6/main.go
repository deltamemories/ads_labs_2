package main

import (
	"fmt"
)

// d is the number of characters in the input alphabet
const d = 256

// SearchRabinKarp searches for all occurrences of pattern in text using Rabin-Karp algorithm.
// q is a prime number for hashing.
func SearchRabinKarp(text, pattern string, q int) []int {
	m := len(pattern)
	n := len(text)
	var result []int

	if m == 0 || m > n {
		return result
	}

	h := 1
	p := 0 // hash value for pattern
	t := 0 // hash value for text

	// The value of h would be "pow(d, m-1) % q"
	for i := 0; i < m-1; i++ {
		h = (h * d) % q
	}

	// Calculate the hash value of pattern and first window of text
	for i := 0; i < m; i++ {
		p = (d*p + int(pattern[i])) % q
		t = (d*t + int(text[i])) % q
	}

	// Slide the pattern over text one by one
	for i := 0; i <= n-m; i++ {
		// Check the hash values of current window of text and pattern.
		// If the hash values match then only check for characters one by one
		if p == t {
			match := true
			for j := 0; j < m; j++ {
				if text[i+j] != pattern[j] {
					match = false
					break
				}
			}
			if match {
				result = append(result, i)
			}
		}

		// Calculate hash value for next window of text: Remove leading digit, add trailing digit
		if i < n-m {
			t = (d*(t-int(text[i])*h) + int(text[i+m])) % q

			// We might get negative value of t, converting it to positive
			if t < 0 {
				t = (t + q)
			}
		}
	}
	return result
}

func main() {
	text := "GEEKS FOR GEEKS"
	pattern := "GEEK"
	q := 101 // A prime number

	fmt.Printf("Text:    %s\n", text)
	fmt.Printf("Pattern: %s\n", pattern)

	indices := SearchRabinKarp(text, pattern, q)

	if len(indices) > 0 {
		fmt.Printf("Pattern found at indices: %v\n", indices)
	} else {
		fmt.Println("Pattern not found.")
	}
}
