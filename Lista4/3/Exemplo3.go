package main

import "fmt"

func main() {
	var tamanho int

	fmt.Print("Digite o tamanho do array: ")
	fmt.Scan(&tamanho)

	lista := make([]int, tamanho)
	for i := range lista {
		fmt.Printf("Digite o valor para a posição %d: ", i+1)
		fmt.Scan(&lista[i])
	}

	inverterLista(lista, 0, tamanho-1)
	fmt.Println("Lista invertida:", lista)
}

func inverterLista(lista []int, inicio, fim int) {
	if inicio >= fim {
		return
	}

	lista[inicio], lista[fim] = lista[fim], lista[inicio]
	inverterLista(lista, inicio+1, fim-1)
}
