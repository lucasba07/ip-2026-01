package main
import "fmt"

func main() {
	var (
		lista [5]int
		num int
	)
	for i := range lista {
		fmt.Print("Informe um valor qualquer: ")
		fmt.Scanln(&lista[i])
	}

	fmt.Printf("Informe um valor para procurar na lista: ")
	fmt.Scanln(&num)

	if buscaBinaria(lista, num) == -1 {
		fmt.Println("Valor não encontrado na lista, portanto retornou o valor ", buscaBinaria(lista, num))
	} else {
		fmt.Println("O elemento está na posição ", buscaBinaria(lista, num))
	}
	
}	

func buscaBinaria(lista [5]int, num int) int {
	primeiro := 0
	ultimo := len(lista) - 1

	for primeiro <= ultimo {
		meio := (primeiro + ultimo) / 2

		if lista[meio] == num {
			return meio
		} else if lista[meio] < num {
			primeiro = meio + 1
		} else {
			ultimo = meio - 1
		}
	}
	return -1
}
