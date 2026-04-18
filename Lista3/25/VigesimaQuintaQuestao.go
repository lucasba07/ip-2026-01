package main
import "fmt"

func main(){
	cont:=1.0
	soma:=0.0

	for i:= 15; i>=1; i--{
		if i%2==0{soma-=(cont/float64((i*i)))}
		if i%2!=0{soma+=(cont/float64((i*i)))}
		cont*=2
	}
	fmt.Printf("%.2f\n", soma)
}
