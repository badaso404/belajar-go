package main

import "fmt"

func main() {
	var names [2]string
	names[0] = "Bagas"
	names[1] = "Pratama"

	fmt.Println(names)

	var values = [...]int{
		10,
		20,
		30,
		40,
		50,
		60,
	}

	fmt.Println(values)
	fmt.Println("Panjang array values :", len(values))
}
