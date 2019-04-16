package main

import (
	"fmt"
)

func rectProps(length, width float64) (float64, float64) {
	var area = length * width
	var perimeter = (length + width) * 2
	return area, perimeter
}
func main() {
	area, _ := rectProps(10.8, 5.6) // 返回值周长被丢弃
	fmt.Printf("Area %f ", area)
	fmt.Println()
}
