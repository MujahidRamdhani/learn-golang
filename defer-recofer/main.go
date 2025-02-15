package main

import (
	"fmt"
	"sort"
	"sync"
)
func helloWorldDefer() {
	defer fmt.Println("world")
	fmt.Println("hello")
}

func handlepanic() {
	if r := recover(); r != nil {
		fmt.Println("Recovered from panic:", r)
	}
}

func calculate(a, b int) interface{} {
	defer handlepanic()
	 if(b != 0) {
		return a / b
	} else {
		panic("Cannot divide by zero")
	}
}

func addFigures(figures *int) {
	*figures += 10
}

func sendFruits(fruits *[]string, f ...string) {
	*fruits = append(*fruits, f...)
	sort.Strings(*fruits)
	fmt.Println("Sorted fruits:")
	for i, fruit := range *fruits {
		fmt.Printf("Fruit %d: %s\n", i+1, fruit)
	}
}

func printNumber(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(num)
}

func handleWg() {
	var wg sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg.Add(1)
		go printNumber(i, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func sendCities(ch chan string) {
    cities := []string{"Jakarta", "Bandung", "Surabaya", "Yogyakarta", "Medan"}
    for _, city := range cities {
        ch <- city
    }
    close(ch)
}


func main() {
	fmt.Println("----------------")
	fmt.Println("Question 1")
	fmt.Println("----------------")
	helloWorldDefer()

	fmt.Println("----------------")
	fmt.Println("Question 2")
	fmt.Println("----------------")
	a := 10
	b := 2
	result := calculate(a, b)
	fmt.Println("Result:", result)

	fmt.Println("----------------")
	fmt.Println("Question 3")
	fmt.Println("----------------")
	figures := 5
	addFigures(&figures)
	fmt.Println("Figures:", figures)

	fmt.Println("----------------")
	fmt.Println("Question 4")
	fmt.Println("----------------")
	fruits := []string{"apple", "banana", "orange"}
	sendFruits(&fruits, "mango", "grapes")

	fmt.Println("----------------")
	fmt.Println("Question 5")
	fmt.Println("----------------")
	handleWg()
	
	fmt.Println("----------------")
	fmt.Println("Question 6")
	fmt.Println("----------------")
    ch := make(chan string)
	go sendCities(ch)

	for city := range ch {
		fmt.Println(city)
	}

}
