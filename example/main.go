package main

import (
	"fmt"
	"gf"
)

func main() {
	fmt.Println("Array Contains")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayContains([]int{1, 2, 3}, 1))
	fmt.Println(gf.ArrayContains([]float64{1.2, 3.4, 5.6}, 7.8))
	fmt.Println(gf.ArrayContains([]string{"a", "b", "c"}, "c"))
	fmt.Println(gf.ArrayContains([]any{1, 2.3, "4.5"}, "4.5"))
	fmt.Println("")

	fmt.Println("Array Deduplicate")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayDeduplicate([]int{1, 2, 3, 4, 1, 2, 4, 3}))
	fmt.Println(gf.ArrayDeduplicate([]float64{1.1, 2.2, 3.3, 4.4, 1.1, 2.1, 4.1, 3.3}))
	fmt.Println(gf.ArrayDeduplicate([]string{"a", "b", "c", "a", "b", "c"}))
	fmt.Println(gf.ArrayDeduplicate([]any{1, 2.3, "4.5", 1, struct {
		name string
	}{name: "name"}, struct {
		name string
	}{name: "name"}, "4.5"}))
	fmt.Println("")

	fmt.Println("Array Equal")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayEqual([]int{1, 2, 3}, []int{1, 2, 3}))
	fmt.Println(gf.ArrayEqual([]float64{1, 2, 3}, []float64{1, 2, 3, 4.0}))
	fmt.Println("")

	fmt.Println("Array Intersect")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayIntersect([]int{1, 2}, []int{2, 3}))
	fmt.Println(gf.ArrayIntersect([]float64{1, 2, 3}, []float64{1, 2.0, 3, 4.0}))
	fmt.Println(gf.ArrayIntersect([]string{"a", "b"}, []string{"a", "b", "c"}))
	fmt.Println("")

	fmt.Println("Array Union")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayUnion([]int{1, 2}, []int{2, 3}))
	fmt.Println(gf.ArrayUnion([]float64{1, 2, 3, 5}, []float64{1, 2.0, 3, 4.0}))
	fmt.Println(gf.ArrayUnion([]string{"a", "b", "z"}, []string{"a", "b", "c"}))
	fmt.Println("")

	fmt.Println("Array Difference")
	fmt.Println("-----------------------------")
	fmt.Println(gf.ArrayDifference([]int{1, 2}, []int{2, 3}))
	fmt.Println(gf.ArrayDifference([]float64{1, 2, 3, 5, 10}, []float64{1, 2.0, 3, 4.0}))
	fmt.Println(gf.ArrayDifference([]string{"a", "b", "z", "y"}, []string{"a", "b", "c"}))
	fmt.Println("")

	fmt.Println("Convent String Bytes")
	fmt.Println("-----------------------------")
	fmt.Println(gf.StringToBytes("123"))
	fmt.Println(gf.BytesToString([]byte{49, 50, 51}))
	fmt.Println("")

	fmt.Println("Float Keep Decimal")
	fmt.Println("-----------------------------")
	f64 := gf.FloatKeepDecimal(1.2345, 2)
	fmt.Printf("type: %T, value: %v\n", f64, f64)
	f32 := gf.FloatKeepDecimal(float32(1.2345), 3)
	fmt.Printf("type: %T, value: %v\n", f32, f32)
	fmt.Println("")

	fmt.Println("String Gzip")
	fmt.Println("-----------------------------")
	str := "12345678901234567890123456789012345678901234567890123456789012345678901234567890"
	b, _ := gf.StringCreateGzip(str)
	fmt.Printf("ole size: %d, new size: %d, value: %v\n", len(str), len(b), b)
	s, _ := gf.StringParseGzip(b)
	fmt.Printf("value: %v\n", s)
}
