package main

import "fmt"

func main() {
	var number32 int32 = 32768
	var number64 int64 = int64(number32)
	var number16 int16 = int16(number32)

	fmt.Println("number32:", number32)
	fmt.Println("number64:", number64)
	fmt.Println("number16:", number16)

	var name string = "Bagas"
	var e = name[4]
	var eString = string(e)

	fmt.Println(name)
	fmt.Println(e)
	fmt.Println(eString)

}
