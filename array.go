package gf

import (
	"reflect"
)

// ArrayDeduplicate 数组去重
func ArrayDeduplicate[T any](array []T) []T {
	arr := make([]T, 0)
	set := NewSet[T]()
	for i := range array {
		if !set.Has(array[i]) {
			set.Add(array[i])
			arr = append(arr, array[i])
		}
	}
	return arr
}

// ArrayEqual 数组比较
func ArrayEqual[T any](array1, array2 []T) bool {
	return reflect.DeepEqual(array1, array2)
}

// ArrayContains 数组包含
func ArrayContains[T any](array []T, val T) bool {
	for i := range array {
		if reflect.DeepEqual(array[i], val) {
			return true
		}
	}
	return false
}

// ArrayIntersect 数组交集
func ArrayIntersect[T any](array1, array2 []T) []T {
	arr := make([]T, 0)
	set := NewSet[T]()
	set.Add(array1...)
	for i := range array2 {
		if set.Has(array2[i]) {
			arr = append(arr, array2[i])
		}
	}
	return arr
}

// ArrayUnion 数组并集
func ArrayUnion[T any](array1, array2 []T) []T {
	arr := make([]T, 0)
	set := NewSet[T]()
	arr = ArrayDeduplicate(array1)
	set.Add(array1...)
	for i := range array2 {
		if !set.Has(array2[i]) {
			arr = append(arr, array2[i])
		}
	}
	return arr
}

// ArrayDifference 数组差集
func ArrayDifference[T any](array1, array2 []T) []T {
	arr := make([]T, 0)
	set := NewSet[T]()
	inter := ArrayIntersect(array1, array2)
	set.Add(inter...)
	array1 = ArrayDeduplicate(array1)
	for i := range array1 {
		if !set.Has(array1[i]) {
			arr = append(arr, array1[i])
		}
	}
	return arr
}

// ArrayFilter 数组过滤
func ArrayFilter[T any](array []T, predicate func(T) bool) []T {
	arr := make([]T, 0)
	for i := range array {
		if predicate(array[i]) {
			arr = append(arr, array[i])
		}
	}
	return arr
}

// ArrayReverse 数组反转
func ArrayReverse[T any](array []T) []T {
	arr := make([]T, len(array))
	for i := range array {
		arr[len(array)-i-1] = array[i]
	}
	return arr
}

// ArrayBubbleSort 数组冒泡排序
func ArrayBubbleSort[T any](array []T, less func(array []T, i, j int) bool) []T {
	for i := 0; i <= len(array)-1; i++ {
		for j := i; j <= len(array)-1; j++ {
			if !less(array, i, j) {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
	return array
}

// ArrayInsertionSort 数组插入排序
func ArrayInsertionSort[T any](array []T, less func(array []T, i, j int) bool) []T {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && less(array, j, j-1); j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
	return array
}

// ArrayQuickSort 数组快速排序
func ArrayQuickSort[T any](array []T, less func(array []T, i, j int) bool) []T {
	n := len(array)
	return quickSort(array, 0, n, maxDepth(n), less)
}

// ArrayHeapSort 数组堆排序
func ArrayHeapSort[T any](array []T, less func(array []T, i, j int) bool) []T {
	n := len(array)
	return heapSort(array, 0, n, less)
}

func maxDepth(n int) int {
	var depth int
	for i := n; i > 0; i >>= 1 {
		depth++
	}
	return depth * 2
}

func siftDown[T any](array []T, lo, hi, first int, less func(array []T, i, j int) bool) []T {
	root := lo
	for {
		child := 2*root + 1
		if child >= hi {
			break
		}
		if child+1 < hi && less(array, first+child, first+child+1) {
			child++
		}
		if !less(array, first+root, first+child) {
			return array
		}
		array[first+root], array[first+child] = array[first+child], array[first+root]
		root = child
	}
	return array
}

func heapSort[T any](array []T, a, b int, less func(array []T, i, j int) bool) []T {
	first := a
	lo := 0
	hi := b - a

	// Build heap with greatest element at top.
	for i := (hi - 1) / 2; i >= 0; i-- {
		array = siftDown(array, i, hi, first, less)
	}

	// Pop elements, largest first, into end of array.
	for i := hi - 1; i >= 0; i-- {
		array[first], array[first+i] = array[first+i], array[first]
		array = siftDown(array, lo, i, first, less)
	}
	return array
}

func medianOfThree[T any](array []T, m1, m0, m2 int, less func(array []T, i, j int) bool) []T {
	// sort 3 elements
	if less(array, m1, m0) {
		array[m1], array[m0] = array[m0], array[m1]
	}
	// array[m0] <= array[m1]
	if less(array, m2, m1) {
		array[m1], array[m2] = array[m2], array[m1]
		// array[m0] <= array[m2] && array[m1] < array[m2]
		if less(array, m1, m0) {
			array[m1], array[m0] = array[m0], array[m1]
		}
	}
	// now array[m0] <= array[m1] <= array[m2]
	return array
}

func doPivot[T any](array []T, lo, hi int, less func(array []T, i, j int) bool) ([]T, int, int) {
	m := int(uint(lo+hi) >> 1) // Written like this to avoid integer overflow.
	if hi-lo > 40 {
		// Tukey's ``Ninther,'' median of three medians of three.
		s := (hi - lo) / 8
		array = medianOfThree(array, lo, lo+s, lo+2*s, less)
		array = medianOfThree(array, m, m-s, m+s, less)
		array = medianOfThree(array, hi-1, hi-1-s, hi-1-2*s, less)
	}
	array = medianOfThree(array, lo, m, hi-1, less)

	// Invariants are:
	//	array[lo] = pivot (set up by ChoosePivot)
	//	array[lo < i < a] < pivot
	//	array[a <= i < b] <= pivot
	//	array[b <= i < c] unexamined
	//	array[c <= i < hi-1] > pivot
	//	array[hi-1] >= pivot
	pivot := lo
	a, c := lo+1, hi-1

	for ; a < c && less(array, a, pivot); a++ {
	}
	b := a
	for {
		for ; b < c && !less(array, pivot, b); b++ { // array[b] <= pivot
		}
		for ; b < c && less(array, pivot, c-1); c-- { // array[c-1] > pivot
		}
		if b >= c {
			break
		}
		// array[b] > pivot; array[c-1] <= pivot
		array[b], array[c-1] = array[c-1], array[b]
		b++
		c--
	}
	// If hi-c<3 then there are duplicates (by property of median of nine).
	// Let's be a bit more conservative, and set border to 5.
	protect := hi-c < 5
	if !protect && hi-c < (hi-lo)/4 {
		// Lets test some points for equality to pivot
		dups := 0
		if !less(array, pivot, hi-1) { // array[hi-1] = pivot
			array[c], array[hi-1] = array[hi-1], array[c]
			c++
			dups++
		}
		if !less(array, b-1, pivot) { // array[b-1] = pivot
			b--
			dups++
		}
		// m-lo = (hi-lo)/2 > 6
		// b-lo > (hi-lo)*3/4-1 > 8
		// ==> m < b ==> array[m] <= pivot
		if !less(array, m, pivot) { // array[m] = pivot
			array[m], array[b-1] = array[b-1], array[m]
			b--
			dups++
		}
		// if at least 2 points are equal to pivot, assume skewed distribution
		protect = dups > 1
	}
	if protect {
		// Protect against a lot of duplicates
		// Add invariant:
		//	array[a <= i < b] unexamined
		//	array[b <= i < c] = pivot
		for {
			for ; a < b && !less(array, b-1, pivot); b-- { // array[b] == pivot
			}
			for ; a < b && less(array, a, pivot); a++ { // array[a] < pivot
			}
			if a >= b {
				break
			}
			// array[a] == pivot; array[b-1] < pivot
			array[a], array[b-1] = array[b-1], array[a]
			a++
			b--
		}
	}
	// Swap pivot into middle
	array[pivot], array[b-1] = array[b-1], array[pivot]
	return array, b - 1, c
}

func quickSort[T any](array []T, a, b, maxDepth int, less func(array []T, i, j int) bool) []T {
	for b-a > 12 { // Use ShellSort for slices <= 12 elements
		if maxDepth == 0 {
			return heapSort(array, a, b, less)
		}
		maxDepth--
		data, mlo, mhi := doPivot(array, a, b, less)
		// Avoiding recursion on the larger subproblem guarantees
		// a stack depth of at most lg(b-a).
		if mlo-a < b-mhi {
			data = quickSort(data, a, mlo, maxDepth, less)
			a = mhi // i.e., quickSort(array, mhi, b)
		} else {
			data = quickSort(data, mhi, b, maxDepth, less)
			b = mlo // i.e., quickSort(array, a, mlo)
		}
	}
	if b-a > 1 {
		// Do ShellSort pass with gap 6
		// It could be written in this simplified form cause b-a <= 12
		for i := a + 6; i < b; i++ {
			if less(array, i, i-6) {
				array[i], array[i-6] = array[i-6], array[i]
			}
		}
		array = ArrayInsertionSort(array, less)
	}
	return array
}
