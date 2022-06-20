package gf

import (
	"reflect"
	"sort"
)

// ArrayDeduplicate 数组去重
func ArrayDeduplicate[T comparable](array []T) []T {
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
func ArrayIntersect[T comparable](array1, array2 []T) []T {
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
func ArrayUnion[T comparable](array1, array2 []T) []T {
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
func ArrayDifference[T comparable](array1, array2 []T) []T {
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
func ArrayBubbleSort[T any](array []T, less func(array []T, i, j int) bool) {
	for i := 0; i <= len(array)-1; i++ {
		for j := i; j <= len(array)-1; j++ {
			if !less(array, i, j) {
				array[i], array[j] = array[j], array[i]
			}
		}
	}
}

// ArrayInsertionSort 数组插入排序
func ArrayInsertionSort[T any](array []T, less func(array []T, i, j int) bool) {
	for i := 1; i < len(array); i++ {
		for j := i; j > 0 && less(array, j, j-1); j-- {
			array[j], array[j-1] = array[j-1], array[j]
		}
	}
}

// ArraySort 数组排序(go原生排序)
func ArraySort[T any](array []T, less func(array []T, i, j int) bool) {
	sort.Slice(array, func(i, j int) bool {
		return less(array, i, j)
	})
}

// ArrayHeapSort 数组堆排序
func ArrayHeapSort[T any](array []T, less func(array []T, i, j int) bool) {
	siftDown := func(array []T, lo, hi, first int, less func(array []T, i, j int) bool) {
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
				return
			}
			array[first+root], array[first+child] = array[first+child], array[first+root]
			root = child
		}
	}
	first := 0
	lo := 0
	hi := len(array)
	for i := (hi - 1) / 2; i >= 0; i-- {
		siftDown(array, i, hi, first, less)
	}
	for i := hi - 1; i >= 0; i-- {
		array[first], array[first+i] = array[first+i], array[first]
		siftDown(array, lo, i, first, less)
	}
}

// ArrayQuickSort 数组快速排序
func ArrayQuickSort[T any](array []T, less func(array []T, i, j int) bool) {
	quickSort(array, 0, len(array)-1, less)
}

func quickSort[T any](array []T, left, right int, less func(array []T, i, j int) bool) {
	if left < right {
		partitionIndex := partition(array, left, right, less)
		quickSort(array, left, partitionIndex-1, less)
		quickSort(array, partitionIndex+1, right, less)
	}
}

func partition[T any](array []T, left, right int, less func(array []T, i, j int) bool) int {
	pivot := left
	index := pivot + 1
	for i := index; i <= right; i++ {
		if less(array, i, pivot) {
			array[i], array[index] = array[index], array[i]
			index++
		}
	}
	array[pivot], array[index-1] = array[index-1], array[pivot]
	return index - 1
}

// ArraySplit 把数组按步长拆分为多个数组
func ArraySplit[T any](array []T, step int) [][]T {
	var startIndex = 0
	var endIndex = 0
	var result [][]T
	length := len(array)
	for {
		endIndex = startIndex + step
		if endIndex >= length {
			result = append(result, array[startIndex:])
			return result
		}
		result = append(result, array[startIndex:endIndex])
		startIndex += step
	}
}
