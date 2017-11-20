package main

import "fmt"

func finder(lar ...int) int {
	fmt.Printf("%T \n", lar)
	var x int
	for _, v := range lar {
		if v > x {
			x = v
		}
	}
	return x
}
func main() {
	x := finder(90, 10, 0, 40, 330)
	fmt.Println(x)
}
