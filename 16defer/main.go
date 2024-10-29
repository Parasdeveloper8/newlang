package main

import "fmt"

//defer means to execute at end or reverse order
func main() {
	fmt.Println("hello")
	mydefer()
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
}
func mydefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i, "\n")
	}
}
