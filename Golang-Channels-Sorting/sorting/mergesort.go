package sorting

// MergeSort is a sorting algorithm with O(n * log n).
// A new sorted slice will be returned after the algorithm has finished.
func MergeSort(s []int) []int {
	length := len(s)
	resultSlice := make([]int, length)
	copy(resultSlice, s)

	if length == 1 {
		return resultSlice
	}

	ch := make(chan int)
	go mergeSortParallel(resultSlice, ch)

	for i := 0; i < length; i++ {
		resultSlice[i] = <-ch
	}

	return resultSlice
}

// mergeSortParallel is the actual function doing all the work for MergeSort.
// s is the slice that needs to be sorted.
// c is a channel that is used to communicate with the caller.
func mergeSortParallel(s []int, c chan<- int) {
	defer close(c)

	length := len(s)

	if length == 0 {
		return
	}

	if length == 1 {
		c <- s[0]
		return
	}

	middle := length / 2
	leftChannel := make(chan int)
	rightChannel := make(chan int)

	go mergeSortParallel(s[:middle], leftChannel)

	go mergeSortParallel(s[middle:], rightChannel)

	leftVal, isOpenL := <-leftChannel
	rightVal, isOpenR := <-rightChannel

	for isOpenL || isOpenR {
		switch {
		case !isOpenR:
			c <- leftVal
			leftVal, isOpenL = <-leftChannel
		case !isOpenL:
			c <- rightVal
			rightVal, isOpenR = <-rightChannel
		default:
			if leftVal < rightVal {
				c <- leftVal
				leftVal, isOpenL = <-leftChannel
			} else {
				c <- rightVal
				rightVal, isOpenR = <-rightChannel
			}
		}
	}
}
