package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	jogadas := make([]int, 0, 20)
	dados := make(map[int]int)
	for i:=0; i<20; i++{
		n:=rand.Intn(6) + 1
		jogadas = append(jogadas, n)
		dados[n]++
	}

	fmt.Printf("Jogadas: %v\n", jogadas)

	fmt.Println("NÚMERO SORTEADO | FREQUÊNCIA")
	for i := 1; i <= 6; i++ {
		fmt.Printf("%15d | %d\n", i, dados[i])
	}
}