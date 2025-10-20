package main

import "fmt"

func main() {
	//Opration normal, +, -, *, /, %(modulus)
	var a = 10
	var b = 10
	var c = 5

	var d = a + b - c
	fmt.Println(d)

	//Augmented assignment +=, -=, *=, /=, %=
	var i = 10
	i += 10
	fmt.Println(i)

	i -= 5
	fmt.Println(i)

	//Unary operator ++, --, -, +, !(boolean kebalikan)
	var j = 1
	j++
	fmt.Println(j)

	j++
	fmt.Println(j)

	j--
	fmt.Println(j)

	//perbandingan Opration <,>,<=,>=,==,!=
	var number1 = 10
	var number2 = 20

	var result bool = number1 < number2
	fmt.Println(result)

	//Logical Opration &&(and), ||(or), !(not)
	var a1 = 10
	var a2 = 20
	var a3 = 30
	var result2 bool = (a1 < a2) && (a2 < a3)
	fmt.Println(result2)
}
