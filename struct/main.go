package main

import "fmt"

// Make a struct
type laptop struct {
	model string
	color string
	year  uint
	brand string
}

// Make a struct with a key and value from other struct
type person struct {
	name string
	age uint
	laptop laptop
}

// Make a struct with embeded other struct inside
type employee struct {
	person
	id string
}

//Create a struct with the method
type rect struct {
	height float64
	width float64
}

// ADD AN area() METHOD TO RECT STRUCT
func (r rect) area() float64 {
	return r.height * r.width
}

func printLaptop(l laptop) {
	fmt.Println("\nLaptop details : \n ")
	fmt.Println(l.brand, l.model, l.color, l.year)
	fmt.Println("\n======================\n ")
}

func printPerson(p person) {
	fmt.Println("\n\n\nPerson details : \n ")
	fmt.Println(p.name, p.age)
	if p.laptop != (laptop{}) {
		fmt.Println("\nPerson Laptop : \n ")
		printLaptop(p.laptop)
	} else {
		fmt.Println("\nTHIS PERSON NOT HAVE A LAPTOP\n ")
	}
	fmt.Println("\n======================\n ")
}

func printEmployee(e employee) {
	fmt.Println("\n\n\nEmployee details : \n ")
	fmt.Println(e.id, e.name, e.age)
	if e.laptop != (laptop{}) {
		fmt.Println("\nEmployee Laptop : \n ")
		printLaptop(e.laptop)
	} else {
		fmt.Println("\nTHIS EMPLOYEE NOT HAVE A LAPTOP\n ")
	}
	fmt.Println("\n==================== \n ")
}

func main() {
	// A struct with partial data initiate
	myLaptop := laptop{
		brand: "Asus",
		model: "AP777",
	}

	// Set a data to struct
	myLaptop.brand = "Asus"
	myLaptop.color = "Black"
	myLaptop.year = 2021
	myLaptop.model = "AP3701-UD"

	printLaptop(myLaptop)

	// Initiate a people struct
	me := person{}

	me.name = "Iman Maliki"
	me.age = 24
	me.laptop = myLaptop

	//You can use me.laptop.brand to set brand property on laptop for 'me' person
	me.laptop.brand = "Acer"

	printPerson(me)

	// make an anonymous struct
	car := struct {
		brand string
		model string
	} {
		brand: "Honda",
		model: "FX-3451-L",
	}

	fmt.Println("CAR Details : \n ", car.model, car.brand, "\n\n ")

	// For initiate if you want to add value on extended always use 'struct name : struct name {..set value here..}'
	theEmployee := employee{
		person: person{
			name: "Angelia Gunn",
			age: 27,
		},
		id: "MGM-00208213-2020",
	} 

	theEmployee.name = "Taylor Angelia Gunn";

	printEmployee(theEmployee)

	// STRUCT WITH METHOD
	rectangle := rect {
		width: 10.13,
		height: 7.5,
	}

	fmt.Println("\n== CALL AREA METHON ON RECT STRUCT ==\n ", rectangle.area())
}
