//awww shit, here we go again
package piscine

import "github.com/01-edu/z01"

func QuadD(x, y int) {
	if x > 0 && y > 0 {//checks if values are positive
		if x == 1 && y == 1 { //when x = 1 and y = 1 then it prints only A
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if x == 1 { //when x = 1 and y = ~ number then it prints A at the start and A at the end
			z01.PrintRune('A')
			for t := 1; t < y-1; t++ {//prints upper bars middle
				if t <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')
				}//end of upper bars middle
				if y == 1 {//checks if y = 1
					z01.PrintRune('\n')
					z01.PrintRune('C')
					return
				}
			}//upper bars end
			z01.PrintRune('\n')
			z01.PrintRune('A')
			z01.PrintRune('\n')
		} else if y == 1 { //when y == 1 and x == ~ number then it prints A at the start and C at the end and fills between a-s with B
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {//fills between a-s with B
				z01.PrintRune('B')
			}
			z01.PrintRune('C')//prints c at the end of the line
			z01.PrintRune('\n')
			return
		} else { // this else is for just other variations of the sqare so A at top left, C at top right, A at bottom left and C at bottom right
			z01.PrintRune('A')//upper left conner
			for i := 1; i < x-1; i++ {//fill between conners with B
				z01.PrintRune('B')
			}
			z01.PrintRune('C')//lower left corner
			z01.PrintRune('\n')
			for t := 1; t < y-1; t++ {
				z01.PrintRune('B')

				for j := 1; j < x-1; j++ {//for squares middle 
					z01.PrintRune(' ')//prints spaces inside of the square
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}
			z01.PrintRune('A')//upper right corner
			for i := 1; i < x-1; i++ {//fill the emptynes
				z01.PrintRune('B')
			}
			z01.PrintRune('C')//lower right corner
			z01.PrintRune('\n')

		}
	}
}
