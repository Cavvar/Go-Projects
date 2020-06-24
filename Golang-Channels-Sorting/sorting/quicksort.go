package sorting

import (
	"sync"
)

// QuickSort implements the QuickSort-Algorithm O(n * log n).
// The given slice will be sorted in place.
func QuickSort(s []int) []int {
	length := len(s)
	if length > 1 {
		var wg sync.WaitGroup

		wg.Add(1)

		go quickSortParallel(s, &wg)
		wg.Wait()
	}

	return s
}

func quickSortParallel(a []int, wg *sync.WaitGroup) {
	defer wg.Done()

	length := len(a)

	if length < 2 {
		return
	}

	pivot := a[length-1]
	leftWall := 0

	for i := 0; i < length; i++ {
		if a[i] < pivot {
			a[i], a[leftWall] = a[leftWall], a[i]
			leftWall++
		}
	}

	a[leftWall], a[length-1] = a[length-1], a[leftWall]

	wg.Add(2)

	go quickSortParallel(a[:leftWall], wg)

	go quickSortParallel(a[leftWall+1:], wg)
}
