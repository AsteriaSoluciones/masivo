package generator

import (
	"math/rand"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// Genera una cadena alfanumérica de longitud fija con caracteres aleatorios
func GeneraCadena(longitud int) string {
	b := make([]byte, longitud)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}
