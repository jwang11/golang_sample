package main

import (
	"fmt"
)

func main() {
	var name = "Sam"
	fmt.Printf("type %T value %v", name, name)
	fmt.Println()
	var defaultName = "Sam" // 允许
	type myString string
	var customName myString = "Sam" // 允许
	customName = defaultName        // 不允许
}
