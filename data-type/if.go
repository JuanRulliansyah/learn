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

	//	Short if
	if length := len(name); length > 5 {
		fmt.Println(name, "Nama terlalu panjang")
	} else {
		fmt.Println(name, "Nama lebih besar dari ", length)
	}
}
