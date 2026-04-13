package main
import "fmt"

func main() {
	var(
		lista [5]int
		valor int
	) 
	for i := range lista{
		fmt.Print("Informe um valor qualquer: ")
		fmt.Scanln(&lista[i])
	}

	fmt.Printf("Informe um valor para procurar na lista: ")
	fmt.Scanln(&valor)
	
	if buscaSequencial(lista, valor) == -1 {
		fmt.Println("Valor não encontrado na lista, portanto retornou o valor", buscaSequencial(lista, valor))
	} else {
		fmt.Println("O elemento está na posição", buscaSequencial(lista, valor))
	}	
}

func buscaSequencial(lista [5]int, valor int) int {
	for i := range lista {
		if lista[i] == valor {
			return i
		}
	}
	return -1
}
