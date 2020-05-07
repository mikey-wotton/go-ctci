package chapter1

import (
	"bytes"
	"fmt"
)

/*
  String Compression: Implement a method to perform basic string compression
  using the counts of repeated characters. For example, the string aabcccccaaa
  would become a2b1c5a3. If the "compressed" string would not become smaller
  than the original string, your method should return the original string.

  You can assume the string has only uppercase and lowercase letters (a - z).

  Hint #92: Do the easy thing first. Compress the string, then compare the
  lengths.

  Hint #110: Be careful that you aren't repeatedly concatenating strings
  together. This can be very inefficient.
*/
//Single pass: O(n)
func compressString(s string) string {
	length := len(s)
	if length < 3 {
		return s
	}

	var b bytes.Buffer

	count, temp := 0, byte(0)
	for i := range s {
		byt := s[i]
		if temp != byt {
			if count != 0 {
				b.WriteString(fmt.Sprintf("%s%d", string(temp), count))
			}
			temp = byt
			count = 0
		}
		count++
	}

	//write final letter
	b.WriteString(fmt.Sprintf("%s%d", string(temp), count))

	final := b.String()
	if len(final) >= length {
		return s
	}

	return final
}
