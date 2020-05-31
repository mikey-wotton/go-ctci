package extras

import "testing"

func Test_isBalanced(t *testing.T) {
	tests := map[string]struct {
		str   string
		pairs map[rune]rune
		want  bool
	}{
		"simple positive": {
			str:   "()",
			pairs: map[rune]rune{'(': ')'},
			want:  true,
		},
		"simple negative": {
			str:   ")(",
			pairs: map[rune]rune{'(': ')'},
			want:  false,
		},
		"nested positive": {
			str:   "((()))",
			pairs: map[rune]rune{'(': ')'},
			want:  true,
		},
		"nested negative": {
			str:   "(((())",
			pairs: map[rune]rune{'(': ')'},
			want:  false,
		},
		"balanced multiple open close": {
			str:   "()()()",
			pairs: map[rune]rune{'(': ')'},
			want:  true,
		},
		"unbalanced string": {
			str:   "((())",
			pairs: map[rune]rune{'(': ')'},
			want:  false,
		},
		"complex positive": {
			str:   "({[]})",
			pairs: map[rune]rune{'(': ')', '[': ']', '{': '}'},
			want:  true,
		},
		"unspecified characters generate false": {
			str:   "(d{[]}d)",
			pairs: map[rune]rune{'(': ')', '[': ']', '{': '}'},
			want:  false,
		},
		"handles wrong closer bracket": {
			str:   "([{(}}])",
			pairs: map[rune]rune{'(': ')', '[': ']', '{': '}'},
			want:  false,
		},
	}

	for desc, test := range tests {
		if got := isBalanced(test.str, test.pairs); got != test.want {
			t.Errorf("%s test failed, got %v but want %v", desc, got, test.want)
		}
	}
}
