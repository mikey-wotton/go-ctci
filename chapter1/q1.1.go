package chapter1

import (
	"sort"
	"strings"
)

/*
  Is Unique: Implement an algorithm to determine if a string has all unique characters. What if you
  cannot use additional data structures?
  Hints: #44, #117, #132

  Hint #44: Try a hash table.

  Hint #117: Could a bit vector be useful?

  Hint #132: Can you solve it in O(N log N) time? What might a solution like that look like?
*/

func isUnique(s string) bool {
	m := make(map[rune]struct{}, len(s))

	for _, r := range s {
		_, exists := m[r]
		if exists {
			return false
		}

		m[r] = struct{}{}
	}

	return true
}

//Brute force: O(n^2)
func isUniqueBrute(s string) bool {
	for i := range s {
		for k := i + 1; k < len(s); k++ {
			if s[i] == s[k] {
				return false
			}
		}
	}

	return true
}

//Brute force: O(n log n)
func isUniqueSorted(s string) bool {
	sorted := strings.Split(s, "")
	sort.Strings(sorted)

	for i := 0; i < len(sorted)-1; i++ {
		if sorted[i] == sorted[i+1] {
			return false
		}
	}

	return true
}
