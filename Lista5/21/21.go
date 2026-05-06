package main
import "fmt"

func main(){
	vetor := make([]float64, 10)
	
	for i:=range vetor{
		fmt.Print("Informe um número qualquer: ")
		fmt.Scanln(&vetor[i])
	}

	num:= 3
	for{
		fmt.Println("Como você deseja ver o vetor digitado?")
		fmt.Print("0 - NÃO DESEJO VER | 1 - FORMA DIRETA | 2 - FORMA INVERSA ")
		fmt.Scanln(&num)

		switch num{
			case 0:
				return
			case 1:
				fmt.Println(vetor)
			case 2:
				vetorInverso := make([]float64, 10)
				cont:=9
				for i:=range vetorInverso{
					vetorInverso[i] = vetor[cont]
					cont--
				}

				fmt.Println(vetorInverso)
			default:
				fmt.Println("DIGITE UM VALOR VÁLIDO!")
		}
	}
}