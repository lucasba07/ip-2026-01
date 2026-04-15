package main

import "fmt"

type Pessoa struct {
	idade  int
	altura float64
	peso   float64
}

func main() {
	var pessoas []Pessoa
	var continuar int

	for {
		var p Pessoa
		fmt.Print("Informe a idade: ")
		fmt.Scanln(&p.idade)

		fmt.Print("Informe a altura: ")
		fmt.Scanln(&p.altura)

		fmt.Print("Informe o peso: ")
		fmt.Scanln(&p.peso)

		pessoas = append(pessoas, p)

		fmt.Print("Deseja continuar? (1 - Sim, outro valor - Não): ")
		fmt.Scanln(&continuar)
		if continuar != 1 {
			break
		}
	}
	
	var countIdade50 int
	var somaAltura1020 float64
	var countAltura1020 int
	var countPeso40 int

	for _, p := range pessoas {
		if p.idade > 50 {
			countIdade50++
		}
		if p.idade >= 10 && p.idade <= 20 {
			somaAltura1020 += p.altura
			countAltura1020++
		}
		if p.peso < 40 {
			countPeso40++
		}
	}

	fmt.Printf("Quantidade de pessoas com idade superior a 50 anos: %d\n", countIdade50)

	if countAltura1020 > 0 {
		mediaAltura := somaAltura1020 / float64(countAltura1020)
		fmt.Printf("Média das alturas das pessoas com idade entre 10 e 20 anos: %.2f\n", mediaAltura)
	} else {
		fmt.Println("Não há pessoas com idade entre 10 e 20 anos.")
	}

	if len(pessoas) > 0 {
		porcentagemPeso := (float64(countPeso40) / float64(len(pessoas))) * 100
		fmt.Printf("Porcentagem de pessoas com peso inferior a 40 quilos: %.2f%%\n", porcentagemPeso)
	} else {
		fmt.Println("Nenhuma pessoa foi cadastrada.")
	}
}
