package main
import ("fmt")

func main(){
	var conta int
	var consumo, valorConta float32
	var tipo string

	fmt.Print("Informe o número da sua conta: ")
	fmt.Scan(&conta)
	fmt.Print("\n")

	fmt.Print("Informe seu consumo de água por m³: ")
	fmt.Scan(&consumo)
	fmt.Print("\n")

	fmt.Print("Informe o seu tipo de consumo: C - COMERCIAL | I - INDUSTRIAL | R - RESIDENCIAL | ")
	fmt.Scan(&tipo)
	fmt.Print("\n")

	switch tipo {
		case "C", "c":
			if consumo > 80 {
				valorConta = 500 + ((consumo - 80) * 0.25)
			} else {
				valorConta = 500
			}
		case "I", "i":
			if consumo > 100 {
				valorConta = 800 + ((consumo - 100) * 0.04)
			} else {
				valorConta = 800
			}
		case "R", "r":
			valorConta = 5 + (0.05 * consumo)
		default:
			fmt.Print("VALOR DO TIPO DE CONSUMO INVALIDO!")
	}
	fmt.Printf("CONTA = %v\n", conta)
	fmt.Printf("VALOR DA CONTA = %.2f", valorConta)
}
