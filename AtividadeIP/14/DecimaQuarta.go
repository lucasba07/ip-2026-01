package main
import ("fmt"; "math")

func main(){
	var alturaP, arestaP float64

	fmt.Print("Informe, em metros, a altura da pirâmide: ")
	fmt.Scan(&alturaP)
	fmt.Print("\n")

	fmt.Print("Informe, em metros, a aresta da pirâmide: ")
	fmt.Scan(&arestaP)
	fmt.Print("\n")

	areaBaseP := (3 * (arestaP * arestaP) * (math.Sqrt(3))) / 2
	volume := (areaBaseP / 3) * alturaP

	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS\n\n", volume)
}