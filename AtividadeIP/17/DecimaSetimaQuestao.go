package main
import ("fmt")

func main(){
	var nX, nY int

	fmt.Print("Informe dois números inteiros separados por espaço: ")
	fmt.Scan(&nX, &nY)
	fmt.Print("\n")

	switch  nX % 2 {
		case 0:
			for x := 0; x < nY; x++ { 
					fmt.Printf("%v ", (nX))
					nX += 2 
			}
			fmt.Print("\n")
		default:
			fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}