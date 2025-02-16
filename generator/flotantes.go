package generator

import (
	"math/rand"
)

func generarNumeroReal() float64 {
	min := 0.00
	max := 999999999.99
	return min + rand.Float64()*(max-min)
}
