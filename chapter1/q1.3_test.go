package chapter1

import "testing"

func Test_urlify(t *testing.T) {
	tests := map[string]struct {
		str  []rune
		len  int
		want []rune
	}{
		"original test": {
			str:  []rune("Mr John Smith    "),
			len:  13,
			want: []rune("Mr%20John%20Smith"),
		},
		"single space is removed": {
			str:  []rune("Hello Bob  "),
			len:  9,
			want: []rune("Hello%20Bob"),
		},
		"two spaces removed": {
			str:  []rune("check this out    "),
			len:  14,
			want: []rune("check%20this%20out"),
		},
		"two spaces removed at once": {
			str:  []rune("beep  boop    "),
			len:  10,
			want: []rune("beep%20%20boop"),
		},
		"improper len provided": {
			str:  []rune("beep  boop"),
			len:  10,
			want: []rune("beep  boop"),
		},
	}

	for desc, test := range tests {
		urlify(test.str, test.len)
		for i := range test.want {
			if test.str[i] != test.want[i] {
				t.Errorf("%s failed at index %d got '%s' but want '%s'", desc, i, string(test.str[i]), string(test.want[i]))
			}
		}
	}
}
