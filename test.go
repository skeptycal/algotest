// Package algotest is a sandbox for testing algorithms
package algotest

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

var mathPow = math.Pow

func pow(x, y int) int {
	if x == 0 {
		return 0
	}
	if y == 0 || x == 1 {
		return 1
	}
	if y == 1 {
		return x
	}
	retval := x
	for i := 1; i < y; i++ {
		retval *= x
	}
	return retval
}

func TenToThe(y int) int {
	return pow(10, y)
}

func TwoToThe(y int) int {
	return pow(2, y)
}

func PowersofTwo(num int) string {
	for {
		if num%2 != 0 {
			return "false"
		}
		num /= 2
		if num == 2 {
			return "true"
		}
		if num < 2 {
			return "false"
		}
	}
}

func QuestionsMarks(str string) string {
	seenAny := false
	s := []byte{}

	// get rid of all of the junk ... only concerned with digits and ?'s
	// using bytes ... cuz
	for _, r := range str {
		if unicode.IsDigit(r) || r == '?' {
			s = append(s, byte(r))
		}
	}

	var first byte = 0
	c := 0
	seenFirst := false

	for _, r := range s {
		if r != '?' {
			if seenFirst { // check 2nd number

				// ASCII for '0' is 48
				// 48 + 48 + 10 (bytes add up to 10)
				if first+r == 106 { // sum == 10

					// ... but not 3 ?'s
					if c != 3 {
						return "false"
					}
					seenAny = true // seen sum == 10
				}

				first = r // 2nd number is now the first in the next search

			} else {
				first = r
				seenFirst = true // seen first number (loop flag)
			}
			c = 0 // reset counter after checking a pair of numbers
		} else {
			c++
		}
	}

	if !seenAny {
		return "false"
	}

	return "true"
}

func BracketCombinations(num int) int {

	// if num < 2 {
	// 	return 1
	// }
	return Cat(num)
}

func FirstReverse(str string) string {
	l := len(str) - 1
	sb := strings.Builder{}

	for i := 0; i < l+1; i++ {
		sb.WriteByte(str[l-i])
	}

	fmt.Println(sb.String())

	return sb.String()
}

func isAlpha(str string) bool {
	for i := range str {
		if str[i] < 'A' || str[i] > 'z' {
			return false
		} else if str[i] > 'Z' && str[i] < 'a' {
			return false
		}
	}
	return true
}

func readline() int {
	// test function for CoderByte
	return 510
}

func catNoPreCalc(num int) int {
	var b, c big.Int

	n := int64(num)
	s := fmt.Sprint(c.Div(b.Binomial(n*2, n), c.SetInt64(n+1)))
	j, _ := strconv.Atoi(s)
	return j
}

// Cat calculates the num'th Catalan number
func Cat(num int) int {
	if num < 15 {
		precalc := map[int]int{
			0:  1,
			1:  1,
			2:  2,
			3:  5,
			4:  14,
			5:  42,
			6:  132,
			7:  429,
			8:  1430,
			9:  4862,
			10: 16796,
			11: 58786,
			12: 208012,
			13: 742900,
			14: 2674440,
		}
		return precalc[num]
	}

	// return catNoPreCalc(num)

	var b, c big.Int

	n := int64(num)
	s := fmt.Sprint(c.Div(b.Binomial(n*2, n), c.SetInt64(n+1)))
	j, _ := strconv.Atoi(s)
	return j

}
