package main
import "fmt"

func main(){
	idades:= make(map[int]int)

	for i:=1; i<=50; i++{
		idade := 0
		fmt.Printf("Informe a %vº idade: ", i)
		fmt.Scanln(&idade)

		idades[idade]++
	}

	maisRep:= 0
	for _, rep:=range idades{
		if rep > maisRep{
			maisRep = rep
		}
	}

	var modas []int
	for idade, rep:=range idades{
		if rep == maisRep{
			modas = append(modas, idade)
		}
	}

	fmt.Printf("A maior frequência das idades digitadas é %v\n", maisRep)
	fmt.Printf("Moda(s): %v\n", modas)
}