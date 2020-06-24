package sorting

import (
	"math/rand"
	"sort"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMergeSort(t *testing.T) {
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
		{[]int{5, 7, 4, 6, 9, 10, 3, 8, 2, 1}, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{5, 7, 4, 4, 9, 10, 7, 8, 2, 1}, []int{1, 2, 4, 4, 5, 7, 7, 8, 9, 10}},
	}
	for _, c := range cases {
		got := MergeSort(c.in)
		if !cmp.Equal(got, c.want) {
			t.Errorf("%v == %v, want %v", c.in, got, c.want)
		}
	}
}

func BenchmarkMergeSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MergeSort(rand.Perm(100000))
	}
}

func BenchmarkSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sort.Ints(rand.Perm(100000))
	}
}
