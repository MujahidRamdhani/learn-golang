package main

import (
	"fmt"
	"strings"
)
// Question 1
type FlatShape interface {
	flatBuildings() float64
}


type SolidShape interface {
	buildingSpaces() int
}

type Circle struct {
	radius float64
}

func (c Circle) flatBuildings() float64 {
	return 3.14 * c.radius * c.radius
}


type Cube struct {
	side int
}

func (c Cube) buildingSpaces() int {
	return c.side * c.side
}


func PrintFlatArea(f FlatShape) {
	fmt.Printf("Luas: %.2f\n", f.flatBuildings())
}


func PrintSolidArea(s SolidShape) {
	fmt.Printf("Luas: %d\n", s.buildingSpaces())
}

// Question 2
type Handphone struct {
	Merk       string
	Model      string
	TahunRilis int
	Warna      []string
}


type DataHP interface {
	TampilkanData() string
}



func (h Handphone) TampilkanData() string {
	return fmt.Sprintf(`
	Merk: %s
	Model: %s
	Tahun Rilis: %d
	Warna: %s`, h.Merk, h.Model, h.TahunRilis, strings.Join(h.Warna, ", "))
}



func areaRectangle(radius float64, displayWord bool) interface{} {
	if displayWord {
		return fmt.Sprintf("Luas lingkaran dengan jari-jari %.2f adalah %.2f", radius, 3.14*radius*radius)
	}
	return 3.14 * radius * radius
}




func main() {

	fmt.Println("----------------")
	fmt.Println("Question 1")
	fmt.Println("----------------")
	c := Circle{radius: 10}
	z := Cube{side: 10}
	PrintFlatArea(c)  
	PrintSolidArea(z)

	fmt.Println("----------------")
	fmt.Println("Question 2")
	fmt.Println("----------------")
	hp := Handphone{
		Merk:       "Samsung",
		Model:      "Galaxy S23",
		TahunRilis: 2023,
		Warna:      []string{"Hitam", "Biru", "Putih"},
	}
	fmt.Println(hp.TampilkanData())


	fmt.Println("----------------")
	fmt.Println("Question 3")
	fmt.Println("----------------")
	fmt.Println(areaRectangle(10, true))
	fmt.Println(areaRectangle(10, false))

	fmt.Println("----------------")
	fmt.Println("Question 4")
	fmt.Println("----------------")

	var figure1 interface{} = []int{2, 4, 6}
	var figure2 interface{} = []int{8, 10, 12}
 	// Type assertion for change  interface{} to []int
	f1 := figure1.([]int)
	f2 := figure2.([]int)

	join := append(f1, f2...)
	total := 0
	for _, num := range join {
		total += num
	}

	fmt.Println("Total:", total) // Output: 42

	fmt.Println(join)

	fmt.Println("----------------")
	fmt.Println("Question 5")
	fmt.Println("----------------")
	var figure1q5 interface{} = []int{2, 5, 7}
	var figure2q5 interface{} = []int{9, 11, 13}
 	// Type assertion for change  interface{} to []int
	f1q5 := figure1q5.([]int)
	f2q5 := figure2q5.([]int)

	joinq5 := append(f1q5, f2q5...)
	// Menghitung total semua angka
	totalq5 := 0
	for _, num := range joinq5 {
		totalq5 += num
	}

	// Menampilkan hasil dengan format yang benar
	fmt.Print("Hasil penjumlahan dari ")
	for i, num := range joinq5 {
		if i == len(joinq5)-1 {
			fmt.Printf("%d", num)
		} else {
			fmt.Printf("%d + ", num)
		}
	}
	fmt.Printf(" = %d\n", totalq5)

}
