package piscine

import "github.com/01-edu/z01"

func QuadC(x, y int) {
	if x > 0 && y > 0 { //main if, it checks if value is bigger than 0, if value is smaller than 1 then it returns nothing
		if x == 1 && y == 1 {// if x and y are just 1 then it prints only A
			z01.PrintRune('A')
			z01.PrintRune('\n')
			return
		} else if x == 1 {// if x is 1 and y is another bigger number then it prints A at the start and C at the end and B in the middle
			z01.PrintRune('A')//start
			for t := 1; t < y-1; t++ {//middle
				if t <= 0 {
					break
				} else {
					z01.PrintRune('\n')
					z01.PrintRune('B')
				}
				if y == 1 {//checks again if y is 1, if y is 1 then it prints A and returns
					z01.PrintRune('\n')
					z01.PrintRune('A')
					return
				}
			}
			z01.PrintRune('\n')//end
			z01.PrintRune('C')
			z01.PrintRune('\n')
		} else if y == 1 {// if y is 1 and X is another bigger number then it prints A in the start, B at the middle and A again in the end
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {//this for loop is for the upper bars middle so it prints B 
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')//upper bar end
			return
		} else {// if all values are random then it uses this part. i made a basic layot that the upper bar starts and ends with A always
			z01.PrintRune('A')
			for i := 1; i < x-1; i++ {
				z01.PrintRune('B')
			}
			z01.PrintRune('A')
			z01.PrintRune('\n')
			for t := 1; t < y-1; t++ {//this is for middle stuff so it fills the sides with B
				z01.PrintRune('B')

				for j := 1; j < x-1; j++ {
					z01.PrintRune(' ')
				}
				z01.PrintRune('B')
				z01.PrintRune('\n')
			}//middle part end
			//
			// this is for lower bar. it always prints c at the start and at the end
			z01.PrintRune('C')//lower bar start, prints c at the start
			for i := 1; i < x-1; i++ {//start of lower bar middle, prints B between lower bar conners
				z01.PrintRune('B')
			}//end of lower bar middle
			z01.PrintRune('C')//ower bar end, types C and a new line to complete the square
			z01.PrintRune('\n')
//lower bar end
		}
	}
}
