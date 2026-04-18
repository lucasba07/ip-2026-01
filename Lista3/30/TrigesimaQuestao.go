package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Raio | Volume")

	for r := 0.0; r <= 20; r += 0.5 {
		volume := (4.0 / 3.0) * math.Pi * (r * r * r)
		fmt.Printf("%4.1f | %.2f\n", r, volume)
	}
}