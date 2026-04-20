package main

import "fmt"

func main() {
	var (
		id, idBoiGordo, idBoiMagro int
		peso, boiGordo, boiMagro   float64
	)

	for i := 1; i <= 90; i++ {
		fmt.Printf("Boi %d\n", i)

		fmt.Print("ID: ")
		fmt.Scanln(&id)

		fmt.Print("Peso(kg): ")
		fmt.Scanln(&peso)

		if i == 1 || peso > boiGordo {
			boiGordo = peso
			idBoiGordo = id
		}

		if i == 1 || peso < boiMagro {
			boiMagro = peso
			idBoiMagro = id
		}
	}

	fmt.Println("RESULTADO:")
	fmt.Printf("Boi mais gordo → ID: %d | Peso: %.2f kg\n", idBoiGordo, boiGordo)
	fmt.Printf("Boi mais magro → ID: %d | Peso: %.2f kg\n", idBoiMagro, boiMagro)
}