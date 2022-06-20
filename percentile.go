package gf

import (
	"errors"
	"math"
)

// Percentile 获取百分位数
func Percentile[T any](
	percentile int,
	slice []T,
	less func(slice []T, i, j int) bool,
	get func(slice []T, i int) float64,
	avg func(slice []T, i, j int) float64,
) (float64, error) {
	// 验证输入
	if percentile > 100 || percentile < 1 {
		return 0, errors.New("error! percentile should between 1 to 99")
	}
	if len(slice) <= 0 {
		return 0, errors.New("error! slice is empty")
	}
	ArraySort(slice, less)

	// 计算百分位数的索引
	var index int
	i := float64(len(slice)*percentile) / 100
	if math.Ceil(i) == math.Floor(i) {
		index = int(i - 1)
		return avg(slice, index, index+1), nil
	} else {
		index = int(math.Ceil(i)) - 1
		return get(slice, index), nil
	}
}
