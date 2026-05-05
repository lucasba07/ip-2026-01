package main
import "fmt"

func main(){
	vetor:= make([]int, 10)
	vetorPrimos:=make([]int, 0, 10)
	vetorPosicoes:=make([]int, 0, 10)

	for i:= range vetor{
		fmt.Print("Informe um valor inteiro: ")
		fmt.Scanln(&vetor[i])

		if isPrimo(vetor[i]){
			vetorPrimos = append(vetorPrimos, vetor[i])
			vetorPosicoes = append(vetorPosicoes, i)
		}
	}
	if len(vetorPrimos) == 0{
		fmt.Println("Não foram digitados números primos!")
	} else{
		fmt.Println("PRIMOS | POSIÇÃO")
		for i:= range vetorPrimos{
			fmt.Printf("%6v | %v\n", vetorPrimos[i], vetorPosicoes[i])
		}
	}
}

func isPrimo(a int)bool{
	if a <= 1 { 
		return false 
	}else if a == 2{ 
		return true 
	}else{
		for i:=2; i < a; i++{
			if a % i == 0{
				return false
			}
		}
	}
	return true
}