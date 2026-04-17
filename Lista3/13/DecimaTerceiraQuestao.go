package main
import "fmt"

func main(){
	var h float64
	x:= 1.0

	for i:=1.0; i<=50; i++{
		h+= x / i
		x+=2
	}
	fmt.Printf("H resulta em: %.2f\n", h)
}
