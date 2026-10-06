package main

import "fmt"

func main() {
	a, b, c := 14.5, 82.1, 45.0

	maxVal := a
	if b > maxVal {
		maxVal = b
	}
	if c > maxVal {
		maxVal = c
	}

	minVal := a
	if b < minVal {
		minVal = b
	}
	if c < minVal {
		minVal = c
	}

	fmt.Printf("Dãy số: %.2f, %.2f, %.2f\n", a, b, c)
	fmt.Printf("Giá trị Lớn nhất (Max): %.2f\n", maxVal)
	fmt.Printf("Giá trị Nhỏ nhất (Min): %.2f\n", minVal)
}
