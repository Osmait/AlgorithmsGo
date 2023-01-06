package challenger

func IterativeFactorial(n int) int {
	fact := 1
	for i := 2; i <= n; i++ {
		fact *= i
	}
	return fact
}

func RecurFactorial(n int) int {
	if n == 1 {
		return n
	}
	return RecurFactorial(n-1) * n

}
