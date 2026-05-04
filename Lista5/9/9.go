package main
import "fmt"

func main(){
	alturas := make([]float64, 10)
	soma, media := 0.0, 0.0

	for i:= range alturas{
		fmt.Printf("Informe a altura do %vº atleta: ", (i+1))
		fmt.Scanln(&alturas[i])
		soma+=alturas[i]
	}
	media = (soma / 10)

	for i, v := range alturas{
		if v > media{
			fmt.Printf("A altura do %vº atleta, %.2fm, é maior que a média das alturas\n", (i+1), v)
		}
	}
}
