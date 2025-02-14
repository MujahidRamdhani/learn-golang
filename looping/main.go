package main

import (
	"fmt"
)

func main() {

	// Buatlah sebuah program yang mencetak angka dari 1 hingga 30 dengan aturan sebagai berikut:
	// Jika angka tersebut habis dibagi 4, cetak "Coding is Fun".
	// Jika angka tersebut habis dibagi 2 tetapi tidak habis dibagi 4, cetak "Keep Learning".
	// Jika tidak memenuhi kedua kondisi di atas, cetak "Never Give Up".
	fmt.Println("----------------")
	fmt.Println("Question 1")
	fmt.Println("----------------")

	for i := 1; i <= 30; i++ {
		if i % 4 == 0 {
			fmt.Println(i, "Coding is Fun")
		}else if i % 2 == 0 {
			fmt.Println(i, "Keep Learning")
		}else {
			fmt.Println(i, "Never Give Up")
		}
	}


	// Buatlah program yang mencetak pola segitiga terbalik dengan karakter *. Program akan mencetak n baris, di mana n adalah jumlah baris yang ditentukan oleh pengguna.
	// Contoh output untuk n = 5:
	// markdown
	// Copy
	// Edit
	// *****
	// ****
	// ***
	// **
	// *
	fmt.Println("----------------")
	fmt.Println("Question 2")
	fmt.Println("----------------")
	for i := 5; i >= 1; i-- {
		for j := 1; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//mencetak hanya kata dari indeks ke 6 sampai akhir
	fmt.Println("----------------")
	fmt.Println("Question 3")
	fmt.Println("----------------")
	var sentence = [...]string{"i", "like", "to", "learn", "golang" , "because", "i", "want", "to", "become", "a", "golang", "developer"}
	// slice := array[start:end] Accesing Some Element with a Slice, start is included, end is excluded
    slice := sentence[6:]
	fmt.Println(slice)
	
	// Buatlah program yang menyimpan daftar nama buah dalam sebuah slice, lalu cetak semua nama buah beserta nomor urutnya.
	//append() digunakan untuk menambahkan elemen ke slice.
	// var buah []string
	// buah = append(buah, "Apel")
	// buah = append(buah, "Jeruk")
	fmt.Println("----------------")
	fmt.Println("Question 5")
	fmt.Println("----------------")
	var fruit []string
	fruit = append(fruit, "apple", "banana", "cherry")
	fruit = append(fruit, "grape")

	for i, name := range fruit {
		fmt.Println(i+1, name)
	}

	// Buatlah sebuah map yang menyimpan informasi panjang, lebar, dan tinggi suatu kotak dalam satuan cm. Kemudian hitung dan cetak volume kotak tersebut.

	// Contoh output:
	
	// makefile
	// Copy
	// Edit
	// panjang = 10  
	// lebar = 5  
	// tinggi = 8  
	// volume kotak = 400 cm³  
	fmt.Println("----------------")
	fmt.Println("Question 6")
	fmt.Println("----------------")

	boxes := map[string]int{
		"length": 10,
		"width": 5,
		"height": 8,
	}

	fmt.Println("Length =", boxes["length"])
	fmt.Println("Width =", boxes["width"])
	fmt.Println("Height =", boxes["height"])

	volume := boxes["length"] * boxes["width"] * boxes["height"]
	fmt.Println("Volume =", volume)
	
}