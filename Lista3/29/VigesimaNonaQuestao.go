package main
import "fmt"

func main(){
	var n int

	for{
		fmt.Print("Informe um número inteiro qualquer: ")
		fmt.Scanln(&n)
		if n > 0 {
			break
		} else {
		fmt.Println("Valor inválido! Digite um número positivo.")
		}
	}

	soma:= 0
	for i:=1; i<=n; i++{
		soma+=i
	}
	fmt.Println(soma)
}
