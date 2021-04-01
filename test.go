// Package test is a sandbox for testing algorithms
package test

import (
	"fmt"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

func PowersofTwo(num int) string {

	// code goes here

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

	// code goes here

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

	// code goes here

	var b, c big.Int
	var r int64

	n := int64(num)

	retval := c.Div(b.Binomial(n*2, n), c.SetInt64(n+1))

	retval.SetInt64(r)

	return int(r)
}

func FirstReverse(str string) string {

	// code goes here

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
	return 510
}

func cat(num int) int {
	var b, c big.Int

	n := int64(num)
	s := fmt.Sprint(c.Div(b.Binomial(n*2, n), c.SetInt64(n+1)))
	j, _ := strconv.Atoi(s)
	return j

}
