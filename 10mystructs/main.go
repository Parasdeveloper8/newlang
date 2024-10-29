package main

import "fmt"

func main() {
	person := User{"paras", "pp@gmail.com", 14}
	fmt.Println(person)
	fmt.Printf("Type of person %+v", person)
}

type User struct {
	Name  string
	Email string
	Age   int
}
