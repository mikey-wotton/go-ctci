package chapter2

import (
	"reflect"
	"testing"
)

func Test_dedupeHashMape(t *testing.T) {
	tests := map[string]struct {
		node *node
		want *node
	}{
		"empty node returns empty node": {
			node: &node{},
			want: &node{},
		},
		"removes duplicate from middle of list": {
			node: &node{1, &node{2, &node{2, &node{3, nil}}}},
			want: &node{1, &node{2, &node{3, nil}}},
		},
		"removes duplicate from end of list": {
			node: &node{1, &node{2, &node{3, &node{3, nil}}}},
			want: &node{1, &node{2, &node{3, nil}}},
		},
		"removes unsorted duplicates from all over list": {
			node: &node{1, &node{4, &node{2, &node{3, &node{2, &node{3, &node{3, &node{4, &node{3, nil}}}}}}}}},
			want: &node{1, &node{4, &node{2, &node{3, nil}}}},
		},
	}
	for desc, test := range tests {
		dedupeHashMap(test.node)
		if !reflect.DeepEqual(test.node, test.want) {
			t.Errorf("%s test failed, got %v but want %v", desc, test.node, test.want)
		}
	}
}
