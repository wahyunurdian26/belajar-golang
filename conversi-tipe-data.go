package main

import "fmt"

// Deklarasi variabel dengan berbagai tipe data
var (
	// Konversi tipe data numerik
	nilai32 int32 = 2000
	nilai64 int64 = int64(nilai32) // Konversi dari int32 ke int64
	nilai16 int16 = int16(nilai32) // Konversi dari int32 ke int16 (potensi overflow jika nilai terlalu besar)
	nilai8  int8  = int8(nilai32)  // Konversi dari int32 ke int8 (akan terjadi overflow karena 2000 > 127)

	// Konversi string ke byte dan sebaliknya
	name      string = "dino"
	e         byte   = name[0]          // Mengambil karakter pertama ('d') dalam bentuk byte
	eString   string = string(e)        // Konversi byte menjadi string

	filem     string = "dono indro"
	a         byte   = filem[2]         // Mengambil karakter ketiga ('n') dalam bentuk byte
	eStringg  string = string(a)        // Konversi byte menjadi string
	
)

func main() {
	// Menampilkan hasil konversi numerik
	fmt.Println("Nilai 32-bit:", nilai32)
	fmt.Println("Nilai 64-bit:", nilai64)
	fmt.Println("Nilai 16-bit:", nilai16) // Bisa terjadi overflow jika nilai melebihi kapasitas int16
	fmt.Println("Nilai 8-bit:", nilai8)   // Bisa terjadi overflow jika nilai melebihi kapasitas int8

	// Menampilkan hasil konversi string ke byte lalu kembali ke string
	fmt.Println("Karakter pertama dari 'dino':", eString)   // Output: "d"
	fmt.Println("Karakter ketiga dari 'dono indro':", eStringg) // Output: "n"
}
