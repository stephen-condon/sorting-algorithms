package bubblesort

import (
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {
	type test struct {
		input []int
		want  []int
	}

	tests := []test{
		{[]int{3, 2, 5, 1, 4}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 2, 2, 1, 4}, []int{1, 2, 2, 3, 4}},
		{[]int{3, 2, 5, 1, 4, 8, 12, 45, -4, 1, 56, 0, -13}, []int{-13, -4, 0, 1, 1, 2, 3, 4, 5, 8, 12, 45, 56}},
	}

	for _, tc := range tests {
		got := Sort(tc.input)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf(`Want %v, got: %v, input: %v`, tc.want, got, tc.input)
		}
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Sort([]int{3, 2, 5, 1, 4, 8, 12, 45, -4, 1, 56, 0, -13})
	}
}
