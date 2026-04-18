package main

import (
	"fmt"
	"math"
)

func main(){
	var(
		newA, senA float64
	) 
	fmt.Println("ÂNGULO | SENO")
	for a:= 0.0; a<=6.3; a+=0.1{
		newA = math.Round(a*10) / 10
		senA = newA - (math.Pow(newA,3)/6) + (math.Pow(newA,5)/120) - (math.Pow(newA,7)/5040)
		fmt.Printf("   %.1f | %.2f\n", newA, senA)
	}
}