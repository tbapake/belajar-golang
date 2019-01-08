//menghitung luas persegi
package main

import "fmt"
var hasil, alas, tinggi int

func main() {
	fmt.Print("Masukkan nilai Alas:")
	//var alas int
	fmt.Scanln(&alas)
    fmt.Print("Masukkan nilai Tinggi:")
	//var tinggi int
	fmt.Scanln(&tinggi)
    hasil = alas * tinggi / 2
    fmt.Print("Luas segitiga adalah: ", hasil)
}