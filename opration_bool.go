package main

import "fmt"

func main() {
	var nilaiAkhir = 90
	var nilaiAbsen = 81

	var lulusNilaiAkhir = nilaiAkhir > 80
	var lulusNilaiAbsen = nilaiAbsen > 80

	fmt.Println("Selamat kamu lulus : ", lulusNilaiAkhir && lulusNilaiAbsen)
}
