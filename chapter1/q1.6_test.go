package chapter1

import "testing"

func Test_compressString(t *testing.T) {
	tests := map[string]struct {
		str  string
		want string
	}{
		"example": {
			str:  "aabcccccaaa",
			want: "a2b1c5a3",
		},
		"example 2": {
			str:  "zzeeerrrrttttt",
			want: "z2e3r4t5",
		},
		"string length increased by compression returns original": {
			str:  "abcdefg",
			want: "abcdefg",
		},
		"compression of same size returns original": {
			str:  "aabb", //a2b2 compressed
			want: "aabb",
		},
		"empty string should return empty string": {
			str:  "",
			want: "",
		},
	}

	for desc, test := range tests {
		if got := compressString(test.str); got != test.want {
			t.Errorf("%s test failed, got %s but want %s", desc, got, test.want)
		}
	}
}
