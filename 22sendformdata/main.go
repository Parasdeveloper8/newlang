package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	performpostreqformdata()
}
func performpostreqformdata() {
	const newurl = "http://localhost:7800/form"
	//form data

	data := url.Values{}
	data.Add("name", "paras prajapat")
	data.Add("class", "9")

	response, err := http.PostForm(newurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))
}
