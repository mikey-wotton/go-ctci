package arraysandstrings

/*
  Palindrome Permutation: Given a string, write a function to check if it is a
  permutation of a palindrome. A palindrome is a word or phrase that is the same
  forwards and backwards. A permutation is a rearrangement of letters.

  The palindrome does not need to be limited to just dictionary words.

  EXAMPLE
  Input: Tact Coa
  Output: True (permutations: "taco cat". "atco cta". etc.)

  Hint #106: You do not have to-and should not-generate all permutations. This
  would be very inefficient.

  Hint #121: What characteristics would a string that is a permutation of a
  palindrome have?

  Hint #134: Have you tried a hash table? You should be able to get this down
  to 0 (N) time.

  Hint #136: Can you reduce the space usage by using a bit vector?
*/
//Note: will not handle casing. Nan is not treated as a palindrome for example.
func isPalindrome(s string) bool {
	m := make(map[rune]int, len(s))

	for _, r := range s {
		if r != spaceASCII {
			if _, exists := m[r]; !exists {
				m[r] = 0
			}
			m[r]++
		}
	}

	oddCount := 0
	for _, v := range m {
		remainder := v % 2
		if remainder > 0 {
			oddCount++
		}

		if oddCount > 1 {
			return false
		}
	}

	return true
}
