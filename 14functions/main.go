package main

import "fmt"

func main() {
	sum := adder(2, 5)
	fmt.Println(sum)

	sum2, msg := proadder(1, 2, 3, 4)
	fmt.Println(sum2, msg)
}

func adder(val1 int, val2 int) int {
	// Arguments data types and data type which will return
	return val1 + val2
}

func proadder(values ...int) (int, string) {
	total := 0
	for _, va := range values {
		total += va
	}
	return total, "hi"
}
