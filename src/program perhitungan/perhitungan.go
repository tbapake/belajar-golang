package main

import "fmt"

func main() {

	var pilihan int

	fmt.Println("1. Luas Persegi")
	fmt.Println("2. Luas Segitiga")
	fmt.Println("3. Luas Lingkaran")

	fmt.Print("Masukkan pilihan anda: ")
	fmt.Scanln(&pilihan)
	
	switch pilihan{
    case 1 : luas_persegi()
    case 2 : luas_segitiga()
    case 3 : luas_lingkaran()
    default : main()

	}
}


func luas_persegi() {
	var (sisi, Hasil int)
	fmt.Println("                                      ")
	fmt.Println("======================================")
	fmt.Println("======Perhitungan Luas Persegi========")
	fmt.Println("======================================")
	fmt.Println("                                      ")
	fmt.Print("Masukkan panjang sisi persegi: ")
	fmt.Scanln(&sisi)
	Hasil = sisi * sisi
	fmt.Println("                                      ")
	fmt.Println("Luas persegi adalah: ", Hasil)
	fmt.Println("                                      ")
	fmt.Println("======================================")
}

func luas_segitiga() {
	var alas, tinggi, hasil int

	fmt.Print("Masukkan nilai Alas:")
	fmt.Scanln(&alas)
    fmt.Print("Masukkan nilai Tinggi:")
	fmt.Scanln(&tinggi)
    hasil = alas * tinggi / 2
    fmt.Print("Luas segitiga adalah: ", hasil)
}

func luas_lingkaran() {
	var jari2, hasil int
	fmt.Print("Masukkan nilai jari-jari: ")
	hasil = (22/7) * jari2 * jari2
	fmt.Println("Luas lingkaran adalah: ", hasil)
}