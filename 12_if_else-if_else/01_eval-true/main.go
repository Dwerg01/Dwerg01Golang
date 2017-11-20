package main

import "fmt"

func main() {

	x := 1

	if x > 1 {
		fmt.Println("This ran")
	}

	if x < 2 {
		fmt.Println("This did not run")
	}
}
