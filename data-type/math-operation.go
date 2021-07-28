package main

import "fmt"

func main() {
	var result = 10 + 10
	fmt.Println(result)

	//Augmented Assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)

	//Unary Operator
	i++
	fmt.Println(i) //21
	i--
	fmt.Println(i) //20

}
