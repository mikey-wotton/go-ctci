package chapter1

import (
	"reflect"
	"testing"
)

func Test_rotateMatrix(t *testing.T) {
	tests := map[string]struct {
		matrix [][]byte
		want   [][]byte
	}{
		"2x2 matrix": {
			matrix: [][]byte{
				{1, 2},
				{3, 4},
			},
			want: [][]byte{
				{3, 1},
				{4, 2},
			},
		},
		"4x4 matrix": {
			matrix: [][]byte{
				{1, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]byte{
				{13, 9, 5, 1},
				{14, 10, 6, 2},
				{15, 11, 7, 3},
				{16, 12, 8, 4},
			},
		},
	}

	for desc, test := range tests {
		got := rotateMatrix(test.matrix)

		if !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s failed, got %v but want %v", desc, got, test.want)
		}
	}
}
