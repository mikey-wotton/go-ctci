package arraysandstrings

import "testing"

func Test_IsUnique(t *testing.T) {
	tests := map[string]struct {
		str    string
		answer bool
	}{
		"Testing hello": {
			str:    "hello",
			answer: false,
		},
		"Testing subdermatoglyphic": {
			str:    "subdermatoglyphic",
			answer: true,
		},
		"Testing empty string": {
			str:    "",
			answer: true,
		},
		"Testing rrrr": {
			str:    "rrrr",
			answer: false,
		},
	}

	for desc, test := range tests {
		got := isUnique(test.str)

		if got != test.answer {
			t.Errorf("%s failed - got %v but want %v", desc, got, test.answer)
		}
	}
}
