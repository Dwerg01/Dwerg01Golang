package main

import "fmt"

func main() {
	var largenum int
	var smallnum int
	fmt.Print("please enter a large number: ")
	fmt.Scanln(&largenum)
	fmt.Print("enter a smaller divisor: ")
	fmt.Scanln(&smallnum)
	fmt.Println(largenum, "%", smallnum, "=", largenum%smallnum)

}
