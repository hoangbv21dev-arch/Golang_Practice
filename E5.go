package main

import "fmt"

func main() {
	var c float64 = 37.0 // Thân nhiệt cơ thể

	f := (c * 9.0 / 5.0) + 32.0
	k := c + 273.15

	fmt.Printf("Celsius   : %.2f °C\n", c)
	fmt.Printf("Fahrenheit: %.2f °F\n", f)
	fmt.Printf("Kelvin    : %.2f K\n", k)
}
