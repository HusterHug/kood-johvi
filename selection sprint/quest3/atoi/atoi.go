package piscine

func Atoi(s string) int {
	str := []rune(s)
	l := len(s)
	a := 0
	for i := range str {
		l = i + 1
	}
	if l == 0 {
		return 0
	}
	b := 1
	for k := 0; k < l-1; k++ {
		if str[k] == '+' || str[k] == '-' {
			continue
		}
		b *= 10
	}
	for i := range str {
		if (str[0] == '-' || str[0] == '+') && i == 0 {
			continue
		}
		if str[i] < '0' || str[i] > '9' {
			return 0
		}
		nr := 0
		for j := '0'; j < str[i]; j++ {
			nr++
		}
		a += nr * b
		b /= 10
	}
	if str[0] == '-' {
		a *= -1
	}
	return a
}
