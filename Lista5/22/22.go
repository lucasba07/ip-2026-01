package main
import "slices"

import "fmt"

func main(){
	conta:= make(map[int]float64)
	codigo := make([]int, 0, 10)

	for i:=0; i<10; i++{
		saldo:=0.0
		for{
			num:=0		
			fmt.Printf("Informe o código da %vº conta: ", (i+1))
			fmt.Scanln(&num)

			if !codigoExiste(num, codigo){
				codigo = append(codigo, num)
				break
			} else{
				fmt.Println("Código já existente, tente novamente!")
			}
		}

		for{
			fmt.Printf("Informe o saldo da conta de código %v: ", codigo[i])
			fmt.Scanln(&saldo)
			if saldo < 0 {
				fmt.Println("Saldo inicial inválido!")
			} else{
				conta[codigo[i]] = saldo
				break
			}
		}
	}
	for{
		escolha:=0
		fmt.Println("---------MENU----------")
		fmt.Printf("1 - Efetuar Depósito\n2 - Efetuar Saque\n3 - Consultar o ativo bancário\n4 - Finalizar o programa\n")
		fmt.Scanln(&escolha)

		if escolha == 1{
			efetuarDeposito(conta, codigo)
		} else if escolha == 2{
			efetuarSaque(conta, codigo)
		} else if escolha == 3{
			somaSaldo := consultaAtivoBancario(conta)
			fmt.Printf("O ativo bancário atual é de R$ %.2f\n", somaSaldo)
		} else if escolha == 4{
			fmt.Println("Programa encerrado!")
			break
		} else{
			fmt.Println("Informe um valor válido!")
		}
	}
}

func codigoExiste(codigo int, c []int) bool{
	return slices.Contains(c, codigo)
}

func efetuarDeposito(conta map[int]float64, codigo []int){
	teste, valor := 0, 0.0
	fmt.Print("Informe o codigo da conta: ")
	fmt.Scanln(&teste)

	if !codigoExiste(teste, codigo){
		fmt.Println("Conta não encontrada!")
		return
	}
	cod := teste

	fmt.Printf("Informe o valor que deseja depositar na conta %v: ", cod)
	fmt.Scanln(&valor)

	if valor <= 0 {
		fmt.Println("Valor inválido!")
		return
	}

	conta[cod] += valor
	fmt.Println("Operação realizada com sucesso!")
}

func efetuarSaque(conta map[int]float64, codigo []int){
	teste, valor := 0, 0.0
	fmt.Print("Informe o codigo da conta: ")
	fmt.Scanln(&teste)

	if !codigoExiste(teste, codigo){
		fmt.Println("Conta não encontrada!")
		return
	}
	cod := teste

	fmt.Printf("Informe o valor que deseja sacar na conta %v: ", cod)
	fmt.Scanln(&valor)

	if valor <= 0 {
		fmt.Println("Valor inválido!")
		return
	}

	if conta[cod] < valor{
		fmt.Println("Saldo insuficiente!")
	} else{
		conta[cod] -= valor
		fmt.Println("Operação realizada com sucesso!")
	}
}

func consultaAtivoBancario(conta map[int]float64) float64{
	somaSaldo:= 0.0
	for _, s:=range conta{
		somaSaldo += s
	}
	return somaSaldo
}