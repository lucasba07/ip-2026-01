package main
import "fmt"

func main(){
	soma:= 0.0
	cont:=100.0
	for i:=0; i<20; i++{
		soma += (cont / float64(fatorial(i)))
		cont--
	}
	fmt.Printf("%.2f\n", soma)
}

func fatorial(n int) int {
	if n == 0{return 1}
	return n * fatorial(n-1)
}
