package main

import "fmt"

// Go is procedural programming language. c
// oops concepts : (c++)
/*
Object oriented programming.

1 . You should be able to define your own types. (Encapsulation)
type Student int
type Student Struct {
	ID int
}

2. Inheritance : you build new type from other existing types.
				 New build type is known as child and the inherited one is parent.
3. Polymorphism : one type can act as different other types in different conditions.
4. Overriding of a function: updating the functionality of a parent class in child class.

*/

// Encapsulation
type myInterface interface {
	f1()
	f2()
}

type myList []int

type MyString string
type MyStruct struct {
	A int
	B string
}

type Kitchen struct {
	Equipments []Equipment // map[string]float32, []int
	Type       string
	Id         string
}

type Equipment struct {
	Name string
	// Description string
	Price float32
}

type MyNewStruct MyStruct

func Encapsulation() {
	// Encapsulation
	MyString := "xyz This is Encapsulation"
	myNewStruct := MyNewStruct{5, "New Struct"}
	myStruct := MyStruct{5, "New Struct1"}
	newStruct := MyStruct{A: myNewStruct.A, B: myNewStruct.B}
	if myStruct == newStruct {
		fmt.Println(MyString, myNewStruct, myStruct)
	}
}

// person --> Employee

type Hobby struct {
	Name        string `Json:"name,omitempty"`
	Description string
	Level       int
	Place       string
}
type Person struct {
	Name    string
	Age     int
	Place   string
	Hobbies []Hobby
}

type Employee struct {
	Person     // Inheritance
	EId        string
	Salary     float32
	Department string
}

func main() {
	// Inheritance
	emp := Employee{
		Person: Person{
			Name: "Kiran",
			// Age:   32,
			Place: "Pune",
			// Hobbies: []Hobby{
			// 	{
			// 		Name:        "singing",
			// 		Description: "he can sing",
			// 		Place:       "US",
			// 		Level:       7,
			// 	},
			// },
		},
		EId:        "userId1",
		Salary:     50000.00,
		Department: "Manager",
	}
	fmt.Println(emp)
	fmt.Println(emp.EId)
}

//
/*

{
	"key1": "value1",
	"key1": [
		"data1", "data2"
	]
}

*/
