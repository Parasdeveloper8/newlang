package main

import "fmt"

func main() {
	var num int = 24
	pointer := &num
	fmt.Println(pointer)
	fmt.Println(&num)
}
