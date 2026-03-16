package main
import ("fmt")

func main(){
	var n int

	fmt.Print("Informe um número inteiro maior que 5 e menor que 2000: ")
	fmt.Scan(&n)
	fmt.Print("\n")

	switch {
		case n > 5 && n < 2000:
			for x := 2; x <= n; x += 2 {
				quad := (x * x)
				fmt.Printf("%v ^ 2 = %v\n", x, quad)
			}
		default:
			fmt.Println("VALOR INVALIDO")
	}
}