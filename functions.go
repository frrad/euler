package euler

//Implement special mathematical functions

func TriangleNumber(n int) int64 {
	a := int64(n)
	return a * (a + 1) / 2
}

var fibtable = map[int]int64{0: 0, 1: 1}

//Fibonacci returns the nth Fibonacci number. Fibonacci(1) == Fibonacci(2) == 1.
//The last correct value is Fibonacci(92) = 7540113804746346429, with
//Fibonacci(93) = overflow.
func Fibonacci(n int) int64 {
	if n < 0 {
		return -1
	}
	if ans, ok := fibtable[n]; ok {
		return ans
	}
	ans := Fibonacci(n-1) + Fibonacci(n-2)
	fibtable[n] = ans
	return Fibonacci(n)
}
