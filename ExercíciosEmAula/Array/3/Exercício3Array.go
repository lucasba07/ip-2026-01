package main
import "fmt"

func main() {
	var num [10]float64

	for i:= range num{
		fmt.Print("Informe um valor qualquer: ")
		fmt.Scan(&num[i])
	}

	for i:= range num{
		fmt.Printf("%.2f\n", num[9-i])
	}
}