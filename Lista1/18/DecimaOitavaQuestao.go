package main
import ("fmt")

func main(){
	var a1, r, n, v, soma int

	fmt.Print("Informe o primeiro termo, a razão e a quantidade de termos da PA: ")
	fmt.Scan(&a1, &r, &n)

	switch {
		case n > 0:
			for x := 0; x < n; x++ {
				if x == 0 {
					v = a1
					soma += v 
				} else {
					v += r
					soma += v
				}
			}
			fmt.Printf("%v\n", soma)
		default:
			fmt.Println("VALOR INVALIDO")
	}
}