package main

import (
	"fmt"
	"strconv"
)

func main(){
	var num int

	for{
		fmt.Print("Informe um número natural qualquer: ")
		fmt.Scanln(&num)

		if num > 0 {break}
		if num <= 0 {fmt.Println("INVÁLIDO!")}
	}

	fmt.Print("Hexadecimal: ")
	decimalParaHexadecimal(num)
	fmt.Println()
}

func decimalParaHexadecimal(num int){
	var(
		hexa []int
		newHexa []string
		resto int
	)

	for num > 0{
		resto = num%16
		hexa = append(hexa, resto)
		num/= 16
	}

	for i:=(len(hexa)-1); i>=0; i--{
		switch hexa[i]{
		case 10:
			newHexa = append(newHexa, "A")
		case 11:
			newHexa = append(newHexa, "B")
		case 12:
			newHexa = append(newHexa, "C")
		case 13:
			newHexa = append(newHexa, "D")
		case 14:
			newHexa = append(newHexa, "E")
		case 15:
			newHexa = append(newHexa, "F")
		default:
			newHexa = append(newHexa, strconv.Itoa(hexa[i]))
		}
	}
	for i:= range newHexa{
		fmt.Print(newHexa[i])
	}
}
