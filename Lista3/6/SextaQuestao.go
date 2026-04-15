package main
import "fmt"

func main (){
	var num int
	
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scanln(&num)
	
	if num > 0{
		isTriangular(num)
	} else{
		fmt.Println("Número inválido!")
	}
}

func isTriangular(num int) {
	for i := 0; i <= num; i++{
		if (i * (i+1) * (i+2)) == num{
			fmt.Printf("O número %v é triangular\n", num)
			return
		} 
	}
	fmt.Printf("O número %v não é triangular\n", num)
}