package main

import (
	"fmt"
	"net/url"
)

const Url string = "https://devparas.com/"

func main() {
	result, _ := url.Parse(Url)
	fmt.Println(result)
	fmt.Printf("type of url : %T", result)
}
