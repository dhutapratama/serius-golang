package main

import "fmt"

func main() {
	a := 1
	b := &a

	fmt.Println("A: ", a)
	fmt.Println("B: ", *b)
}
