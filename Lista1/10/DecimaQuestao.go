package main
import ("fmt")

func main (){
	var aM, bM, cM, dM float32

	fmt.Print("Informe o valor do elemento 'a' da matriz: ")
	fmt.Scan(&aM)

	fmt.Print("Informe o valor do elemento 'b' da matriz: ")
	fmt.Scan(&bM)

	fmt.Print("Informe o valor do elemento 'c' da matriz: ")
	fmt.Scan(&cM)

	fmt.Print("Informe o valor do elemento 'd' da matriz: ")
	fmt.Scan(&dM)

	det := (aM * dM) - (bM * cM)

	fmt.Printf("O VALOR DO DETERMINANTE E = %.2f\n\n", det)
}