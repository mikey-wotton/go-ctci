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
		"same length similar but not a permutation": {
			str:  "qwertyuiop",
			perm: "qwertyuiol",
			want: false,
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
		"empty string is a permutation": {
			str:  "",
			perm: "",
			want: true,
		},
		"all repeating letters is a permutation ": {
			str:  "rrrrrrrrrrr",
			perm: "rrrrrrrrrrr",
			want: true,
		},
	}

	for desc, test := range tests {
		if got := isPermutation(test.str, test.perm); got != test.want {
			t.Errorf("%s permutation test failed, got %v but want %v", desc, got, test.want)
		}
	}
}

func Test_isPermutationExtraSpace(t *testing.T) {
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
		"same length similar but not a permutation": {
			str:  "qwertyuiop",
			perm: "qwertyuiol",
			want: false,
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
		"empty string is a permutation": {
			str:  "",
			perm: "",
			want: true,
		},
		"all repeating letters is a permutation ": {
			str:  "rrrrrrrrrrr",
			perm: "rrrrrrrrrrr",
			want: true,
		},
	}

	for desc, test := range tests {
		if got := isPermutationExtraSpace(test.str, test.perm); got != test.want {
			t.Errorf("%s permutation test failed, got %v but want %v", desc, got, test.want)
		}
	}
}
