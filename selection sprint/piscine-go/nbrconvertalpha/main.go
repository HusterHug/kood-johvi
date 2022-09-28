package main

import (
	"os"

	"github.com/01-edu/z01"
)

func strTOint(s string) int {
	number := 0
	for _, r := range s {
		if '0' <= r && r <= '9' {
			start := 0
			for i := '1'; i <= r; i++ {
				start++
			}
			number = number*10 + start
		}
	}
	return number
}

func main() {
	program := os.Args

	length := len(program)

	Capitalized := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Lowercase := []rune("abcdefghijklmnopqrstuvwxyz")

	result := []rune{}

	if length == 1 {
	} else {
		if program[1] == "--upper" {
			for i := 2; i < length; i++ {
				number := strTOint(program[i])
				if number > 26 || number < 1 {
					result = append(result, rune(32))
				} else {
					result = append(result, Capitalized[number-1])
				}
			}
		} else {
			for i := 1; i < length; i++ {
				number := strTOint(program[i])
				if number > 26 || number < 1 {
					result = append(result, rune(32))
				} else {
					result = append(result, Lowercase[number-1])
				}
			}
		}
		for _, j := range result {
			z01.PrintRune(j)
		}
		z01.PrintRune('\n')
	}
}
