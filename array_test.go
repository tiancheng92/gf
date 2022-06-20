package gf

import (
	"reflect"
	"testing"
)

func TestArrayContainsInt(t *testing.T) {
	type args struct {
		array []int
		val   int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{array: []int{1, 2, 3}, val: 2}, want: true},
		{name: "failed", args: args{array: []int{1, 2, 3}, val: 12}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayContains(tt.args.array, tt.args.val); got != tt.want {
				t.Errorf("ArrayContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayContainsFloat(t *testing.T) {
	type args struct {
		array []float64
		val   float64
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{array: []float64{1.1, 2.0, 3.3}, val: 2}, want: true},
		{name: "failed", args: args{array: []float64{1.3, 21.2, 3}, val: 12}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayContains(tt.args.array, tt.args.val); got != tt.want {
				t.Errorf("ArrayContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayContainsString(t *testing.T) {
	type args struct {
		array []string
		val   string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{array: []string{"A", "B", "C"}, val: "B"}, want: true},
		{name: "failed", args: args{array: []string{"A", "B", "C"}, val: "b"}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayContains(tt.args.array, tt.args.val); got != tt.want {
				t.Errorf("ArrayContains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDeduplicateInt(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success", args: args{array: []int{1, 2, 3, 1, 2, 3}}, want: []int{1, 2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDeduplicate(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDeduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDeduplicateFloat(t *testing.T) {
	type args struct {
		array []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "float", args: args{array: []float64{1.0, 2.2, 3.1, 1}}, want: []float64{1.0, 2.2, 3.1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDeduplicate(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDeduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDeduplicateString(t *testing.T) {
	type args struct {
		array []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "success", args: args{array: []string{"A", "b", "C", "b"}}, want: []string{"A", "b", "C"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDeduplicate(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDeduplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayEqual(t *testing.T) {
	type args struct {
		a []int
		b []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{a: []int{1, 2, 3}, b: []int{1, 2, 3}}, want: true},
		{name: "failed", args: args{a: []int{1, 2, 3}, b: []int{1, 2, 4}}, want: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayEqual(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("ArrayEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayIntersectInt(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success", args: args{array1: []int{1, 2, 2, 2, 3}, array2: []int{2, 3, 4}}, want: []int{2, 3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayIntersect(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayIntersectFloat(t *testing.T) {
	type args struct {
		array1 []float64
		array2 []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "success", args: args{array1: []float64{1, 2.2, 3}, array2: []float64{2, 3.0, 4}}, want: []float64{3}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayIntersect(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayIntersectString(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "success", args: args{array1: []string{"a", "b", "c"}, array2: []string{"x", "a", "z"}}, want: []string{"a"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayIntersect(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayIntersect() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUnionInt(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success", args: args{array1: []int{1, 2, 4}, array2: []int{3, 4, 5}}, want: []int{1, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUnion(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUnionFloat(t *testing.T) {
	type args struct {
		array1 []float64
		array2 []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "success", args: args{array1: []float64{1, 1, 1, 1, 2, 4}, array2: []float64{3, 4, 5}}, want: []float64{1, 2, 4, 3, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUnion(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayUnionString(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "success", args: args{array1: []string{"a", "b", "z"}, array2: []string{"c", "z", "e"}}, want: []string{"a", "b", "z", "c", "e"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayUnion(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayUnion() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDifferenceInt(t *testing.T) {
	type args struct {
		array1 []int
		array2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{name: "success", args: args{
			array1: []int{1, 1, 1, 2},
			array2: []int{2, 3},
		}, want: []int{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDifference(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDifferenceFloat(t *testing.T) {
	type args struct {
		array1 []float64
		array2 []float64
	}
	tests := []struct {
		name string
		args args
		want []float64
	}{
		{name: "success", args: args{
			array1: []float64{1, 2},
			array2: []float64{2, 3},
		}, want: []float64{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDifference(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayDifferenceString(t *testing.T) {
	type args struct {
		array1 []string
		array2 []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{name: "success", args: args{
			array1: []string{"1", "2"},
			array2: []string{"2", "3"},
		}, want: []string{"1"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDifference(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayFilter(t *testing.T) {
	type args struct {
		array     []int
		predicate func(int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8},
				predicate: func(i int) bool {
					return i%2 == 0
				},
			},
			want: []int{2, 4, 6, 8},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayFilter(tt.args.array, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayFilter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArrayReverse(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				[]int{1, 2, 3, 4, 5},
			},
			want: []int{5, 4, 3, 2, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayReverse(tt.args.array); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayReverse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestArraySort(t *testing.T) {
	type args struct {
		array []int
		less  func(array []int, i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				array: []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5},
				less: func(array []int, i, j int) bool {
					return array[i] < array[j]
				},
			},
			want: []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 10, 14, 14, 14, 14, 14, 14, 14, 14, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArraySort(tt.args.array, tt.args.less)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("ArraySort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}

func TestArrayInsertionSort(t *testing.T) {
	type args struct {
		array []int
		less  func(array []int, i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				array: []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5},
				less: func(array []int, i, j int) bool {
					return array[i] < array[j]
				},
			},
			want: []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 10, 14, 14, 14, 14, 14, 14, 14, 14, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArrayInsertionSort(tt.args.array, tt.args.less)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("ArrayInsertionSort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}

func TestArrayBubbleSort(t *testing.T) {
	type args struct {
		array []int
		less  func(array []int, i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				array: []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5},
				less: func(array []int, i, j int) bool {
					return array[i] < array[j]
				},
			},
			want: []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 10, 14, 14, 14, 14, 14, 14, 14, 14, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArrayBubbleSort(tt.args.array, tt.args.less)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("ArrayBubbleSort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}

func TestArrayHeapSort(t *testing.T) {
	type args struct {
		array []int
		less  func(array []int, i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				array: []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5},
				less: func(array []int, i, j int) bool {
					return array[i] < array[j]
				},
			},
			want: []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 10, 14, 14, 14, 14, 14, 14, 14, 14, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArrayHeapSort(tt.args.array, tt.args.less)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("ArrayHeapSort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}

func TestArrayQuickSort(t *testing.T) {
	type args struct {
		array []int
		less  func(array []int, i, j int) bool
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "success",
			args: args{
				array: []int{10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 99, 2, 3, 14, 5, 10, 2, 3, 14, 5, 10, 2, 3, 14, 5},
				less: func(array []int, i, j int) bool {
					return array[i] < array[j]
				},
			},
			want: []int{2, 2, 2, 2, 2, 2, 2, 2, 3, 3, 3, 3, 3, 3, 3, 3, 5, 5, 5, 5, 5, 5, 5, 5, 10, 10, 10, 10, 10, 10, 10, 10, 14, 14, 14, 14, 14, 14, 14, 14, 99},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ArrayQuickSort(tt.args.array, tt.args.less)
			if !reflect.DeepEqual(tt.args.array, tt.want) {
				t.Errorf("ArrayQuickSort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}

func TestArraySplit(t *testing.T) {
	type args struct {
		array []int
		step  int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "success",
			args: args{
				array: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 0},
				step:  2,
			},
			want: [][]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}, {9, 0}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if !reflect.DeepEqual(ArraySplit(tt.args.array, tt.args.step), tt.want) {
				t.Errorf("ArrayQuickSort() = %v, want %v", tt.args.array, tt.want)
			}
		})
	}
}
