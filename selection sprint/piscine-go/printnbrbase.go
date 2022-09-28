package piscine

import (
	"github.com/01-edu/z01"
)

func PrintNbrBase(number int, base string) {
	length := len(base)

	runes := []rune(base)
	valid := true

	if length < 2 {
		valid = false
	} else {
		for i := 0; i < length; i++ {

			if runes[i] == '-' || runes[i] == '+' {
				valid = false
			}

			for j := i + 1; j < length; j++ {
				if runes[i] == runes[j] {
					valid = false
				}
			}
		}
	}
	if !valid {

		z01.PrintRune('N')
		z01.PrintRune('V')

	} else {
		if number == 0 {
			z01.PrintRune(runes[0])
		} else {

			if number < 0 {
				z01.PrintRune('-')
			}

			printNbrBaseRecurrsive(number, runes, length)

		}
	}
}

func printNbrBaseRecurrsive(nr int, runes []rune, length int) {
	if nr/length != 0 {
		printNbrBaseRecurrsive(nr/length, runes, length)
	}

	index := nr % length
	if index < 0 {
		index = -index
	}

	z01.PrintRune(runes[index])
}
