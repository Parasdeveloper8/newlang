package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

const Myurl string = "http://localhost:7800/getdata"

func main() {
	response, err := http.Get(Myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	var responsestring strings.Builder
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	bytecount, _ := responsestring.Write(content)
	fmt.Println("status code", response.StatusCode)
	fmt.Println("Content length", response.ContentLength)
	fmt.Println("bytecount is", bytecount)
	fmt.Println(responsestring.String())
	//fmt.Println(string(content))
	//you can use any method
}
