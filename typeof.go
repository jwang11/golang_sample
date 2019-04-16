package main

import "fmt"

func main() {
	v := "hello world"
	fmt.Println(typeof(v))
}
func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	//... etc
	default:
		_ = t
		return "unknown"
	}
}
