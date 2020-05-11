package chapter1

import (
	"reflect"
	"testing"
)

func Test_zeroMatrix(t *testing.T) {
	tests := map[string]struct {
		matrix [][]int
		want   [][]int
	}{
		"empty matrix should return empty matrix": {
			matrix: [][]int{},
			want:   [][]int{},
		},
		"2x2 with one zero only zeros that column and row": {
			matrix: [][]int{
				{1, 2},
				{3, 0},
			},
			want: [][]int{
				{1, 0},
				{0, 0},
			},
		},
		"4x4 with two zero zeros that column and row": {
			matrix: [][]int{
				{0, 2, 3, 4},
				{5, 6, 7, 8},
				{9, 10, 0, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 6, 0, 8},
				{0, 0, 0, 0},
				{0, 14, 0, 16},
			},
		},
		"2x4 with one zero  that column and row": {
			matrix: [][]int{
				{0, 2, 3, 4},
				{5, 6, 7, 8},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 6, 7, 8},
			},
		},
		"4x4 with row of zeros makes entire matrix zero": {
			matrix: [][]int{
				{0, 0, 0, 0},
				{5, 6, 7, 8},
				{9, 10, 11, 12},
				{13, 14, 15, 16},
			},
			want: [][]int{
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
				{0, 0, 0, 0},
			},
		},
	}

	for desc, test := range tests {
		if got := zeroMatrix(test.matrix); !reflect.DeepEqual(got, test.want) {
			t.Errorf("%s test failed, got %v but want %v", desc, got, test.want)
		}
	}
}
