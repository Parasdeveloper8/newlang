package main

import "fmt"

const Publicvar = "public" //public variable because of capital letter in variable at first
func main() {
	var username string = "paras prajapat"
	fmt.Println(username)
	//%T means type \n means new line
	fmt.Printf("type is %T \n", username)

	var isHuman bool = true
	fmt.Println(isHuman)
	fmt.Printf("type is %T \n", isHuman)

	var number int = 255
	fmt.Println(number)
	fmt.Printf("type %T \n", number)

	var numberFloat float64 = 255.6655887872979809289890
	fmt.Println(numberFloat)
	fmt.Printf("type %T \n", numberFloat)

	var emptystr string
	fmt.Println(emptystr)
	fmt.Printf("type %T \n", emptystr)

	//implicit type
	var name = "paras"
	fmt.Println(name)

	//no var style
	user := "aras"
	fmt.Println(user)
}
