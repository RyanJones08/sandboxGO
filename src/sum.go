package main

import (
	"fmt"
)

func main(){
	var a int
	fmt.Println("Enter Number")
	fmt.Scanln(&a)	

	var b int
	fmt.Println("Enter Another Number")
	fmt.Scanln(&b)

	result := sum(a, b)
	fmt.Println("Sum", result)
}

func sum(x int, y int) int{
	return x + y
}