package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	performpostreq()
}
func performpostreq() {
	//fake payload
	requestBody := strings.NewReader(`
	{
	    "name":"paras",
		"class":"9th"
	}
	`)
	const Myurl = "http://localhost:7800/post"
	response, err := http.Post(Myurl, "text/html", requestBody)
	if err != nil {
		panic(err)
	}
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println("response", string(content))
}
