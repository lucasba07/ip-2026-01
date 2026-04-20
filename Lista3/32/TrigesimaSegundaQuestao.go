package main
import "fmt"

func main(){
	c

	fmt.Println(multiplicar(n1, n2))
}

func multiplicar(n1, n2 int) int {
	mult := 0
	negativo := false

	if n1 < 0 {
		n1 = -n1
		negativo = !negativo
	} 
	
	if n2 < 0 {
		n2 = -n2
		negativo = !negativo
	}

	for i := 0; i < n2; i++ {mult += n1}

	if negativo {mult = -mult}
	return mult
}
