package main
import "fmt"

func main(){
	var soma int
	for i:= 1; i <= 20; i++{
		fmt.Printf("%v ", i)
		soma+=i
	}
	fmt.Printf("\nA soma resulta em %v\n", soma)
}
