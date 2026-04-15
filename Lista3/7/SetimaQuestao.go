package main
import(
	"fmt"
	"math"
) 
		

func main(){
	var(
		lista []float64
		num float64
		soma float64
		qntd int
		qntdImpar int
		maior float64
		menor float64
		porcen float64
		media float64
		i int
		qntdPar int
		somaPar float64
		mediaPar float64
	)

	for {
    	fmt.Printf("Informe o %vº valor: ", (i + 1))
    	fmt.Scanln(&num)

    	if num == 30000 {
        	fmt.Println("Entrada de dados finalizada!") 
        	fmt.Printf("SOMA: %.2f | QNTD: %v | MÉDIA: %.2f | MAIOR: %.2f | MÉDIA PARES: %.2f | MENOR: %.2f | %% ÍMPARES: %.2f%%\n", 
            	soma, qntd, media, maior, mediaPar, menor, porcen)
        	return
    	}

    	lista = append(lista, num)
    	soma += num
    	qntd++

    	if int(math.Round(num))%2 != 0 {
        	qntdImpar++
    	} else {
        	somaPar += num
        	qntdPar++
        	mediaPar = somaPar / float64(qntdPar)
    	}

    	if qntd == 1 { 
        	maior = num
        	menor = num
    	} else {
        	if num > maior { maior = num }
        	if num < menor { menor = num }
    	}

    	media = soma / float64(qntd)
    	porcen = (float64(qntdImpar) / float64(qntd)) * 100
    	i++
	}
}