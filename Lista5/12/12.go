package main
import "fmt"

func main(){
	notas := make(map[int]int)

	for i:=0; i<15; i++{
		num := 0
		fmt.Printf("Informe a %vº nota: ", (i+1))
		fmt.Scanln(&num)
		
		notas[num]++
	}
	frequencias(notas)
}

func frequencias(a map[int]int){
	fmt.Println("NOTA | FA | FR")
	for v, i:= range a{
		fmt.Printf("%4d | %d  | %.2f\n", v, i, (float64(i)/15.0))
	}
}
