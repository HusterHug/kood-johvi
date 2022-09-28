// Instructions

// Write a function that receives three arguments:

// nbr: A string representing a numberic value in a base.

// baseFrom: A string representing the base nbr it's using.

// baseTo: A string representing the base nbr should be represented in the returned value.

// Only valid bases will be tested.

// Negative numbers will not be tested.

// Expected function

// func ConvertBase(nbr, baseFrom, baseTo string) string {

// }
// Usage

// Here is a possible program to test your function :

// package main

// import (
// 	"fmt"
// 	"piscine"
// )

// func main() {
// 	result := piscine.ConvertBase("101011", "01", "0123456789")
// 	fmt.Println(result)
// }
// And its output :

// $ go run .
// 43
// $

package piscine

func ConvertBase(nbr, baseFrom, baseTo string) string {
	numberLength := 0
	baseFromLength := 0
	baseToLength := 0
	answer := ""

	result := map[rune]int{}

	for c := range nbr {
		numberLength = c + 1
	}

	for i, c := range baseFrom {
		result[c] = i
		baseFromLength = i + 1
	}

	for c := range baseTo {
		baseToLength = c + 1
	}

	power := 1
	count := 0
	for i := numberLength - 1; i >= 0; i-- {
		count = count + result[[]rune(nbr)[i]]*power
		power = power * baseFromLength
	}

	for count != 0 {
		answer = string(baseTo[count%baseToLength]) + answer
		count = count / baseToLength
	}
	return answer
}
