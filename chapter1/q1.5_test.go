package chapter1

import "testing"

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
		"exchange a letter 2nd string is longer": {
			str1: "pale",
			str2: "kales",
			want: false,
		},
		"remove letter 2nd string is longer": {
			str1: "bale",
			str2: "bales",
			want: true,
		},
		"first string is too long to be one away": {
			str1: "bale12",
			str2: "bale",
			want: false,
		},
		"not one away because same": {
			str1: "bake",
			str2: "bake",
			want: false,
		},
		"more than one away": {
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
