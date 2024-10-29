package main

import "fmt"

func main() {
	days := []string{"monday", "tuesday", "wednesday", "thursday", "friday", "saturday", "sunday"}
	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for key, value := range days {
		fmt.Printf("Key is %v and value is %v\n", key, value)
	}

	values := 10
	for values < 100 {
		if values == 99 {
			goto paras
		}
		if values == 50 {
			values++
			continue
			//break
		}
		fmt.Println("value is:", values)
		values++
	}
paras:
	fmt.Println("devparas.com")
}
