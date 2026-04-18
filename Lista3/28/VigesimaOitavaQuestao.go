package main

import (
	"fmt"
	"math"
)

func main() {
	serie := 0.0
	cont := 1

	for i := 1; i <= 51; i++ {
		if i%2 == 0 {
			serie -= 1.0 / potencia(cont, 3)
		} else {
			serie += 1.0 / potencia(cont, 3)
		}
		cont += 2
	}
	pi := math.Cbrt(serie * 32)
	fmt.Printf("Valor de π: %.5f\n", pi)
}

func potencia(n int, e int) float64 {
	if e == 0 {
		return 1
	}
	return float64(n) * potencia(n, e-1)
}