package chapter1

import "testing"

func Test_isRotated(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		want   bool
	}{
		"empty string test": {
			s1:   "",
			s2:   "",
			want: false,
		},
		"mismatching lengths test": {
			s1:   "fish",
			s2:   "steve",
			want: false,
		},
		"small negative test": {
			s1:   "boot",
			s2:   "toob",
			want: false,
		},
		"example waterbottle test": {
			s1:   "waterbottle",
			s2:   "erbottlewat",
			want: true,
		},
	}

	for desc, test := range tests {
		if got := isRotated(test.s1, test.s2); got != test.want {
			t.Errorf("%s failed, got %v but want %v", desc, got, test.want)
		}
	}
}
