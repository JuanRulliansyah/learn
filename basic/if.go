package main

import "fmt"

func main() {
	name := "Juan"

	if name == "Juan" {
		fmt.Println("Halo Juan")
	} else if name == "Joko" {
		fmt.Println("Halo Joko")
	} else {
		fmt.Println("Wah, salah")
	}

	slice := []string{
		"Pablo",
		"Escobar",
	}

	slice2 := make([]string, 3, 5)
	slice2[0] = "Pacho"
	slice2[1] = "Herrera"

	slice3 := make([]string, len(slice2), cap(slice2))
	copy(slice3, slice2)

	fmt.Println(slice, slice2)
	fmt.Println(slice3)

	if length := 3; len(slice[0]) > length {
		fmt.Println(slice[0], " is greater than ", length)
	} else {
		fmt.Println(slice[0], " is equal or not greater than ", length)
	}

}
