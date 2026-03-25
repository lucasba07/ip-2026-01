package main
import ("fmt")

func main(){
	var qnt int
	var temp float32

	fmt.Print("Quantas vezes você deseja realizar a conversão? ")
	fmt.Scan(&qnt)
	fmt.Print("\n")

	for x := 1; x <= qnt; x++ {
		fmt.Printf("%vº conversão...\n", x)
		fmt.Print("Informe a temperatura em Fahrenheit: ")
		fmt.Scan(&temp)
		fmt.Print("\n")

		celsius := (5 * (temp - 32)) / 9

		fmt.Printf("%v° Fahrenheit EQUIVALE A %.2f° CELSIUS\n\n", temp, celsius)
	}
}