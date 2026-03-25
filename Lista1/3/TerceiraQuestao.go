package main

import "fmt"

func main() {
	var n1, n2, n3 int

	fmt.Println("Informe o primeiro número: ")
	fmt.Scan(&n1)

	fmt.Println("Informe o segundo número: ")
	fmt.Scan(&n2)

	fmt.Println("Informe o terceiro número: ")
	fmt.Scan(&n3)

	if n1 > 9 || n1 < -9 || n2 > 9 || n2 < 0 || n3 > 9 || n3 < 0 {
		fmt.Println("DIGITO INVALIDO")
		return
	} else {
		if n1 < 0 {
			n1 = (n1 * -1)
			c := ((n1 * 100) + (n2 * 10) + n3) * -1
			fmt.Printf("%v, %v", c, (c * c))
			return
		}
			c := (n1 * 100) + (n2 * 10) + n3
			fmt.Printf("%v, %v", c, (c * c))
	}
}