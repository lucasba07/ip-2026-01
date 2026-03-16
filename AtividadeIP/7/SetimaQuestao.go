package main
import ("fmt")

func main(){
	var far, chuvaPole float32

	fmt.Print("Informe o valor da temperatura em Fahrenheit: ")
	fmt.Scan(&far)
	fmt.Print("\n")

	fmt.Print("Informe a quantidade de chuva em Polegadas: ")
	fmt.Scan(&chuvaPole)
	fmt.Print("\n")	

	cel := ((5 * far) - 160) / 9
	chuvaMili := chuvaPole * 25.4

	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", cel)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", chuvaMili)
}