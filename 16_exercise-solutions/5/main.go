package main

import "fmt"

func main() {
	array(1, 2, 5)
	array(3, 4, 2)
	area := []int{1, 4, 6, 7}
	array(area...)
	array()
}

func array(number ...int) {
	fmt.Println(number)
}
