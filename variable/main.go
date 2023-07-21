package main

import "fmt"

func main() {
	var word string         // a string
	var number int          // a number (integer)
	var unumber uint        // an unsigned number (integer)
	var floatNumber float32 // a float number

	// Declare with set value
	name := "Iman Maliki" // set var name value to "Iman Maliki" with string type
	age := 24             // set var age with 24 and type is int

	word, number, unumber = "Software Dev", -1, 90
	floatNumber = 9.7

	println(
		word,
		number,
		unumber,
		floatNumber,
		name,
		age,
	)

	// Array
	x := make([]int, 0)
	y := [3]string{"idx-0", "idx-1", "idx-2"}

	x = append(x, 2)
	x = append(x, 100)

	fmt.Println(
		x,
		y,
	)
}
