package euler

import (
	"math"
	"strconv"
	"strings"
)

func ConcatanInt(a int64, b int64) int64 {

	//wrong string conversion
	x, _ := strconv.Atoi(strconv.FormatInt(a, 10) + strconv.FormatInt(b, 10))
	return int64(x)
}

//unnecessarily slow. short though!
func IsSquare(x int64) bool {
	return int64(math.Sqrt(float64(x)))*int64(math.Sqrt(float64(x))) == x
}

//removes duplicate entries in a SORTED vector.
func RemoveDuplicates(input []int64) []int64 {

	if len(input) == 1 {
		return input
	}

	if input[0] == input[1] {
		return RemoveDuplicates(append(input[1:]))
	}

	return append(input[:1], RemoveDuplicates(input[1:])...)

}

func Totient(n int64) int64 {
	if n < 2 {
		return 0
	}
	if n < totientTableLength && totientTable[n] != 0 {
		return totientTable[n]
	}

	factors := Factor(n)

	if factors[0] == factors[len(factors)-1] {
		answer := IntExp(factors[0], int64(len(factors))) - IntExp(factors[0], int64(len(factors)-1))
		if n < totientTableLength {
			totientTable[n] = answer
		}
		return answer
	}

	for i := 0; i < len(factors); i++ {
		if factors[i] != factors[0] {
			split := IntExp(factors[0], int64(i))
			answer := Totient(split) * Totient(n/split)
			if n < totientTableLength {
				totientTable[n] = answer
			}
			return answer
		}
	}

	//bad things have happenned if we're here
	return 0

}

func Sort(word string) string {
	wordtable := strings.Split(word, "")
	for j := 0; j < len(word); j++ {

		for i := 0; i < len(word)-1; i++ {
			if wordtable[i] < wordtable[i+1] {
				temp := wordtable[i]
				wordtable[i] = wordtable[i+1]
				wordtable[i+1] = temp
			}
		}
	}
	return strings.Join(wordtable, "")
}

func SortInts(list []int) []int {

	for j := 0; j < len(list); j++ {

		for i := 0; i < len(list)-1; i++ {
			if list[i] < list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}
	return list
}

func SortLInts(list []int64) []int64 {

	for j := 0; j < len(list); j++ {

		for i := 0; i < len(list)-1; i++ {
			if list[i] < list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}
	return list
}

func SortUInt64(list []uint64) []uint64 {
	for j := 0; j < len(list); j++ {

		for i := 0; i < len(list)-1; i++ {
			if list[i] > list[i+1] {
				temp := list[i]
				list[i] = list[i+1]
				list[i+1] = temp
			}
		}
	}
	return list
}

func SortInt(input int64) int64 {

	swapped, _ := strconv.ParseInt(Sort(strconv.FormatInt(input, 10)), 10, 64)

	return swapped

}

func IsPandigital(n int64) bool {

	height := 1 + int64(math.Log10(float64(n)))

	output := int64(0)

	for i := int64(1); i < height+1; i++ {
		current := int64(1)
		for j := int64(1); j < i; j++ {
			current *= 10
		}
		output += (current * (i))
	}

	return output == SortInt(n)
}

func UInt64Prime(n uint64) uint64 {
	return uint64(Prime(int64(n)))
}

func Exp2(n int) int64 {
	answer := int64(1)
	for i := 0; i < n; i++ {
		answer *= 2
	}
	return answer
}

//Number of distinct members of a sorted list
func DistinctNumber(list []int64) int {
	current := list[0]
	total := 1
	for i := 1; i < len(list); i++ {
		if current != list[i] {
			total++
		}
		current = list[i]
	}
	return total

}

func ArePermutations(a int64, b int64) bool {
	A := strconv.FormatInt(a, 10)
	B := strconv.FormatInt(b, 10)

	length := len(A)
	list1 := make([]byte, length)
	list2 := make([]byte, length)

	if len(A) != len(B) {
		return false
	}

	for i := 0; i < length; i++ {
		list1[i] = A[i]
		list2[i] = B[i]
	}

	for i := 0; i < length; i++ {
		flag := false

		for j := 0; j < length; j++ {
			if flag == false && list1[i] == list2[j] {
				list2[j] = 0
				flag = true
			}

		}
		if flag == false {
			return false
		}
	}
	return true

}

func MinInt(m, n int) int {
	if m < n {
		return m
	}
	return n
}

func Min(m int64, n int64) int64 {
	if m < n {
		return m
	}
	return n
}

func MaxInt(m, n int) int {
	if m < n {
		return n
	}
	return m
}

func Max(m int64, n int64) int64 {
	if m < n {
		return n
	}
	return m
}

//Adds fractions quickly, may unreduce
func FastFracAdd(num1 int64, den1 int64, num2 int64, den2 int64) (num int64, den int64) {
	return num1*den2 + num2*den1, den1 * den2
}

//Given two reduced fractions, returns their reduced sum
func FracAdd(num1 int64, den1 int64, num2 int64, den2 int64) (num int64, den int64) {
	gcd := GCD(den1, den2)
	d1 := den1 / gcd
	d2 := den2 / gcd

	den = d1 * d2 * gcd
	num = num1*d2 + num2*d1

	return
}

func FracReduce(num int64, den int64) (int64, int64) {
	gcd := GCD(num, den)
	return num / gcd, den / gcd
}

func DigitSum(N int64) (sum int) {

	sum = 0
	word := strconv.FormatInt(N, 10)

	for i := 0; i < len(word); i++ {
		x, _ := strconv.Atoi(string(word[i]))
		sum += x
	}
	return
}

func StringReverse(a string) string {
	b := ""
	for i := len(a) - 1; i >= 0; i-- {
		b += string(a[i])
	}
	return b
}

func IntReverse(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	s = StringReverse(s)
	x, _ := strconv.ParseInt(s, 10, 64)
	return x
}

func IsPalindrome(n int64) bool {
	if n == IntReverse(n) {
		return true
	}
	return false
}

func IsStringPalindrome(n string) bool {
	if n == StringReverse(n) {
		return true
	}
	return false
}

func UInt64Exp(a, b uint64) uint64 {
	if b == 0 {
		return 1
	}
	if b == 1 {
		return a
	}
	if b%2 == 0 {
		temp := UInt64Exp(a, b/2)
		return temp * temp
	}
	return a * UInt64Exp(a, b-1)
}

func IntExp(a int64, b int64) int64 {
	if a > 0 {
		return int64(UInt64Exp(uint64(a), uint64(b)))
	}
	if a < 0 && b%2 == 0 {
		return IntExp(-1*a, b)
	}

	return -1 * IntExp(-1*a, b)

}

func NumberDigits(n int64) int {
	return (len(strconv.FormatInt(n, 10)))
}
