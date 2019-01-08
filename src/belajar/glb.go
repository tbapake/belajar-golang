// Menghitung gerak lurus beraturan oleh : Hendra Permana yang sudah di edit
package main

import "fmt"

var	s, v, t float64

func main() {
	

	fmt.Println("Masukan nilai jarak (dalam satuan Meter) :")
	fmt.Scanln(&s)
	fmt.Println("Masukan nilai waktu tempuh (dalam satuan detik) :")
	fmt.Scanln(&t)
	v = s / t
	fmt.Println("Hasilnya adalah ",v, "m/s")
}
