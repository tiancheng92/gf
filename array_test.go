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

func TestArrayContainsAny(t *testing.T) {
	type args struct {
		array []any
		val   any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "any", args: args{array: []any{1, "b", nil, struct {
			Name string
		}{Name: "test"}}, val: struct {
			Name string
		}{Name: "test"}}, want: true},
		{name: "any", args: args{array: []any{1, "b", nil, struct {
			Name string
		}{Name: "test"}}, val: struct {
			Name string
		}{Name: "test1"}}, want: false},
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

func TestArrayDeduplicateAny(t *testing.T) {
	type args struct {
		array []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "success", args: args{array: []any{1, "b", nil, nil, "1"}}, want: []any{1, "b", nil, "1"}},
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
		a []any
		b []any
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "success", args: args{a: []any{1, 2, 3}, b: []any{1, 2, 3}}, want: true},
		{name: "failed", args: args{a: []any{1, 2, 3}, b: []any{1, 2, ""}}, want: false},
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

func TestArrayIntersectAny(t *testing.T) {
	type args struct {
		array1 []any
		array2 []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "success", args: args{array1: []any{1, "a", 2.2, "b", "c"}, array2: []any{1, "x", "a", "z", 2.2}}, want: []any{1, "a", 2.2}},
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

func TestArrayUnionAny(t *testing.T) {
	type args struct {
		array1 []any
		array2 []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "success", args: args{array1: []any{"a", 2, "z"}, array2: []any{"c", "z", 5.5}}, want: []any{"a", 2, "z", "c", 5.5}},
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

func TestArrayDifferenceAny(t *testing.T) {
	type args struct {
		array1 []any
		array2 []any
	}
	tests := []struct {
		name string
		args args
		want []any
	}{
		{name: "success", args: args{
			array1: []any{"1", 2, 3.0},
			array2: []any{1, 2, "3"},
		}, want: []any{"1", 3.0}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayDifference(tt.args.array1, tt.args.array2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
