package main

import (
	"fmt"
	"sort"
)

func main() {
	var fruitList = []string{"peach", "pear"}
	fmt.Println(fruitList)
	fruitList = append(fruitList, "apple", "mango", "guava")
	fmt.Println(fruitList)
	fruitList = append(fruitList[1:3])
	fmt.Println(fruitList)

	highScores := make([]int, 3)
	highScores[0] = 100
	highScores[1] = 8980
	highScores[2] = 389
	//highScores[3] = 400 out of range
	fmt.Println(highScores)
	highScores = append(highScores, 500, 600, 700)
	fmt.Println(highScores)

	sort.Ints(highScores) //sorting integers in increasing order
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) //checking if sorted

	var courses = []string{"react", "nest", "next", "nuxt", "express"}
	var index int = 2
	courses = append(courses[:index+1], courses[:index+2]...) //... = leave
	fmt.Println(courses)
}
