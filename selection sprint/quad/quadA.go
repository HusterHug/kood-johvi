package piscine

import "github.com/01-edu/z01"

func QuadA(x, y int) {
	/*Subtract 2 from X and Y, as there is an "o" around the edges. If X and Y (5,5) then we get:
	o-----o
	|     |
	|     |
	|     |
	|     |
	|     |
	o-----o
	But this is the wrong answer. That's why we subtract 2 from x and y. Correct answer:
	o---o
	|   |
	|   |
	|   |
	o---o
	*/
	switch {
	case x > 1 && y > 1:
		z01.PrintRune('o')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
		for i := 1; i <= y-2; i++ {
			z01.PrintRune('|')
			for j := 0; j < x-2; j++ {
				z01.PrintRune(' ')
			}
			z01.PrintRune('|')
			z01.PrintRune('\n')
		}
		z01.PrintRune('o')
		for i := 1; i <= x-2; i++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
	case x <= 1 && y <= 1:
		z01.PrintRune('o')
		z01.PrintRune('\n')
	case x <= 1:
		z01.PrintRune('o')
		z01.PrintRune('\n')
		for i := 0; i < y-2; i++ {
			z01.PrintRune('|')
			z01.PrintRune('\n')
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
	case y <= 1:
		z01.PrintRune('o')
		for i := 0; i < x-2; i++ {
			z01.PrintRune('-')
		}
		z01.PrintRune('o')
		z01.PrintRune('\n')
	}
}
