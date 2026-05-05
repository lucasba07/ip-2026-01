package main
import "fmt" 

func main(){
	funcionarios := make(map[int]int)

	for i:=1; i<=100; i++{
		meses, escolha:= 0,0

		for{
			fmt.Printf("Informe a quantidade de meses de trabalho do funcionário %d: ", i)
			fmt.Scanln(&meses)
			if mesIsValido(meses, funcionarios){
				funcionarios[meses] = i
				break
			}
		}
		
		if i % 10 == 0{
			fmt.Println("VOCÊ DESEJA CONTINUAR? 1 - SIM | 2 NÃO")
			fmt.Scanln(&escolha)
			if escolha == 2 { break }
		}
	}
	empregadosRecentes(funcionarios)
}

func empregadosRecentes(f map[int]int) {
    m1, m2, m3 := 99999, 99999, 99999
    
    for m := range f {
        if m < m1 {
            m3 = m2
            m2 = m1
            m1 = m
        } else if m < m2 {
            m3 = m2
            m2 = m
        } else if m < m3 {
            m3 = m
        }
    }
    
    fmt.Println("TRÊS EMPREGADOS COM MENOS MESES DE TRABALHO:")
    for m, n := range f {
        if m == m1 || m == m2 || m == m3 {
            fmt.Printf("Número: %d | Meses de trabalho: %d\n", n, m)
        }
    }
}

func mesIsValido(m int, l map[int]int)bool{
	for x := range l{
		if x == m{
			return false
		}
	}
	return true
}