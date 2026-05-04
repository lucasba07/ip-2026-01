package main
import "fmt"

func main(){
	vetor:=make([]int, 10)
	posicao:=make([]int, 0, 10)
	valor:=make([]int, 0, 10)

	for i := range vetor{
		fmt.Print("Informe um número inteiro: ")
		fmt.Scanln(&vetor[i])

		if vetor[i] > 50 {
			posicao = append(posicao, i)
			valor = append(valor, vetor[i])
		}
	}

	if len(posicao) == 0 && len(valor) == 0{
			fmt.Println("Nenhum número digitado foi maior que 50.")
			return
		}

	fmt.Println("VALORES ACIMA DE 50:")
	fmt.Println("POSIÇÃO | VALOR ")
	for i:= range valor{
		fmt.Printf("%7d | %d\n", posicao[i], valor[i])
	}
}

