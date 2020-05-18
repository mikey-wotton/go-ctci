package chapter1

import "testing"

func Test_isRotated(t *testing.T) {
	tests := map[string]struct {
		s1, s2 string
		want   bool
	}{
		"small test": {
			s1:   "boot",
			s2:   "toob",
			want: true,
		},
	}

	for desc, test := range tests {
		if got := isRotated(test.s1, test.s2); got != test.want {
			t.Errorf("%s failed, got %v but want %v", desc, got, test.want)
		}
	}
}
