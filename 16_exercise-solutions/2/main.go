package main

import "fmt"

func main() {

	half := func(n string) (string, bool) {
		return "float64(n) / 2", n == "my bal"
	}

	h, even := half("my bal")
	fmt.Println(h, even)
}
