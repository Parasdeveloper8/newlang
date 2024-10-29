package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "my name is paras prajapat"
	fmt.Println("Lets create a file with go")
	file, err := os.Create("./mylogs.txt")
	if err != nil {
		panic(err) //panic == stop execution further
	}
	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("length is:", length)
	defer file.Close()
	readFile("./mylogs.txt")
}
func readFile(filename string) {
	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("data is \n", string(databyte))
}
