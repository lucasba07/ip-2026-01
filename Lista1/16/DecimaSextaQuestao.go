package main
import ("fmt")

func main(){
	var salarioF float32

	fmt.Print("Informe o salário do funcionário: ")
	fmt.Scan(&salarioF)
	fmt.Print("\n")

	switch {
		case salarioF <= 300:
			reajuste := (salarioF * 1.5)
			fmt.Printf("SALARIO COM REAJUSTE = %.2f\n\n", reajuste)
	    case salarioF > 300:
			reajuste := (salarioF * 1.3)
			fmt.Printf("SALARIO COM REAJUSTE = %.2f\n\n", reajuste)
		default:
			fmt.Println("VALOR INVALIDO")	
	}
}