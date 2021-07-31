package main

import "fmt"

func main() {
	data := make(map[string]string)
	data["nama"] = "Juan"
	data["umur"] = "21"

	switch data["nama"] {
	case "Juan":
		fmt.Println("Hello Juan")
	case "Ana":
		fmt.Println("Hello Ana")
	default:
		fmt.Println("Kamu siapaa? Kamu siapaa?")
	}

	switch length := 5; len(data["nama"]) > length {
	case true:
		fmt.Println(data["nama"], " lebih banyak daripada ", length)
	case false:
		fmt.Println(data["nama"], " lebih sedikit atau sama dengan daripada ", length)
	}

	length := len(data["nama"])
	switch {
	case length > 10:
		fmt.Println("Nama lebih besar dari 5 karakter")
	case length < 5:
		fmt.Println("Nama lebih kecil atau sama dengan dari 5 karakter")
	default:
		fmt.Println("Nama sudah benar")
	}

}
