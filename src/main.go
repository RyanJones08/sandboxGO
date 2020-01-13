package main

import (
	"fmt"
)

func main(){
	var name string
	fmt.Println("Enter Name")
	fmt.Scanln(&name)	

	fmt.Println("Hello", name)
}