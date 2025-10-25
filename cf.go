package main

import "fmt"

func main() {
	var fahrenheit float64

	fmt.Println("Enter temperature in Fahrenheit:")
	fmt.Scan(&fahrenheit)

	celsius := (fahrenheit - 32) * 5 / 9

	fmt.Println("Temperature in Celsius:", celsius)
}
