package arraysandstrings

import "testing"

func Test_isPalindrome(t *testing.T) {
	tests := map[string]struct {
		str  string
		want bool
	}{
		"example test case provided": {
			str:  "tact coa",
			want: true,
		},
		"single letter in middle palindrome": {
			str:  "madam",
			want: true,
		},
		"dual letter middle palindrome": {
			str:  "never odd or even",
			want: true,
		},
		"not a palindrome": {
			str:  "beep",
			want: false,
		},
	}

	for desc, test := range tests {
		if got := isPalindrome(test.str); got != test.want {
			t.Errorf("%s test failed, got %v but want %v", desc, got, test.want)
		}
	}
}
