package chapter2

import (
	"reflect"
	"testing"
)

func Test_deleteNode(t *testing.T) {
	tests := map[string]struct {
		list *node
		num  int
		want *node
	}{
		"delete example small": {
			list: &node{1, &node{5, &node{9, &node{12, nil}}}},
			num:  2,
			want: &node{1, &node{5, &node{12, nil}}},
		},
		"delete middle small": {
			list: &node{1, &node{2, &node{3, nil}}},
			num:  1,
			want: &node{1, &node{3, nil}},
		},
		"delete 2nd node": {
			list: &node{1, &node{2, &node{3, &node{4, &node{5, &node{6, &node{7, nil}}}}}}},
			num:  1,
			want: &node{1, &node{3, &node{4, &node{5, &node{6, &node{7, nil}}}}}},
		},
		"delete 2nd to last node": {
			list: &node{1, &node{2, &node{3, &node{4, &node{5, &node{6, &node{7, nil}}}}}}},
			num:  5,
			want: &node{1, &node{2, &node{3, &node{4, &node{5, &node{7, nil}}}}}},
		},
	}

	for desc, test := range tests {
		n := test.list
		for i := 0; i < test.num; i++ {
			n = n.next
		}
		deleteNode(n)
		if !reflect.DeepEqual(test.list, test.want) {
			t.Errorf("%s test failed, got %v but want %v", desc, test.list, test.want)
		}
	}
}
