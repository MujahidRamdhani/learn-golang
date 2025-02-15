package main
import "fmt"

func calculateLengthExtension(length *float64, width *float64, broad *float64){
	*broad= (*length * *width)
}

func calculatePerimeterOfRectangle(length *float64, width *float64, around *float64){
	*around = (2 * (*length + *width))
}

func makeIntroductions(name *string, age *int) {
	fmt.Println("Hello, my name is", *name, "and I am", *age, "years old.")
}

func addAnimalsData(animal *[]string) {
    *animal = append(*animal, "cat", "dog", "hourse")
}

func addBooksData(title string, author string, year string, booksData *[]map[string]string) {
    book := map[string]string{
        "title" : title,
        "author" : author,
        "year" : year,
    }

    *booksData = append(*booksData, book)
}

func main() {
	fmt.Println("----------------")
	fmt.Println("Question 1")
	fmt.Println("----------------")
	l, w, b := 5.0, 3.0, 0.0
	calculateLengthExtension(&l, &w, &b)
	fmt.Println("Broad:", b)

	fmt.Println("----------------")
	l, w, a := 5.0, 3.0, 0.0
	calculatePerimeterOfRectangle(&l, &w, &a)
	fmt.Println("Around:", a)

	fmt.Println("----------------")
	fmt.Println("Question 2")
	fmt.Println("----------------")
	name, age := "John", 25
	makeIntroductions(&name, &age)

	fmt.Println("----------------")
	fmt.Println("Question 3")
	fmt.Println("----------------")
	animals := []string{"elephant", "tiger"}
	fmt.Println("before adding animals data: ", animals)
	addAnimalsData(&animals)
	fmt.Println("after adding animals data: ", animals)


	fmt.Println("----------------")
	fmt.Println("Question 4")
	fmt.Println("----------------")

	var books []map[string]string

	addBooksData("The Lord of the Rings", "J.R.R Tolkien", "2012", &books)
	addBooksData("Harry Potter", "J.K. Rowling", "1997", &books)
	
	// for i, book := range books {
	// 	fmt.Printf("%d. Title: %s, Author: %s, Year: %s\n", i+1, book["title"], book["author"], book["year"])
	// }
}