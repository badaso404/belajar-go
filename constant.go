package main

import "fmt"

func main() {
	const firstName = "Bagas"
	const lastName = "Pratama"
	fmt.Println(firstName)
	fmt.Println(lastName)

	const (
		nickname  = "Bagas Pratama"
		birthYear = 2004
		isMarried = false
		height    = 175.5
	)
	fmt.Println(nickname)
	fmt.Println(birthYear)
	fmt.Println(isMarried)
	fmt.Println(height)
}
