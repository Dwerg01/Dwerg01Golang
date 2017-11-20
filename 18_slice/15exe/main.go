package main

import "fmt"

func main() {
	x := 98
	turbo := make([]int, 5, 20)
	turbo = append(turbo, 4, 8, 9)
	turbo[3] = x
	fmt.Println(len(turbo))
	fmt.Println(cap(turbo))
	fmt.Println(turbo)
	fmt.Println(len(turbo))
}
