package main
import "fmt"

func main (){
	soma := 0
	qntd := 0

	for x:= 50; x <= 70; x+=2{
		soma += x
		qntd++
	}
	media := soma / qntd
	fmt.Printf("SOMA: %v | MÉDIA: %v\n", soma, media)
}