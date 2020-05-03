package arraysandstrings

import "sort"

/*
  Check Permutation: Given two strings, write a method to decide if one is a
  permutation of the other.

  Hint #1: Describe what it means for two strings to be permutations of each other.
  Now, look at that definition you provided. Can you check the strings against
  that definition?

  Hint #84: There is one solution that is 0 (N log N) time. Another solution
  uses some space, but isO(N) time

  Hint #122: Could a hash table be useful?

  Hint #131: Can you do all three checks in a single pass?
*/
// Quicksort: O(n log n)
func isPermutation(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}

	s, p = sortString(s), sortString(p)
	for i := range s {
		if s[i] != p[i] {
			return false
		}
	}

	return true
}

// One time pass but extra space required: O(N)
func isPermutationExtraSpace(s string, p string) bool {
	if len(s) != len(p) {
		return false
	}

	seen := [256]byte{}

	for i := range s {
		seen[s[i]]++
		seen[p[i]]--
	}

	for _, count := range seen {
		if count != 0 {
			return false
		}
	}

	return true
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func sortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
