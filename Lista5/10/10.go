package main
import "fmt"

func main(){
	fibonacci := make([]int, 50)

	for i:= range fibonacci{
		if i == 0 || i == 1{ 
			fibonacci[i] = 1 
		} else {fibonacci[i] = fibonacci[i-1] + fibonacci[i-2]}
	}

	for _, v := range fibonacci{
		fmt.Printf("%v  ", v)
	}
}
