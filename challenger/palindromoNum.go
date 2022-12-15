package challenger

func Palindromo(num int) bool {
	temp := num
	rev := 0
	for num > 0 {
		dig := num % 10
		rev = rev*10 + dig
		num = num / 10
	}
	return temp == rev
}
