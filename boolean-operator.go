package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusUjian = ujian >= 80
	var lulusAbsensi = absensi >= 80
	var lulus = lulusUjian && lulusAbsensi

	fmt.Println(lulusUjian)
	fmt.Println(lulusAbsensi)
	fmt.Println(lulus)

	var ujian2 = 75
	var absensi2 = 80
	fmt.Println(ujian2 >= 80 && absensi2 >= 80)
}
