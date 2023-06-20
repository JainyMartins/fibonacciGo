package fibonacci


func Fibonacci(n int) int {
	f := 0
	ant := 0

	for i := 1; i <= n; i++ {
		if i == 1 {
			f = 1
			ant = 0
		} else {
			f += ant
			ant = f - ant
		}
	}
	return f
}