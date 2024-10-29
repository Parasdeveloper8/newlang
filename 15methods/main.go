package main

import "fmt"

func main() {
	person := User{"paras", "pp@gmail.com", 14}
	fmt.Println(person)
	fmt.Printf("Type of person %+v\n", person)
	person.GetName()
}

type User struct {
	Name  string
	Email string
	Age   int
}

func (u User) GetName() {
	fmt.Println("Name is ", u.Name)
}
