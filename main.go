package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Check if there are enough arguments
	if len(os.Args) < 4 {
		fmt.Println("Usage: go run ex.go from=<unit> to=<unit> value=<number>")
		return
	}

	var fromUnit, toUnit string
	var value float64

	// Loop through the command-line arguments
	for _, arg := range os.Args[1:] {
		parts := strings.Split(arg, "=")
		if len(parts) != 2 {
			continue
		}

		key := parts[0]
		val := parts[1]

		switch key {
		case "from":
			fromUnit = val
		case "to":
			toUnit = val
		case "value":
			v, err := strconv.ParseFloat(val, 64)
			if err != nil {
				fmt.Println("Invalid value:", val)
				return
			}
			value = v
		}
	}

	var result float64

	// Perform conversion
	if fromUnit == "fahrenheit" && toUnit == "celsius" {
		result = (value - 32) * 5 / 9
	} else if fromUnit == "celsius" && toUnit == "fahrenheit" {
		result = (value * 9 / 5) + 32
	} else {
		fmt.Println("Unsupported conversion:", fromUnit, "to", toUnit)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", value, fromUnit, result, toUnit)
}
