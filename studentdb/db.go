package main

import (
	"fmt"
	"strconv"
)

type Person struct {
	firstName string
	lastName  string
	Age       uint
}

type Student struct {
	studentID uint64
	gpa       float32
	standing  string
	person    *Person
	cardsLost uint64
}

func (p *Person) FullName() (fullName string) {
	fullName = p.firstName + p.lastName
	return
}

func (s *Student) PadCardsLost() (paddedCardsLost string) {
	paddedCardsLost = strconv.Itoa(int(s.cardsLost))
	if s.cardsLost < 10 {
		paddedCardsLost = "0" + strconv.Itoa(int(s.cardsLost))
	}
	return
}
func (s *Student) PrintCard() (magstripData string) {
	magstripData = "12345" + strconv.Itoa(int(s.studentID)) + s.PadCardsLost()
	return
}
func main() {
	person := Person{firstName: "Arsh"}
	arsh := Student{person: &person, studentID: 31098356, cardsLost: 10}
	fmt.Println(arsh.PrintCard())
	fmt.Println(person.FullName())
}
