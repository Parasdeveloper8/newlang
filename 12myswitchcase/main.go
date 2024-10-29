package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("welcome to switches")

	rand.Seed(time.Now().UnixNano())
	dicenumber := rand.Intn(6) + 1
	fmt.Println("dice number is ", dicenumber)

	switch dicenumber {
	case 1:
		fmt.Println("open your key")
		fallthrough

	case 2, 3, 4, 5:
		fmt.Println("go")
		fallthrough

	case 6:
		fmt.Println("go twice")
	}
}
