package main

import "fmt"

func main() {

	Myword := "supercalifra gilistic"
	mySlice := ([]int{1, 3, 5, 7, 9, 11})
	fmt.Printf("%T\n", mySlice)
	fmt.Println(mySlice[1:4])
	fmt.Println(Myword[3:7])
	fmt.Println("Gobsmack"[3])
}
