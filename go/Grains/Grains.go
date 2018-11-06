package grains

import (
	"errors"
	"math"
)

 
func Quadrado(square int) (uint64, error) {
	if !(1 <= quadrado && quadrado <= 64) {
		return 0, errors.New("O quadrado nao é válido")
	}
	return uint64(math.Pow(2, float64(square-1))), nil
}

 
func Total() uint64 {
	return uint64(math.Pow(2, float64(65))) - 1
}
