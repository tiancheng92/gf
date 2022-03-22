package main

import (
	"fmt"
	"gf"
)

type T struct {
	N int
}

func main() {
	fmt.Println("Array Contains")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayContains([]int{1, 2, 3}, 1))
	fmt.Println(gf.ArrayContains([]float64{1.2, 3.4, 5.6}, 7.8))
	fmt.Println(gf.ArrayContains([]string{"a", "b", "c"}, "c"))
	fmt.Println(gf.ArrayContains([]T{{1}, {2}, {3}}, T{2}))
	fmt.Println("")

	fmt.Println("Array Deduplicate")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayDeduplicate([]int{1, 2, 3, 4, 1, 2, 4, 3}))
	fmt.Println(gf.ArrayDeduplicate([]float64{1.1, 2.2, 3.3, 4.4, 1.1, 2.1, 4.1, 3.3}))
	fmt.Println(gf.ArrayDeduplicate([]string{"a", "b", "c", "a", "b", "c"}))
	fmt.Println(gf.ArrayDeduplicate([]T{{1}, {2}, {2}}))
	fmt.Println("")

	fmt.Println("Array Equal")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(gf.ArrayEqual([]float64{1, 2, 3}, []float64{1, 2, 3, 4.0}))
	fmt.Println(gf.ArrayEqual([]string{"a", "b", "c", "a", "b", "c"}, []string{"a", "b", "c", "a", "c", "b"}))
	fmt.Println(gf.ArrayEqual([]T{{1}, {2}, {2}}, []T{{1}, {2}, {2}}))
	fmt.Println("")

	fmt.Println("Array Intersect")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayIntersect([]int{1, 2}, []int{2, 3}))
	fmt.Println(gf.ArrayIntersect([]float64{1, 2, 3}, []float64{1, 2.0, 3, 4.0}))
	fmt.Println(gf.ArrayIntersect([]string{"a", "b"}, []string{"a", "b", "c"}))
	fmt.Println(gf.ArrayIntersect([]T{{1}, {2}, {2}}, []T{{-1}, {12}, {2}}))
	fmt.Println("")

	fmt.Println("Array Union")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayUnion([]int{1, 2}, []int{2, 3}))
	fmt.Println(gf.ArrayUnion([]float64{1, 2, 3, 5}, []float64{1, 2.0, 3, 4.0}))
	fmt.Println(gf.ArrayUnion([]string{"a", "b", "z"}, []string{"a", "b", "c"}))
	fmt.Println(gf.ArrayUnion([]T{{1}, {2}, {2}}, []T{{-1}, {12}, {2}}))
	fmt.Println("")

	fmt.Println("Array Difference")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayDifference([]int{1, 2}, []int{2, 3}))
	fmt.Println(gf.ArrayDifference([]float64{1, 2, 3, 5, 10}, []float64{1, 2.0, 3, 4.0}))
	fmt.Println(gf.ArrayDifference([]string{"a", "b", "z", "y"}, []string{"a", "b", "c"}))
	fmt.Println(gf.ArrayDifference([]T{{1}, {2}, {2}}, []T{{-1}, {12}, {2}}))
	fmt.Println("")

	fmt.Println("Array Difference")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayFilter([]int{1, 2, 3, 4, 5}, func(v int) bool {
		return v%2 == 0
	}))
	fmt.Println(gf.ArrayFilter([]float64{1, 2, 3, 4, 5}, func(v float64) bool {
		return v > 3
	}))
	fmt.Println(gf.ArrayFilter([]string{"a", "b", "c", "d"}, func(v string) bool {
		return v != "c"
	}))
	fmt.Println(gf.ArrayFilter([]T{{1}, {2}, {2}}, func(v T) bool {
		return v.N%2 == 0
	}))
	fmt.Println("")

	fmt.Println("Array Reverse")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayReverse([]int{1, 2, 3, 4, 5}))
	fmt.Println(gf.ArrayReverse([]float64{1, 2, 3, 4, 5}))
	fmt.Println(gf.ArrayReverse([]string{"a", "b", "c", "d"}))
	fmt.Println(gf.ArrayReverse([]T{{1}, {2}, {2}}))
	fmt.Println("")

	fmt.Println("Array Insertion Sort")
	fmt.Println("-----------------------------")
	arr1 := []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5}
	gf.ArrayInsertionSort(arr1, func(arr []int, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr1)
	arr2 := []string{"z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "k", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d"}
	gf.ArrayInsertionSort(arr2, func(arr []string, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr2)
	arr3 := []T{{10}, {-9}, {7}, {13}, {11}, {1}, {-6}, {71}, {8}, {31}, {-11}, {1}}
	gf.ArrayInsertionSort(arr3, func(arr []T, i, j int) bool {
		return arr[i].N < arr[j].N
	})
	fmt.Println(arr3)
	fmt.Println("")

	fmt.Println("Array Sort")
	fmt.Println("-----------------------------")
	arr1 = []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5}
	gf.ArraySort(arr1, func(arr []int, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr1)
	arr2 = []string{"z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "k", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d"}
	gf.ArraySort(arr2, func(arr []string, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr2)
	arr3 = []T{{10}, {-9}, {7}, {13}, {11}, {1}, {-6}, {71}, {8}, {31}, {-11}, {1}}
	gf.ArraySort(arr3, func(arr []T, i, j int) bool {
		return arr[i].N < arr[j].N
	})
	fmt.Println(arr3)
	fmt.Println("")

	fmt.Println("Array Bubble Sort")
	fmt.Println("-----------------------------")
	arr1 = []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5}
	gf.ArrayBubbleSort(arr1, func(arr []int, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr1)
	arr2 = []string{"z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "k", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d"}
	gf.ArrayBubbleSort(arr2, func(arr []string, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr2)
	arr3 = []T{{10}, {-9}, {7}, {13}, {11}, {1}, {-6}, {71}, {8}, {31}, {-11}, {1}}
	gf.ArrayBubbleSort(arr3, func(arr []T, i, j int) bool {
		return arr[i].N < arr[j].N
	})
	fmt.Println(arr3)
	fmt.Println("")

	fmt.Println("Array Heap Sort")
	fmt.Println("-----------------------------")
	arr1 = []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5}
	gf.ArrayHeapSort(arr1, func(arr []int, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr1)
	arr2 = []string{"z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "k", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d"}
	gf.ArrayHeapSort(arr2, func(arr []string, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr2)
	arr3 = []T{{10}, {-9}, {7}, {13}, {11}, {1}, {-6}, {71}, {8}, {31}, {-11}, {1}}
	gf.ArrayHeapSort(arr3, func(arr []T, i, j int) bool {
		return arr[i].N < arr[j].N
	})
	fmt.Println(arr3)
	fmt.Println("")

	fmt.Println("Array Quick Sort")
	fmt.Println("-----------------------------")
	arr1 = []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5}
	gf.ArrayQuickSort(arr1, func(arr []int, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr1)
	arr2 = []string{"z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d", "z", "k", "z", "a", "y", "b", "c", "d", "z", "z", "a", "y", "b", "c", "d"}
	gf.ArrayQuickSort(arr2, func(arr []string, i, j int) bool {
		return arr[i] < arr[j]
	})
	fmt.Println(arr2)
	arr3 = []T{{10}, {-9}, {7}, {13}, {11}, {1}, {-6}, {71}, {8}, {31}, {-11}, {1}}
	gf.ArrayQuickSort(arr3, func(arr []T, i, j int) bool {
		return arr[i].N < arr[j].N
	})
	fmt.Println(arr3)
	fmt.Println("")

	fmt.Println("Convent String Bytes")
	fmt.Println("-----------------------------")
	fmt.Println(gf.StringToBytes("123"))
	fmt.Println(gf.BytesToString([]byte{49, 50, 51}))
	fmt.Println("")

	fmt.Println("Float Keep Decimal")
	fmt.Println("-----------------------------")
	f64 := gf.FloatRound(1.2345, 2)
	fmt.Printf("type: %T, value: %v\n", f64, f64)
	f32 := gf.FloatRound(float32(1.2345), 3)
	fmt.Printf("type: %T, value: %v\n", f32, f32)
	fmt.Println("")

	fmt.Println("String Gzip")
	fmt.Println("-----------------------------")
	str := "12345678901234567890123456789012345678901234567890123456789012345678901234567890"
	b, _ := gf.StringCreateGzip(str)
	fmt.Printf("ole size: %d, new size: %d, value: %v\n", len(str), len(b), b)
	s, _ := gf.StringParseGzip(b)
	fmt.Printf("value: %v\n", s)
	fmt.Println("")

	fmt.Println("String Join")
	fmt.Println("-----------------------------")
	fmt.Println(gf.StringJoin("a", "b", "c"))
	fmt.Println("")

	fmt.Println("URL Format")
	fmt.Println("-----------------------------")
	url, _ := gf.UrlFormat("https://u:p@www.a.com:9999////////b/c/d.e?i=2&&i=1&&&&&&&&g=2&h=3&h=3&h=3&f=4#5")
	fmt.Println(url)
}
