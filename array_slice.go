package main

import (
	"fmt"
)

func printarray(a [3][2]string) {
	for _, v1 := range a {
		for _, v2 := range v1 {
			fmt.Printf("%s ", v2)
		}
		fmt.Printf("\n")
	}
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("inside function ", num)
}

func main() {
	{
		a := [3]int{12, 78, 50} // short hand declaration to create array
		fmt.Println(a)
	}
	{
		var a [3]int
		a[0] = 12 // array index starts at 0
		a[1] = 78
		a[2] = 50
		fmt.Println(a)
	}
	{
		a := [...]int{12, 78, 50} // ... makes the compiler determine the length
		fmt.Println(a)
	}
	{
		a := [...]string{"USA", "China", "India", "Germany", "France"}
		b := a // a copy of a is assigned to b
		b[0] = "Singapore"
		fmt.Println("a is ", a)
		fmt.Println("b is ", b)
	}
	{
		num := [...]int{5, 6, 7, 8, 8}
		fmt.Println("before passing to function ", num)
		changeLocal(num) //num is passed by value
		fmt.Println("after passing to function ", num)
	}
	{
		a := [...]float64{67.7, 89.8, 21, 78}
		for i := 0; i < len(a); i++ { // looping from 0 to the length of the array
			fmt.Printf("%d th element of a is %.2f\n", i, a[i])
		}
	}
	{
		a := [...]float64{67.7, 89.8, 21, 78}
		sum := float64(0)
		for i, v := range a { //range returns both the index and value
			fmt.Printf("%d the element of a is %.2f\n", i, v)
			sum += v
		}
		fmt.Println("\nsum of all elements of a", sum)
	}

	{
		a := [3][2]string{
			{"lion", "tiger"},
			{"cat", "dog"},
			{"pigeon", "peacock"}, // this comma is necessary. The compiler will complain if you omit this comma
		}
		printarray(a)
		var b [3][2]string
		b[0][0] = "apple"
		b[0][1] = "samsung"
		b[1][0] = "microsoft"
		b[1][1] = "google"
		b[2][0] = "AT&T"
		b[2][1] = "T-Mobile"
		fmt.Printf("\n")
		printarray(b)
	}
	{
		a := [5]int{76, 77, 78, 79, 80}
		var b []int = a[1:4] // creates a slice from a[1] to a[3]
		fmt.Println(b)
	}
	{
		darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
		dslice := darr[2:5]
		fmt.Println("array before", darr)
		for i := range dslice {
			dslice[i]++
		}
		fmt.Println("array after", darr)
	}
	{
		numa := [3]int{78, 79, 80}
		nums1 := numa[:] // creates a slice which contains all elements of the array
		nums2 := numa[:]
		fmt.Println("array before change 1", numa)
		nums1[0] = 100
		fmt.Println("array after modification to slice nums1", numa)
		nums2[1] = 101
		fmt.Println("array after modification to slice nums2", numa)
	}
}
