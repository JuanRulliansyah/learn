package main

import "fmt"

func main() {
	var months = [...]string{
		"Januari",
		"Februari",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"Oktober",
		"November",
		"Desember",
	}

	var list [10]string
	list[0] = "Juan"
	fmt.Println(list)

	//var slice1 = months[4:7]
	//fmt.Println(slice1)
	//fmt.Println(len(slice1))
	//fmt.Println(cap(slice1))
	//
	//months[5] = "Bulan Bintang"
	//fmt.Println(slice1)

	//slice1[0] = "Mei Lagi"
	//fmt.Println(slice1)
	//
	//tambah := append(slice1, "Bulan 13")
	//fmt.Println(tambah)

	var slice2 = months[2:4]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Eko")
	fmt.Println(slice3)
	slice3[1] = "Bukan Desember"
	fmt.Println(slice3)

	fmt.Println(slice2)
	fmt.Println(months)

	// Creating Slice
	newSlice := make([]string, 2, 5)
	newSlice[0] = "Juan"
	newSlice[1] = "Rulliansyah"
	fmt.Println(newSlice)
	fmt.Println(len(newSlice))
	fmt.Println(cap(newSlice))

	// Copy Slice
	copySlice := make([]string, len(newSlice), cap(newSlice))
	copy(copySlice, newSlice)
	fmt.Println(copySlice)

	//	Perbedaan Slice & Array
	array := [...]string{"Muhammad", "Juan", "Akmal", "Rulliansyah"}
	slice := []string{"Muhammad", "Juan", "Akmal", "Rulliansyah"}

	fmt.Println(array, slice)

}
