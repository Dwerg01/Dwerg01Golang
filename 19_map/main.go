package main

import "fmt"

func main() {
	turbosnaps := map[string]string{
		"H":  "Hydrogen",
		"He": "Helium",
		"C":  "Carbon",
		"Ne": "Neon",
		"Li": "Lithium",
		"Be": "Berilium",
		"B":  "Boron",
	}
	for k, v := range turbosnaps {
		fmt.Println(k, "-", v)
	}
}
