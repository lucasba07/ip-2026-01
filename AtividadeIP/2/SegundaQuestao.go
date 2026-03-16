package main
import ("fmt")

func main(){
	var nj int

	fmt.Print("Informe quantos jogos serão analisados: ")
	fmt.Scan(&nj)
	fmt.Print("\n")

	fmt.Print("A seguir você digitará, consecutivamente: o número de pessoas presentes no jogo, a porcentagem que comprou ingresso na Popular, a porcentagem que comprou ingresso na Geral, a porcentagem que comprou ingresso na Arquibancada e a porcentagem que comprou ingresso nas Cadeiras\n")
	fmt.Print("\n")

	for x:=0; x < nj; x++{
		var np, pp, pg, pa, pc float32
		fmt.Print("\n")
		fmt.Printf("Em relação ao jogo %v, informe os dados ditos anteriormente separados por espaço: ", (x + 1))
		fmt.Scan(&np, &pp, &pg, &pa, &pc)
		fmt.Print("\n")

		rp := (np * (pp / 100))
		rg := (np * (pg / 100)) * 5
		ra := (np * (pa / 100)) * 10
		rc := (np * (pc / 100)) * 20
		rt := (rp + ra + rg + rc)

		fmt.Printf("A RENDA DO JOGO N. %v E = %.2f", (x + 1), rt)
		fmt.Print("\n")
	}
}