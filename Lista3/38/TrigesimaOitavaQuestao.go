package main
import "fmt"

func main(){
	cpf := make([]int, 11)
	var digito int

	for i:= range cpf{
		fmt.Printf("Informe o %vº dígito do cpf: ", (i+1))
		fmt.Scanln(&digito)
		cpf[i] = digito
	}
	if cpfValido(cpf){
		fmt.Println("O CPF é válido!")
	} else{
		fmt.Println("O CPF não é válido!")
	}
}

func cpfValido(cpf []int) bool{
	soma10:=0
	soma11:=0

	verificaDigito10:=0
	verificaDigito11:=0

	cont1:= 10
	for i:=0; i<9; i++{
		soma10 += (cpf[i] * cont1)
		cont1--
	}
	
	if soma10 % 11 >= 2 {
		verificaDigito10 = 11-(soma10%11) 
	}

	cont2:= 11
	for i:=0; i<10; i++{
		soma11 += (cpf[i] * cont2)
		cont2--
	}

	if soma11 % 11 >= 2 {
		verificaDigito11 = 11-(soma11%11) 
	}

	if verificaDigito10 == cpf[9] && verificaDigito11 == cpf[10]{
		return true
	}
	return false
}
