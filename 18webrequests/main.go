package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const Url = "https://devparas.com"

func main() {
	response, err := http.Get(Url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("type of response \n", response)
	defer response.Body.Close()
	databyte, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(databyte))
}
