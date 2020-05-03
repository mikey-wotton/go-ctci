package arraysandstrings

import "testing"

var uniqueTests = map[string]struct {
	str    string
	answer bool
}{
	"Testing hello": {
		str:    "hello",
		answer: false,
	},
	"Testing abcdefghijklmnopqrstuvwxyz": {
		str:    "abcdefghijklmnopqrstuvwxyz",
		answer: true,
	},
	"Testing AbCdEfG": {
		str:    "AbCdEfG",
		answer: true,
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

func Test_IsUnique(t *testing.T) {
	for desc, test := range uniqueTests {
		got := isUnique(test.str)

		if got != test.answer {
			t.Errorf("%s failed - got %v but want %v", desc, got, test.answer)
		}
	}
}

func Test_isUniqueBrute(t *testing.T) {
	for desc, test := range uniqueTests {
		got := isUniqueBrute(test.str)

		if got != test.answer {
			t.Errorf("%s failed - got %v but want %v", desc, got, test.answer)
		}
	}
}

func Test_isUniqueSorted(t *testing.T) {
	for desc, test := range uniqueTests {
		got := isUniqueSorted(test.str)

		if got != test.answer {
			t.Errorf("%s failed - got %v but want %v", desc, got, test.answer)
		}
	}
}
