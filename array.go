package gf

import "reflect"

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
