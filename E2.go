package main

import "fmt"

const (
	StatusOK         = 200
	StatusBadRequest = 400
	StatusNotFound   = 404
)

const (
	Step1 = iota
	Step2
	Step3
)

var globalVar = 100

func main() {

	var a int

	var b = 10

	c := 20

	fmt.Println("--- BIẾN ---")
	fmt.Printf("a = %d (mặc định)\n", a)
	fmt.Printf("b = %d\n", b)
	fmt.Printf("c = %d\n", c)
	fmt.Printf("globalVar = %d\n", globalVar)

	fmt.Println("\n--- HẰNG SỐ ---")
	fmt.Printf("Mã HTTP lỗi: %d\n", StatusNotFound)
	fmt.Printf("Step 3 có giá trị là: %d\n", Step3)
}
