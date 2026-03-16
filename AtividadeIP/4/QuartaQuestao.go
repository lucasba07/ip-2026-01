package main
import ("fmt")

func main(){
	var salario, gasto_kW float32

	fmt.Print("Qual é o valor do salário mínimo em reais? ")
	fmt.Scan(&salario)

	fmt.Print("Quanto sua residência gasta de energia em kW? ")
	fmt.Scan(&gasto_kW)

	valor_kW := (salario * 0.7) / 100
	custoTotal := valor_kW * gasto_kW
	custoDesconto := custoTotal * 0.9

	fmt.Printf("Custo por kW: R$ %.2f\n", valor_kW)
	fmt.Printf("Custo do consumo: R$ %.2f\n", custoTotal)
	fmt.Printf("Custo com desconto: R$ %.2f", custoDesconto)
}