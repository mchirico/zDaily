package main
import (
	"fmt"
)

type A struct {
	 i int
}

func (a *A)Thing() int {
	return 2*a.i
}


func main(){
	fmt.Println("Hello Mouse")
}
