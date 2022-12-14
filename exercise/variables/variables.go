//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier

package main

import "fmt"

func main() {

	var favoriteColor string = "pink and black" // var keyword
	fmt.Println("My favorites colors are", favoriteColor)
	
	var	year,age = 1999, 23 // compound assignment
	fmt.Printf("I burn in %v and I'm %v years old.\n", year, age)

	var (
		firstInitial	string = "R" // block assignment
		lastInitial		string = "C"
	)
	fmt.Printf("My first initial are %v and my last initial are %v\n", firstInitial, lastInitial)

	var daysOfLife int

	daysOfLife = 365 * age
	fmt.Printf("I have %v days of life until today.", daysOfLife)

}
