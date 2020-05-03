package arraysandstrings

import "testing"

func Test_isPermutation(t *testing.T) {
	tests := map[string]struct {
		str  string
		perm string
		want bool
	}{
		"backwards word is permutation": {
			str:  "hello",
			perm: "olleh",
			want: true,
		},
		"different length strings cannot be permutations": {
			str:  "aa",
			perm: "a",
			want: false,
		},
		"jumbled letters are permutations": {
			str:  "qwesd",
			perm: "qdsew",
			want: true,
		},
		"all repeating letters is a permutation ": {
			str:  "rrrrrrrrrrrrrrrrr",
			perm: "rrrrrrrrrrrrrrrrr",
			want: true,
		},
	}

	for desc, test := range tests {
		if got := is_permutation(test.str, test.perm); got != test.want {
			t.Errorf("%s permutation test failed, got %v but want %v", desc, got, test.want)
		}
	}

}
