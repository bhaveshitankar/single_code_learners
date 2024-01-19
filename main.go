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

func (p *Person) ShowPerson() string {
	p.Name = "new name"
	return "I am a function for Person"
}

type Employee struct {
	Person     // Inheritance
	EId        string
	Salary     float32
	Department string
}

func Inheritance() {
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
	a := emp.Person.ShowPerson()
	fmt.Println(a)
}

func Polymorphism() {
	// Polymorphism
	var HoldAnything interface{}
	HoldAnything = 1
	fmt.Println(HoldAnything)
	HoldAnything = "my data" // overring the type of HoldAnything variable to string.
	fmt.Println(HoldAnything)
	// car interface we can audi or mercedes to it
}

func main() {
	/*
		// pointers :
		[block of memory[32bits]][memory location]<-- pointer
		I2 --> [24][56789012345] --> &I2 this returns 56789012345
		*&I2 --> 24
		* --> it will extract value from address
		& --> it will extract address from value
	*/
	var I1 *int // I1 will hold a memory address of an Int--> [0xc00000a0e0][0xc000052020]
	fmt.Println(I1)
	fmt.Println(&I1)
	var I2 int = 24 // I2 will hold value of type int at a memory address
	fmt.Println(&I2)
	I1 = &I2 // &I2 will return address of I2
	fmt.Println(&I1)
	// update values using references.
	fmt.Println(I1)
	*I1 = 36
	fmt.Println(I2)

	fmt.Println("without pointer.. ")
	p1 := Person{}
	fmt.Println(p1)
	fmt.Println(p1.Age)
	p1 = modifyPerson(p1)
	fmt.Println(p1.Age)

	fmt.Println("with pointer.. ")
	p2 := &Person{}
	fmt.Println(p2)
	fmt.Println(p2.Age)
	modifyPersonWithPointer(p2)
	fmt.Println(p2.Age)
}

func modifyPerson(p Person) Person { // over use of memory
	p.Age = 50 // updating copy of p
	return p   // returning the copy.
}

func modifyPersonWithPointer(p *Person) { // p is just a refrence to Person object created in main.
	p.Age = 50 // modifing value at address.
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
