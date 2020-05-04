package arraysandstrings

/*
  URLify: Write a method to replace all spaces in a string with '%20: You may
  assume that the string has sufficient space at the end to hold the additional
  characters, and that you are given the "true" length of the string.
  Note: If implementing in Java, please use a character array so that you can
  perform this operation in place.)
  EXAMPLE
  Input: "Mr John Smith    ", 13
  Output: "Mr%20John%20Smith"

  Hint #58: Try thinking about the array as circular, such that the end of the
  array "wraps around" to the start of the array.

  Hint #118: You might find you need to know the number of spaces. Can you
  just count them?

*/

const spaceASCII = 32 //the ascii byte representation for space

func urlify(s []rune, actual int) {
	l := len(s)

	if l == actual {
		return
	}

	for i := actual - 1; i >= 0; i-- {
		if s[i] == spaceASCII {
			s[l-1] = '0'
			s[l-2] = '2'
			s[l-3] = '%'
			l -= 3
		} else {
			s[l-1] = s[i]
			l--
		}
	}
}
