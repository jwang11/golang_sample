package main

import (
	"fmt"
)

type Employee struct {
	name     string
	salary   int
	currency string
}

/*
  displaySalary() 方法将 Employee 做为接收器类型
*/
func (e Employee) displaySalary() int {
	fmt.Printf("Salary of %s is %s%d", e.name, e.currency, e.salary)
	return 0
}

func main() {
	emp1 := Employee{
		name:     "Sam Adolf",
		salary:   5000,
		currency: "$",
	}
	emp1.displaySalary() // 调用 Employee 类型的 displaySalary() 方法
}
