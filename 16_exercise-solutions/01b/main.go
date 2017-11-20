package main

import "fmt"

func main() {
	h, calculator := calculator(5)
	fmt.Println(h, calculator)
}
func calculator(x int) (float64, bool) {
	return float64(x) / 2, x%2 == 0
}
