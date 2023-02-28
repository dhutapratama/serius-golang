package main

import "fmt"

type DLabs int

func standar() {
	// Deklarasi Standar
	var varInt int

	// Deklarasi standar dengan value
	var varIntWithValue int = 1

	// Deklarasi dengan :=, otomatis asignment
	intAuto := 1
	floatAuto := 1.0
	stringAuto := "string"

	// Custom vars
	var customVar DLabs

	fmt.Println(varInt, varIntWithValue, intAuto, floatAuto, stringAuto, customVar)
}

// Tipe data non standar
func nonStandardType() {
	// Deklarasi standar
	var sliceString []string
	var mapString map[string]string
	var iface interface{}

	// Auto assignment
	sliceStringAuto := []string{
		"a", "b", "c",
	}
	mapStringAuto := map[string]string{
		"nama": "DLabs",
	}

	iface = "string"
	iface = 1
	iface = 18
	iface = true

	fmt.Println("", sliceString, mapString, sliceStringAuto, mapStringAuto, iface)
}

type myCustomType interface{}

func nullableType() {
	var intNull = myCustomType(nil)
	fmt.Println(intNull)
}

func pointer() {
	var intPointer *int = new(int)

	intInjector := 1
	intPointer = &intInjector

	fmt.Println(intPointer)
}
