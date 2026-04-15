package main
import "fmt"

func main (){
	var (
		base int
		expoente int
	)

	fmt.Print("Informe a base da potência: ")
	fmt.Scanln(&base)
	fmt.Print("Informe o expoente da potência: ")
	fmt.Scanln(&expoente)

	fmt.Printf("O resultado da potenciação é %d\n", potencia(base, expoente))
}

func potencia (base int, expoente int) int{
	switch expoente {
		case 0:
			return 1
		case 1:
			return base
	}
	return base * potencia(base, expoente - 1) 
}
