package chapter1

import "testing"

/*
pale, ple -> true
pales, pale -> true
pale, bale -> true
pale, bake -> false

*/
func Test_isOneAway(t *testing.T) {
	tests := map[string]struct {
		str1 string
		str2 string
		want bool
	}{
		"remove letter": {
			str1: "pale",
			str2: "ple",
			want: true,
		},
		"add a letter": {
			str1: "pales",
			str2: "pale",
			want: true,
		},
		"exchange a letter": {
			str1: "pale",
			str2: "bale",
			want: true,
		},
		"not one away": {
			str1: "pale",
			str2: "bake",
			want: false,
		},
	}

	for desc, test := range tests {
		if got := isOneAway(test.str1, test.str2); got != test.want {
			t.Errorf("%s test failed, got %v but want %v", desc, got, test.want)
		}
	}
}
