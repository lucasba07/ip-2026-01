package main
import ("fmt")

func main(){
	var valorInt int
	var soma float64 

	fmt.Print("Informe um número inteiro maior que 1: ")
	fmt.Scan(&valorInt)

	switch {
		case valorInt > 1:
			for x := 1; x <= valorInt; x++ {
				valor := (1 / float64(x))
				soma += valor
			}
			fmt.Printf("%.6f\n", soma)
		default:
			fmt.Println("Numero invalido!")
	}
}