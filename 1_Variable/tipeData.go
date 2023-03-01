package main

import "fmt"

func main() {
	// Tipedata standart
	// var mydataType int
	// fmt.Println(mydataType)

	// Tipedata non-standart
	// var mydataType interface{}
	// mydataType = true
	// fmt.Println(mydataType)

	// Tipe data slice
	// var mydataType []string
	// fmt.Println(mydataType)

	// var mydataTypeMap map[string]string
	// fmt.Println(mydataTypeMap)

	// Auto assignment tipe data
	// angka := "1"
	// fmt.Println(angka)

	// Pointer
	mobil := "toyota"
	minifan := &mobil

	mobil = "tesla"
	mobil = "tesla sssss"

	fmt.Println("Mobil: ", mobil)
	fmt.Println("Minifan: addr ", minifan)
	fmt.Println("Minifan: ", *minifan)
}
