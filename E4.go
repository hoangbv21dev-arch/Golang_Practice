package main

import "fmt"

func main() {
	num := -8

	// Cú pháp 'if initialization; condition' đặc sản của Go
	if rem := num % 2; rem == 0 {
		fmt.Printf("%d là số Chẵn\n", num)
	} else {
		fmt.Printf("%d là số Lẻ\n", num)
	}

	// Biến 'rem' ở trên đã tự động hết vòng đời (out of scope), không gây ô nhiễm biến ngoài!

	if num > 0 {
		fmt.Printf("%d là số Dương\n", num)
	} else if num < 0 {
		fmt.Printf("%d là số Âm\n", num)
	} else {
		fmt.Println("Số 0 (Không âm cũng không dương)")
	}
}
