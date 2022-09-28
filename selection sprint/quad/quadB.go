package piscine

import (
	"github.com/01-edu/z01"
)

// Instructions
// Write a function QuadB that prints a valid rectangle of width x and of height y.

// The function must draw the rectangles as in the examples.

// If x and y are positive numbers, the program should print the rectangles as seen in the examples,
//  otherwise, the function should print nothing.

func QuadB(x, y int) {

	// 1. First line starts
	// if x or y is negative or 0, then don't print anything
	if x <= 0 || y <= 0 {
		return
	}
	// if x and y are 1, then just print the first symbol
	if x == 1 && y == 1 {
		z01.PrintRune('/')
		z01.PrintRune('\n')
		return
	}
	z01.PrintRune('/')
	// print 2 asterix less than x value as we have slashes for both sides horisontally
	for i := 1; i < x-1; i++ {
		z01.PrintRune('*')
	}
	if x > 1 {
		z01.PrintRune('\\')
	}
	z01.PrintRune('\n')
	// First line ends
	// 2. Mid section lines start starts
	// Print 2 asterix less than y value, as we have slashes for both sides vertically
	for i := 1; i < y-1; i++ {
		// This asterix is printed only if x > 1 because we print another asterix anyway if y > 1 and we don't want double in that case
		if x > 1 {
			z01.PrintRune('*')
		}
		// Once asterix is printed print empty space 2 times less than x value horisontally
		for j := 0; j < x-2; j++ {
			z01.PrintRune(' ')
		}
		z01.PrintRune('*')
		z01.PrintRune('\n')
		// Mid section lines end
	}
	// Last lines start
	if y > 1 {
		z01.PrintRune('\\')
		// Print 2 asterix less than x value as slashes both sides
		for i := 1; i < x-1; i++ {
			z01.PrintRune('*')
		}
	}
	// We only need this slash if x is bigger than 1 
	if x > 1 {
		z01.PrintRune('/')
	}
	// We already print line if y == 1
	if y > 1 {
		z01.PrintRune('\n')
	}
}
