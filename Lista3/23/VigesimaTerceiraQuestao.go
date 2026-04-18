package main
import "fmt"

func main(){
	var n int

	for{
		fmt.Print("Informe um número natural qualquer: ")
		fmt.Scanln(&n)

		if n >= 1{break}
		fmt.Println("VALOR INVÁLIDO!")
	}

	cont:=1000.0
	soma:=0.0
	for i:=1; i<=n; i++{
		if i%2==0{soma-=cont/float64(i)}
		if i%2!=0{soma+=cont/float64(i)}
		cont-=3
	}
	fmt.Printf("%.2f\n", soma)
}
