package gf

import "reflect"

// ArrayDeduplicate 数组去重
func ArrayDeduplicate[T any](array []T) []T {
	var (
		arr = make([]T, 0)
		m   = make(map[any]struct{})
	)
	for i := range array {
		if _, ok := m[array[i]]; !ok {
			m[array[i]] = struct{}{}
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
	var (
		set = make(map[any]struct{})
		res = make([]T, 0)
	)

	for i := range array1 {
		set[array1[i]] = struct{}{}
	}

	for i := range array2 {
		if _, ok := set[array2[i]]; ok {
			res = append(res, array2[i])
		}
	}
	return res
}

// ArrayUnion 数组并集
func ArrayUnion[T any](array1, array2 []T) []T {
	var (
		set = make(map[any]struct{})
		res = make([]T, 0)
	)

	res = ArrayDeduplicate(array1)
	for i := range array1 {
		set[array1[i]] = struct{}{}
	}

	for i := range array2 {
		if _, ok := set[array2[i]]; !ok {
			res = append(res, array2[i])
		}
	}
	return res
}

// ArrayDifference 数组差集
func ArrayDifference[T any](array1, array2 []T) []T {
	var (
		set = make(map[any]struct{})
		res = make([]T, 0)
	)

	inter := ArrayIntersect(array1, array2)
	for i := range inter {
		set[inter[i]] = struct{}{}
	}

	array1 = ArrayDeduplicate(array1)
	for i := range array1 {
		if _, ok := set[array1[i]]; !ok {
			res = append(res, array1[i])
		}
	}
	return res
}
