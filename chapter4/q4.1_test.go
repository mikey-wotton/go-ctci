package chapter4

import "testing"

func Test_isRouteDFS(t *testing.T) {
	tests := map[string]struct {
		root   *node
		target int
		want   bool
	}{
		"single positive root": {
			root: &node{
				val:   0,
				edges: []*node{{1, nil}},
			},
			target: 1,
			want:   true,
		},
		"branching positive root": {
			root: &node{
				val: 0,
				edges: []*node{
					{1, []*node{{2, nil}, {3, nil}}},
				},
			},
			target: 3,
			want:   true,
		},
		"branching multiple": {
			root: &node{
				val: 0,
				edges: []*node{
					{1,
						[]*node{{2, []*node{{3, nil}, {4, nil}, {5, nil}}},
							{10, []*node{{11, []*node{{12, nil}}}}}}},
				},
			},
			target: 12,
			want:   true,
		},
		"branching multiple negative": {
			root: &node{
				val: 0,
				edges: []*node{
					{1,
						[]*node{{2, []*node{{3, nil}, {4, nil}, {5, nil}}},
							{10, []*node{{11, []*node{{12, nil}}}}}}},
				},
			},
			target: 55,
			want:   false,
		},
	}

	for desc, test := range tests {
		if got := isRouteDFS(test.root, test.target); got != test.want {
			t.Errorf("%s test failed, got %v but want %v", desc, got, test.want)
		}
	}
}

func Test_isRouteBFS(t *testing.T) {
	tests := map[string]struct {
		root   *node
		target int
		want   bool
	}{
		"single positive root": {
			root: &node{
				val:   0,
				edges: []*node{{1, nil}},
			},
			target: 1,
			want:   true,
		},
		"branching positive root": {
			root: &node{
				val: 0,
				edges: []*node{
					{1, []*node{{2, nil}, {3, nil}}},
				},
			},
			target: 3,
			want:   true,
		},
		"branching multiple": {
			root: &node{
				val: 0,
				edges: []*node{
					{1,
						[]*node{
							{2, []*node{{3, nil}, {4, nil}, {5, nil}}},
							{10, []*node{{11, []*node{{12, nil}}}}},
						}},
				},
			},
			target: 12,
			want:   true,
		},
		"branching multiple negative": {
			root: &node{
				val: 0,
				edges: []*node{
					{1, []*node{{2, []*node{{3, nil}, {4, nil}, {5, nil}}},
						{10, []*node{{11, []*node{{12, nil}}}}}}},
				},
			},
			target: 55,
			want:   false,
		},
	}

	for desc, test := range tests {
		if got := isRouteBFS(test.root, test.target); got != test.want {
			t.Errorf("%s test failed, got %v but want %v", desc, got, test.want)
		}
	}
}
