package main

import "fmt"

/*
  Write a program which prompts the user to enter a floating point number
  and prints the integer which is a truncated version of the floating point
  number that was entered. Truncation is the process of removing the digits
  to the right of the decimal place.
*/
func main() {
	var num float32

	fmt.Print("Enter a floating point number: ")
	fmt.Scan(&num)

	fmt.Printf("The truncated version of the number %f is %d\n", num, int32(num))
}
