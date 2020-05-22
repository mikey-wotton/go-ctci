package chapter2

import (
	"reflect"
	"testing"
)

func Test_returnKth(t *testing.T) {
	tests := map[string]struct {
		head *node
		num  int
		want *node
	}{
		"return just last node": {
			head: &node{1, &node{2, nil}},
			num:  0,
			want: &node{2, nil},
		},
		"return all nodes": {
			head: &node{1, &node{2, nil}},
			num:  1,
			want: &node{1, &node{2, nil}},
		},
		"return last 4 nodes": {
			head: &node{1, &node{2, &node{3, &node{4, &node{5, &node{6, nil}}}}}},
			num:  3,
			want: &node{3, &node{4, &node{5, &node{6, nil}}}},
		},
		"return nil as target not in list": {
			head: &node{1, &node{2, nil}},
			num:  2,
			want: nil,
		},
	}

	for desc, test := range tests {
		if got := returnKth(test.head, test.num); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s test failed, got %v want %v", desc, got, test.want)
		}
	}
}
