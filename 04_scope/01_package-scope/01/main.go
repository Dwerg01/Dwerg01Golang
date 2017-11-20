package main

import "fmt"

var x = 42

func main() {
	fmt.Println(x)
	y := 16
	println(y)
	foo()
}

func foo() {
	fmt.Println(x)

}
