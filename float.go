package gf

import (
	"fmt"
	"reflect"
	"strconv"
)

type Float interface {
	float64 | float32
}

func FloatKeepDecimal[F Float](floatValue F, decimalCount int) F {
	bitSize := 64
	if reflect.TypeOf(floatValue).Kind() == reflect.Float32 {
		bitSize = 32
	}
	res, _ := strconv.ParseFloat(fmt.Sprintf(fmt.Sprintf("%%.%df", decimalCount), floatValue), bitSize)
	return F(res)
}
