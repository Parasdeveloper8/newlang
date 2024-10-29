package main

import "fmt"

func main() {
	logins := 4
	if logins > 10 {
		fmt.Println("high")
	} else if logins < 10 && logins > 5 {
		fmt.Println("low")
	} else {
		fmt.Println("very low")
	}
	//initalizing variable at same time
	if notebooks := 5; notebooks > 4 {
		fmt.Println("good")
	}
}
