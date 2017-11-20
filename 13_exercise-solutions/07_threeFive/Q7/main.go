package main

import "fmt"

func main() {
	var div3 int
	//	var div5 int
	for i := 0; i < 1000; i++ {
		if i%3 == 0 {
			div3 += i
		} else if i%5 == 0 {
			div3 += i
		}
	}
	fmt.Println(div3)
}
