package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {

	fmt.Println("Hello, Golang!")
	fmt.Println("-------------------------------------------------")

	user := User{Name: "Alice", Age: 25}
	number := 42

	fmt.Printf("%%v  (Mặc định)       : %v\n", user)

	fmt.Printf("%%+v (Kèm tên trường) : %+v\n", user)

	fmt.Printf("%%#v (Cú pháp Go)     : %#v\n", user)

	fmt.Printf("%%T  (Kiểu dữ liệu)   : %T\n", user)

	fmt.Printf("%%b  (Nhị phân của %d)  : %b\n", number, number)

	fmt.Printf("%%x  (Hex của %d)       : %x\n", number, number)
}
