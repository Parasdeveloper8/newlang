package main

import "fmt"

func main() {
	languages := make(map[string]string)
	languages["js"] = "javascript"
	languages["py"] = "python"
	fmt.Println(languages)

	delete(languages, "py")
	fmt.Println(languages)

	//loops in maps
	for key, value := range languages {
		fmt.Printf("for key %v value is %v", key, value)
	}
}
