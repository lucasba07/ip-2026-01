package main
import "fmt"

func main(){
	var n1, n2 int

	c
	resto:=n1
	quociente:=divisao(&resto, n2)
	fmt.Printf("QUOCIENTE(%d,%d) = %d | RESTO(%d,%d) = %d\n", n1, n2, quociente, n1, n2, resto)
}

func divisao(n1 *int, n2 int) int {
	cont := 0

	for *n1 >= n2 {
		*n1 -= n2
		cont++
	}

	return cont
}
