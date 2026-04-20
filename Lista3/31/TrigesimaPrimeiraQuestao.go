package main
import "fmt"

func main(){
	var(
		qntdGraos uint64 = 1
		total uint64 = 0
	)

	for i:=1; i<=64; i++{
		fmt.Printf("Pode-se colocar %d grãos no %dº quadro\n", qntdGraos, i)
		total += qntdGraos
		qntdGraos *= 2
	}
	fmt.Printf("No total, pode-se colocar %d grãos\n", total)
}
