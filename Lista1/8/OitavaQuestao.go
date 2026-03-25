package main
import ("fmt")

func main(){
	var raio, altura float32

	fmt.Print("Informe o raio da lata em metros: ")
	fmt.Scan(&raio)

	fmt.Print("Informe a altura da lata em metros: ")
	fmt.Scan(&altura)

	areaBase := (3.14159 * (raio * raio)) * 2
	areaLateral := (2 * 3.14159 * raio) * altura
	custoLata := (areaBase + areaLateral) * 100

	fmt.Printf("O VALOR DO CUSTO E = %.2f\n\n", custoLata)
}