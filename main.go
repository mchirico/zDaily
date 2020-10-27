package main

import (
	"fmt"
)

type Cat struct {
	Name string
	Breed string
}

func newKitten(Name, Breed string) *Cat {

	m:=&Cat{Name: Name,Breed: Breed}

	return m
}

func (cat *Cat) Print() {
	fmt.Println(cat)
}


func main(){
	 cat := newKitten("Mousey","tuxedo")
	 cat.Print()

}