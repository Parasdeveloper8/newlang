package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("hello")
	welcome := bufio.NewReader(os.Stdin)
	fmt.Println("enter rating of our website")

	//comma ok or err err
	input, _ := welcome.ReadString('\n')
	fmt.Println("Thanks for rating", input)
}
