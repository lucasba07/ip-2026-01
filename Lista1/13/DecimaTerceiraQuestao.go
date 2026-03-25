package main
import ("fmt")

func main(){
	var nota float32

	fmt.Print("Informe a nota: ")
	fmt.Scan(&nota)
	fmt.Print("\n")

		switch {
			case nota >= 0 && nota < 6:
				fmt.Printf("NOTA = %.1f CONCEITO = D\n\n", nota)
			case nota >= 6 && nota < 7.5:
				fmt.Printf("NOTA = %.1f CONCEITO = C\n\n", nota)
			case nota >= 7.5 && nota < 9:
				fmt.Printf("NOTA = %.1f CONCEITO = B\n\n", nota)
		    case nota >= 9 && nota <= 10:
				fmt.Printf("NOTA = %.1f CONCEITO = A\n\n", nota)
			default:
				fmt.Print("NOTA INVALIDA")
				fmt.Print("\n")
			}
}