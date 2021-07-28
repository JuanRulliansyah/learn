package main

import "fmt"

func main() {
	var name string
	name = "Juan Rulliansyah"
	fmt.Println(name)

	var fullName = "Muhammad Juan Akmal Rulliansyah"
	fmt.Println(fullName)

	var age int8 = 21
	fmt.Println(age)

	//Variable Declaration without using Var
	address := "Sawangan"
	fmt.Println(address)

	friendName := "Eko"
	fmt.Println(friendName)

	friendName = "Bima"
	fmt.Println(friendName)

	//Multiple variable declaration (Not list or array)
	var (
		merk  = "Toyota"
		badge = "Innova"
	)

	fmt.Println(merk, badge)

	var (
		firstName = "Juan"
		lastName  = "Rulliansyah"
	)

	fmt.Println(firstName, lastName)

}
