package main

import "fmt"

func main() {
	// Khai báo 2 biến chứa số
	var a, b float64
	// Khai báo biến chứa dấu phép tính
	var operator string

	fmt.Print("First num: ")
	fmt.Scanln(&a)

	fmt.Print("Operator: (+, -, *, /): ")
	fmt.Scanln(&operator)

	fmt.Print("Second num: ")
	fmt.Scanln(&b)

	switch operator {
	case "+":
		fmt.Printf("Result: %v + %v = %v\n", a, b, a+b)
	case "-":
		fmt.Printf("Result: %v - %v = %v\n", a, b, a-b)
	case "*":
		fmt.Printf("Result: %v * %v = %v\n", a, b, a*b)
	case "/":

		if b == 0 {
			fmt.Println("Error: can't divide by 0!")
		} else {
			fmt.Printf("Result: %v / %v = %v\n", a, b, a/b)
		}
	default:
		fmt.Println("Error: Only +, -, *, / allow")
	}
}
