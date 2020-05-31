package extras

func isBalanced(s string, openers map[rune]rune) bool {
	if len(s)%2 != 0 {
		return false
	}

	stack := make([]rune, 0, len(s)/2)
	for _, r := range s {
		if _, exists := openers[r]; !exists {
			if len(stack) < 1 {
				return false //found a closer but not expected one
			}

			lastSeenOpener := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			if closer, exists := openers[lastSeenOpener]; exists {
				if closer != r {
					return false
				}
			}

		} else {
			stack = append(stack, r)
		}
	}

	return len(stack) == 0
}
