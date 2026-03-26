package main

import "fmt"

func main() {
	var destino int
	var retorno int

	fmt.Print("Escolha o destino:\n1 - Região Norte\n2 - Região Nordeste\n3 - Região Centro-Oeste\n4 - Região Sul\nDigite a opção: ")
	fmt.Scan(&destino)

	if destino < 1 || destino > 4 {
		fmt.Println("Erro: Destino inválido!")
		return
	}

	fmt.Print("Incluir retorno? (1 - Sim, 2 - Não): ")
	fmt.Scan(&retorno)

	if retorno != 1 && retorno != 2 {
		fmt.Println("Erro: Opção de retorno inválida!")
		return
	}

	var preco float64
	var nomeDestino string

	switch destino {
	case 1:
		nomeDestino = "Região Norte"
		if retorno == 1 {
			preco = 900.00
		} else {
			preco = 500.00
		}
	case 2:
		nomeDestino = "Região Nordeste"
		if retorno == 1 {
			preco = 650.00
		} else {
			preco = 350.00
		}
	case 3:
		nomeDestino = "Região Centro-Oeste"
		if retorno == 1 {
			preco = 600.00
		} else {
			preco = 350.00
		}
	case 4:
		nomeDestino = "Região Sul"
		if retorno == 1 {
			preco = 550.00
		} else {
			preco = 300.00
		}
	}

	tipoPassagem := "Ida"
	if retorno == 1 {
		tipoPassagem = "Ida e Volta"
	}

	fmt.Printf("Destino: %s\n", nomeDestino)
	fmt.Printf("Tipo: %s\n", tipoPassagem)
	fmt.Printf("Preço: R$ %.2f\n", preco)
}
