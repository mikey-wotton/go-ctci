package chapter1

import "strings"

/*
  Assume you have a method isSubstring which checks if one word is a substring
  of another. Given two strings, S1 and S2, write code to check if S2 is a
  rotation of S1 using only one call to isSubstring
  (e.g., "waterbottle" is a rotation of"erbottlewat").

  Hint #34: If a string is a rotation of another, then it's a rotation at a
  particular point. For example, a rotation of water bottle at character 3 means
  cutting waterbottle at character 3 and putting the right half (erbottle)
  before the left half (wat).

  Hint #88: We are essentially asking if there's a way of splitting the first
  string into two parts, x and y, such that the first string is xy and the
  second string is yx. For example, x = wat and y = erbottle. The first string
  is xy = waterbottle. The second string is yx = erbottlewat.

  Hint #104: Think about the earlier hint. Then think about what happens when
  you concatenate erbottlewat to itself. You get erbottlewaterbottlewat.
*/

func isRotated(s1, s2 string) bool {
	n := len(s1)
	if n == 0 || n != len(s2) {
		return false
	}

	s1 += s1

	return strings.Contains(s1, s2)
}
