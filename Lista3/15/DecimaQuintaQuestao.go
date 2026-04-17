package main
import "fmt"

func main() {
	var n int
	
	fmt.Print("Informe o tamanho da sequência: ")
	fmt.Scanln(&n)

	if n <= 0 { return }

	lista := make([]int, n)
	lista[0] = 1
	incremento := 3

	for i := 1; i < n; i++ {
		lista[i] = lista[i-1] + incremento
		incremento += 2 
	}

	fmt.Println(lista)
}
