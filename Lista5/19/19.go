package main
import "fmt"

func main(){
	num := make([]int, 10)
	div := make([]int, 5)

	for i:=range num{
		fmt.Print("Informe um número inteiro: ")
		fmt.Scanln(&num[i])

		if i < 5{
			fmt.Print("Informe outro número inteiro: ")
			fmt.Scanln(&div[i])
		}
	}
	for _, n := range num {
		fmt.Printf("Número %v:\n", n)

		if isDivisivel(n, div) == 0 {
			fmt.Println("Não é divisível por nenhum dos valores!")
		}
		fmt.Println() 
	}
}

func isDivisivel(n int, div []int) int {
	cont := 0
	for i, v := range div {
		if v != 0 && n%v == 0 {
			fmt.Printf("Divisível por %v na posição %v\n", v, i)
			cont++
		}
	}
	return cont
}