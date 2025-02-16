package generator

import (
	"math/rand"
	"time"
)

func generarFechaAleatoria() time.Time {
	min := time.Date(1920, 1, 1, 0, 0, 0, 0, time.UTC)
	max := time.Date(2090, 12, 31, 23, 59, 59, 0, time.UTC)
	delta := max.Unix() - min.Unix()

	sec := rand.Int63n(delta) + min.Unix()
	return time.Unix(sec, 0)
}
