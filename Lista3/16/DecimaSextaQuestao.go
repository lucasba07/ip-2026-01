package main
import "fmt"

func main(){
	var (
		n1, n2, tamanho int
	)

	fmt.Print("Informe o primeiro número inteiro: ")
	fmt.Scanln(&n1)

	fmt.Print("Informe o segundo número inteiro: ")
	fmt.Scanln(&n2)

	fmt.Print("Informe o tamanho da sequência: ")
	fmt.Scanln(&tamanho)

	if tamanho < 3 { return }
	lista:= make([]int, tamanho)
	lista[0] = n1
	lista[1] = n2

	for i:=2; i < tamanho; i++{
		if (i+1) % 2 == 0{lista[i] = lista[i-1] - lista[i-2]}
		if (i+1) % 2 != 0{lista[i] = lista[i-1] + lista[i-2]}
	}
	fmt.Println(lista)
}

