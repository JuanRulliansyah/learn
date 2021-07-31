package main

import "fmt"

func main() {
	//Array 2D
	var mahasiswa [3][3]string
	mahasiswa[0][0] = "Ana"
	mahasiswa[1][1] = "Juan"
	mahasiswa[2][2] = "Fiqri"
	fmt.Println(mahasiswa[0][0], mahasiswa[1][1], mahasiswa[2][2])

	var mobil = [3]int{
		90,
		50,
		40,
	}
	fmt.Println(mobil[0], mobil[1], mobil[2])

}
