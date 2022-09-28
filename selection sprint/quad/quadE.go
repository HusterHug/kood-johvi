package piscine

import "github.com/01-edu/z01"

func QuadE(x int, y int) {
	/*
		A "switch" is used for better readability and simplicity.
		Switch without an expression is an alternate way to express if/else logic.
	*/
	switch {
	case x > 1 && y > 1:
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
		for i := 0; i < y-2; i++ {
			z01.PrintRune('B')
			for j := 0; j < x-2; j++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('B')
			z01.PrintRune('\n')
		}
		z01.PrintRune('C')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('A')
		z01.PrintRune('\n')
	case x <= 1 && y <= 1:
		z01.PrintRune('A')
		z01.PrintRune('\n')

	case x <= 1:
		z01.PrintRune('A')
		z01.PrintRune('\n')
		for i := 0; i < y-2; i++ {
			z01.PrintRune('B')
			z01.PrintRune('\n')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
	case y <= 1:
		z01.PrintRune('A')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('B')
		}
		z01.PrintRune('C')
		z01.PrintRune('\n')
	}
}
