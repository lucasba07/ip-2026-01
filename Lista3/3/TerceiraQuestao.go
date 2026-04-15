package main
import "fmt" 

func main (){
	var (
		salCarlos float64
		salJoao float64
	)
	fmt.Print("Bom dia Carlos! Qual o seu salário? ")
	fmt.Scanln(&salCarlos)

	salJoao = salCarlos / 3
	mes:= 1

	for{
		fmt.Printf("%vº mês...\n", mes)
		salCarlos *= 1.02
		salJoao *= 1.05
		if (salJoao >= salCarlos){
			fmt.Printf("O salário de João igualou/utrapassou o de Carlos após %v meses\n", mes)
			return
		}
		mes++
	}
}
