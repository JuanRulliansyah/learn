package main

import "fmt"

func main() {
	//data := map[string]string{
	//	"nama": "Muhammad Juan Akmal Rulliansyah",
	//	"jurusan": "Teknik Informatika",
	//}

	//data := make(map[string]string)
	var data map[string]string = make(map[string]string)
	data["nama"] = "Muhammad Juan Akmal Rulliansyah"
	data["jurusan"] = "Teknik Informatika"

	fmt.Println(data, data["nama"], data["jurusan"])

	data["universitas"] = "Universitas Pamulang"
	fmt.Println(data["universitas"])

	data["ups"] = "Kesalahan"

	fmt.Println(data)
	delete(data, "ups")
	fmt.Println(data, len(data))
}
