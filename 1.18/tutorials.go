package tutorials

import (
	"fmt"

	"github.com/shopspring/decimal"
	"github.com/yiGmMk/go-tool/convert"
)

type Number interface {
	int64 | int32 | float32 | float64 | int | int16 | int8
}

func FuzzingTest[V Number](val V) (string, error) {
	in := decimal.Zero
	return convert.Number2ChineseYUAN(in, true)
}

// map key has to be comparable
func Generic[K comparable, V int64 | float64 | int](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}

func Add() {
	fmt.Printf("Generic Sums: %v and %v\n",
		Generic[string, int64](map[string]int64{"a": 1, "b": 2, "c": 3}),
		Generic[string, float64](map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}))

	fmt.Printf("Generic Sums: %v and %v\n",
		Generic(map[string]int64{"a": 1, "b": 2, "c": 3}),
		Generic(map[string]float64{"a": 1.1, "b": 2.2, "c": 3.3}))
}

func Generic2[K comparable, V Number](m map[K]V) V {
	var s V
	for _, v := range m {
		s += v
	}
	return s
}
