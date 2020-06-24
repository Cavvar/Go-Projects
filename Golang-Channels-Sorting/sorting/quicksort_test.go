package sorting

import (
	"math/rand"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestQuickSort(t *testing.T) {
	cases := []struct {
		in   []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 3, 2}, []int{1, 2, 3}},
		{[]int{2, 1, 3}, []int{1, 2, 3}},
		{[]int{2, 3, 1}, []int{1, 2, 3}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{1, 2, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 3, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{2, 1, 3, 4}, []int{1, 2, 3, 4}},
		{[]int{2, 3, 1, 4}, []int{1, 2, 3, 4}},
		{[]int{3, 1, 2, 4}, []int{1, 2, 3, 4}},
		{[]int{3, 2, 1, 4}, []int{1, 2, 3, 4}},
		{[]int{1, 2, 4, 3}, []int{1, 2, 3, 4}},
		{[]int{1, 3, 4, 2}, []int{1, 2, 3, 4}},
		{[]int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{[]int{2, 3, 4, 1}, []int{1, 2, 3, 4}},
		{[]int{3, 1, 4, 2}, []int{1, 2, 3, 4}},
		{[]int{3, 2, 4, 1}, []int{1, 2, 3, 4}},
		{[]int{1, 4, 2, 3}, []int{1, 2, 3, 4}},
		{[]int{1, 4, 3, 2}, []int{1, 2, 3, 4}},
		{[]int{2, 4, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{2, 4, 3, 1}, []int{1, 2, 3, 4}},
		{[]int{3, 4, 1, 2}, []int{1, 2, 3, 4}},
		{[]int{3, 4, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{4, 1, 2, 3}, []int{1, 2, 3, 4}},
		{[]int{4, 1, 3, 2}, []int{1, 2, 3, 4}},
		{[]int{4, 2, 1, 3}, []int{1, 2, 3, 4}},
		{[]int{4, 2, 3, 1}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 1, 2}, []int{1, 2, 3, 4}},
		{[]int{4, 3, 2, 1}, []int{1, 2, 3, 4}},
		{[]int{5, 7, 4, 6, 9, 10, 3, 8, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{5, 7, 4, 4, 9, 10, 7, 8, 2, 1}, []int{1, 2, 4, 4, 5, 7, 7, 8, 9, 10}},
	}
	for _, c := range cases {
		got := make([]int, len(c.in))
		copy(got, c.in)
		QuickSort(got)

		if !cmp.Equal(got, c.want) {
			t.Errorf("%v == %v, want %v", c.in, got, c.want)
		}
	}
}

func BenchmarkQuickSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		QuickSort(rand.Perm(100000))
	}
}
