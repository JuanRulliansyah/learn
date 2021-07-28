package main

import "fmt"

func main() {
	var (
		name1 = "Juan"
		name2 = "Juan"
	)

	var result bool = name1 == name2
	fmt.Println(result)

	var (
		value1 = 100
		value2 = 200
	)

	fmt.Println(value1 == value2)
	fmt.Println(value1 < value2)

	var nilaiAbsensi = 60
	var nilaiAkhir = 70

	var lulusNilaiAbsensi = nilaiAbsensi > 80
	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulus = lulusNilaiAkhir && lulusNilaiAbsensi

	fmt.Println(lulus)
}
