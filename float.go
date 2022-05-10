package gf

import (
	"fmt"
	"reflect"
	"strconv"
)

func FloatRound[F ~float64 | ~float32](floatValue F, decimalCount int) F {
	bitSize := 64
	if reflect.TypeOf(floatValue).Kind() == reflect.Float32 {
		bitSize = 32
	}
	res, _ := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", decimalCount), floatValue), bitSize)
	return F(res)
}
