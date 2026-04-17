package main
import "fmt"

func main() {
	a1 := make([]int, 10)
	a2 := make([]int, 10)

	for i := range a1 {
		for x := range a2 {
			if x > i{
				fmt.Printf("[%vx%v] ", (i+1), (x+1))
			} else{
				fmt.Print("  .   ")
			}
		}
		fmt.Println()
	}
}
