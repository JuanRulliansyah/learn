package main

import "fmt"

func main() {
	type NoKTP string

	var ktpJuan NoKTP = "181011400479"
	fmt.Println(ktpJuan)
	fmt.Println(NoKTP("181011400515"))

	type toyota string
	var mobil toyota = "innova"
	var mobil2 toyota = "kijang"

	fmt.Println(mobil, mobil2)
}
