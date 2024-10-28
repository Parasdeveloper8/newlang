package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time study")
	presentTime := time.Now()
	fmt.Println(presentTime)
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday"))
	createdate := time.Date(2024, time.October, 28, 23, 51, 0, 0, time.UTC)
	fmt.Println(createdate)
}
