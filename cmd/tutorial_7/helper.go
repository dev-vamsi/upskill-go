package main

import (
	"fmt"
	"math"
)

type Student struct {
	Name  string
	Age   uint8
	Email string
	// Section Section // won't provide direct access to krishna.SectionName
	Section
}

type Section struct {
	SectionName string
}

func (s Student) checkIfAdult() string {
	if s.Age >= 18 {
		return fmt.Sprintf("%s is an adult", s.Name)
	}
	return fmt.Sprintf("%s is not an adult", s.Name)
}

// Define the interface
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Circle type that implements the Shape interface
type Circle struct {
	radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.radius
}
