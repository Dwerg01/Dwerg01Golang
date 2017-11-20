package main

import "fmt"

func main() {

	MyMap := map[int]string{
		0: "Good morning!",
		1: "Bonjour!",
		2: "Buenos dias!",
		3: "Bongiorno!",
	}

	for key, val := range MyMap {
		fmt.Println(key, " - ", val)
	}
}
