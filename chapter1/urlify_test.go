package arraysandstrings

import "testing"

func Test_urlify(t *testing.T) {
	tests := map[string]struct {
		str  string
		want string
	}{
		"original test": {
			str:  "Mr John Smith    ",
			want: "Mr%20John%20Smith",
		},
		"single space is removed": {
			str:  "Hello Bob  ",
			want: "Hello%20Bob",
		},
		"two spaces removed": {
			str:  "check this out    ",
			want: "check%20this%20out",
		},
	}

	for desc, test := range tests {
		if got := urlify(test.str); got != test.want {
			t.Errorf("%s failed got '%s' but want '%s'", desc, got, test.want)
		}
	}
}
