package chapter1

/*
  One Away: There are three types of edits that can be performed on strings:
  insert a character, remove a character, or replace a character. Given two
  strings, write a function to check if they are one edit (or zero edits) away.
  EXAMPLE
  pale, ple -> true
  pales, pale -> true
  pale, bale -> true
  pale, bake -> false

  Hint #23: Start with the easy thing. Can you check each of the conditions
  separately?

  Hint #97: What is the relationship between the "insert character" option and
  the "remove character" option? Do these need to be two separate checks?

  Hint #130: Can you do all three checks in a single pass?
*/
//Single pass: O(n)
func isOneAway(str1, str2 string) bool {
	l1 := len(str1)
	l2 := len(str2)

	if Abs(l1-l2) > 1 {
		return false
	}

	i, k, differenceCount := 0, 0, 0
	for i < l1 && k < l2 {
		if str1[i] != str2[k] {
			if differenceCount == 1 {
				return false
			}

			if l1 > l2 {
				i++
			} else if l1 < l2 {
				k++
			} else {
				i++
				k++
			}

			differenceCount++
		} else {
			i++
			k++
		}
	}

	if i < l1 || k < l2 {
		differenceCount++
	}

	return differenceCount == 1
}

// Abs returns the absolute value of x.
func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
