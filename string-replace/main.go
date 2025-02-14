package main

import (
    "fmt"
	"strings"
)

func main() {
	//learn 1
	nama := "ahmad"
	umur := 23
	kota := "bandung"
	fmt.Printf("%s %d %s\n", nama, umur, kota)

	//learn 2
	halo := "hello Word"
	halo = strings.Replace(halo, "Word", "GOLANG", 1)
	fmt.Println(halo)

	//learn 3
	var firstWod = "i"
	var secondWord = "am"
	var thirdWord = "learning"
	var fourthWord = "golang"
	firstWod = strings.ToUpper(firstWod)
	secondWord = strings.Replace(secondWord, "a", "A", 1)
	thirdWord= strings.Replace(thirdWord, "l", "l", 1)
	fourthWord = strings.Replace(fourthWord, "g", "g", 1)

	fmt.Printf("%s %s %s %s\n", firstWod, secondWord, thirdWord, fourthWord)
}
