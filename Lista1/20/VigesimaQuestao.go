package main
import ("fmt")

func main(){
	var hrs, min, seg int

	fmt.Print("Informe as horas: ")
	fmt.Scan(&hrs)

	fmt.Print("Informe os minutos: ")
	fmt.Scan(&min)

	fmt.Print("Informe os segundos: ")
	fmt.Scan(&seg)

	tempo := (hrs * 3600) + (min * 60) + seg

	fmt.Printf("O TEMPO EM SEGUNDOS E = %v\n", tempo)
}