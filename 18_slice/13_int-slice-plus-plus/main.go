package main

import "fmt"

func main() {
	mySlice := make([]int, 1, 4)
	fmt.Println(mySlice[0])
	mySlice[0] = 9
	mySlice = append(mySlice, 8)
	mySlice[1] = 7
	fmt.Println(mySlice[1])
	mySlice[1]++
	fmt.Println(mySlice[1])
	fmt.Println(mySlice[0])
	fmt.Println(cap(mySlice), len(mySlice))
}
