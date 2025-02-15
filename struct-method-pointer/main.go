package main
import (
	"fmt"
)

type fruit struct {
	name string
	color string
	withSeeds bool
	price int
}

func (f *fruit) changeValue(name string, color string, withSeeds bool, price int) {
	f.name = name
	f.color = color
	f.withSeeds = withSeeds
	f.price = price
}

type triangle struct {
	pedestal float64
	height float64
}

type square struct {
	side float64
}

type rectangle struct {
	length float64
	width float64
}

type handphone struct {
	name string
	brand string
	year int
	availableColors []string
}

type film[] struct {
	
}


func (t *triangle) calculateAreaTriangle() float64 {
	return 0.5 * t.pedestal * t.height
}

func (s *square) calculateAreaSquare() float64 {
	return s.side  * s.side
}

func (r *rectangle) calculateAreaRectangle() float64 {
	return r.length * r.width
}

func (h *handphone) addColor(color string) {
	h.availableColors = append(h.availableColors, color)
}
//Using variadic parameters to accept more than one color
func (h *handphone) addArrayColor(color ...string) {
	h.availableColors = append(h.availableColors, color...)
}

type Movie struct {
	title string
	genre string
	duration int
	year int
}

func addMovieData(title string, genre string, duration int, year int, m *[]Movie) {
	movie := Movie{title: title, genre: genre, duration: duration, year: year}
	*m = append(*m, movie)
}

type student struct {
	name string
	major string
	year int
	gpa float64
	hobbies []string
}

// func addStudentData(name string, major string, year int, gpa float64, hobbies []string, s *[]student) {
// 	student := student{name: name, major: major, year: year, gpa: gpa, hobbies: hobbies}
// 	*s = append(*s, student)
// }

func addStudentData(name string, major string, year int, gpa float64, s *[]student, hobbies ...string) {
	student := student{name: name, major: major, year: year, gpa: gpa, hobbies: hobbies}
	*s = append(*s, student)
}

func (m Movie) String() string {
	return fmt.Sprintf("Title : %s | Genre: %s | Duration: %d min | Year: %d", m.title, m.genre, m.duration, m.year)
}

func (s student) String() string {
	return fmt.Sprintf("Name : %s | Major: %s | Year: %d | GPA: %f | Hobbies: %v", s.name, s.major, s.year, s.gpa, s.hobbies)
}

func main() {
	fmt.Println("----------------")
	fmt.Println("Question 1")
	fmt.Println("----------------")
	fruits := fruit{}
	fruits.changeValue("apple", "red", true, 10000)
	fmt.Println(fruits)
	fruits.changeValue("banana", "yellow", false, 5000)
	fmt.Println(fruits)

	fmt.Println("----------------")
	fmt.Println("Question 2")
	fmt.Println("----------------")
	t:= triangle{pedestal: 10, height: 5}
	fmt.Println("triangle value :",t.calculateAreaTriangle())
	s := square{ side: 10}
	fmt.Println("square value :",s.calculateAreaSquare())
	r := rectangle{ length: 10, width: 5}
	fmt.Println("rectangle value :",r.calculateAreaRectangle())

	fmt.Println("----------------")
	fmt.Println("Question 3")
	fmt.Println("----------------")
	hp := handphone{name: "nokia", brand: "nokia", year: 2022, availableColors: []string{"black", "white"}}
	hp.addColor("green")
	hp.addArrayColor("yellow", "blue", "red")
	fmt.Println(hp)

	fmt.Println("----------------")
	fmt.Println("Question 4")
	fmt.Println("----------------")
	// movies := []Movie{}
	movies := make([]Movie, 0)
	addMovieData("Avengers", "Action", 120, 2022, &movies)
	addMovieData("The Matrix", "Action", 120, 1999, &movies)

	
	// for i, movie := range movies {
	// 	fmt.Println("No :", i + 1, "Title :",  movie.title, "Genre :", movie.genre, "Duration :", movie.duration, "Year :", movie.year)
	// }

	for i, movie := range movies {
		fmt.Println("No :", i + 1, movie)
	}

	fmt.Println("----------------")
	fmt.Println("Question 5")
	fmt.Println("----------------")
	student := make([]student, 0)
	// addStudentData("ahmad", "IT", 2022, 3.5, []string{"coding", "learning"}, &student)
	// addStudentData("budi", "IT", 2022, 3.5, []string{"coding", "learning"}, &student)

	addStudentData("budi", "IT", 2022, 3.5, &student ,"coding", "learning")
	addStudentData("ahmad", "IT", 2022, 3.5, &student ,"coding", "learning")
	addStudentData("cindy", "IT", 2022, 3.5, &student ,"coding", "learning")
	// for i, student := range student {
	// 	fmt.Println("No :", i + 1, "Name :",  student.name, "Major :", student.major, "Year :", student.year, "GPA :", student.gpa, "Hobbies :", student.hobbies)
	// }

	for i, student := range student {
		fmt.Println("No :", i + 1, student)
	}


}