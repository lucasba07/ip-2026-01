package main
import "fmt"

func main(){
	cont, soma:= 1.0, 0.0
	for i:= 37.0; i>=1; i--{
		soma += (i*(i+1)) / cont
		cont++
	}
	fmt.Printf("%.2f\n", soma)
}
