package main

import (
	"fmt"
	"math"
)

func main(){
	var x float64

	fmt.Print("Informe o valor de x: ")
	fmt.Scanln(&x)

	cos:= 1.0
	cont:=2
	for i:=1; i<=20; i++{
		if i%2==0{
			cos+= (potencia(x,cont)/(fatorial(cont)))
		} else{
			cos-= (potencia(x,cont)/(fatorial(cont)))
		}
		cont+=2
	}
	outroCos:= math.Cos(x)
	diferenca:= cos - outroCos

	fmt.Printf("Cos(série): %.5f\n", cos)
	fmt.Printf("Cos(math):  %.5f\n", outroCos)
	fmt.Printf("Diferença:  %.10f\n", diferenca)
}

func potencia(x float64, e int) float64{
	if e == 0{return 1}
	return x * potencia(x, e-1)
}

func fatorial(x int) float64 {
	if x == 0 {
		return 1
	}
	return float64(x) * fatorial(x-1)
}