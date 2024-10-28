package main

import "fmt"

func main() {
	fmt.Println("hello from array")
	var fruits [4]string
	fruits[0] = "apple"
	fruits[1] = "guava"
	fruits[2] = "lemon"
	fmt.Println(fruits)
	//length of fruits array
	fmt.Println(len(fruits))
	var veges = [3]string{"apple", "muskmelon", "pear"}
	fmt.Println(veges)
}
