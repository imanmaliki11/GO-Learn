package main

import "fmt"

// Func with return an integer
func add(x, y int) int {
	return x+y
}


// Func with return first and last name string
func getName() (string, string) {
	return "Iman", "Maliki"
}


// Func with pointer
func deductSalary(salary *uint, amount uint) {
	*salary -= amount
}

// Func with declarated return
func returnVal() (x int) {
	x = 5
	return // Automatic return x value
}

func main() {
	firstName, lastName := getName()
	fmt.Println(firstName, lastName)

	var salary, loans uint 

	salary = 1000
	loans = 200

	deductSalary(&salary, loans)
	fmt.Println(salary)

	sum := add(5, 7)
	fmt.Println(sum)

	fmt.Println(returnVal())
}
