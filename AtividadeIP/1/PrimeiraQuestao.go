package main
import ("fmt")

func main(){
	var n1, n2, n3, media float32

	fmt.Print("Informe as três notas do estudante: ")
	fmt.Scan(&n1, &n2, &n3)

	media = (n1 + n2 + n3) / 3

	fmt.Printf("MEDIA = %.2f\n", media)
	if media >= 6 {
		fmt.Print("APROVADO\n")
	} else {
		fmt.Print("REPROVADO\n")
	}
}