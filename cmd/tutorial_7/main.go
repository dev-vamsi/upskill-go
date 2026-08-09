package main

import "fmt"

func main() {
	// structs
	var vamsi Student = Student{Name: "Vamsi", Age: 24, Email: "vamsi@gmail.com"}
	fmt.Println(vamsi.Name, "-", vamsi.Age)
	vamsi.Age = 23
	fmt.Println(vamsi.Name, "-", vamsi.Age)

	var krishna Student = Student{"Krishna", 16, "krishna@gmail.com", Section{"Elite"}}
	fmt.Println(krishna.Name, "-", krishna.SectionName, "-", krishna.Section.SectionName)

	fmt.Println(vamsi.checkIfAdult())
	fmt.Println(krishna.checkIfAdult())

	// interface
	var s Shape

	s = Circle{radius: 5}
	fmt.Println("C Area:", s.Area())
}
