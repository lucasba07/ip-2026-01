package main
import "fmt"

func main(){
	var(
		b int
		n int
	)

	for{
		fmt.Print("Informe o valor da base: ")
		fmt.Scanln(&b)

		fmt.Print("Informe o valor do expoente: ")
		fmt.Scanln(&n)
		
		if n > 1 && b >= 2{break}
		fmt.Println("VALORES INVÁLIDOS!")
	}

	mult := 1
	for i:= 0; i < n; i++{
		mult *= b
	}
	fmt.Printf("Resultado: %v\n", mult)
}
