package algotest

import (
	"math/rand"
)

func IsPalindrome(s string) bool {
	if len(s) < 3 {
		return false
	}
	return s == Reverse(s)
}

// Reverse returns the reversed string
func Reverse(s string) (result string) {
	for _, v := range s {
		result = string(v) + result
	}
	return
}

// CreatePalindrome creates a random palindrome for testing.
//
// 'sizeLimit' is the max length of the string returned. 'oddPercent' is
// the percent chance of an odd length string. 'padPercent' is the
// percent chance that padding is added, meaning that the palindrome
// is enclosed in an outer string
func CreatePalindrome(padPercent, oddPercent, sizeLimit int) string {
	odd := Flip(oddPercent)
	pad := Flip(padPercent)
	s := RandomString(rand.Intn(sizeLimit / 2))
	if odd {
		s += string(RandomLower())
	}
	s += Reverse(s)
	if pad {
		for i := 0; i < rand.Intn(sizeLimit/3)+2; i++ {
			if Flip(50) {
				s = string(RandomLower()) + s
			} else {
				s += string(RandomLower())
			}
		}
	}

	return s
}

// CreatePalindromeSet creates a set of random palindromes for testing.
// 'Count' is the number of palindromes returned.
//
// 'sizeLimit' is the max length of the string returned. 'oddPercent' is
// the percent chance of an odd length string. 'padPercent' is the
// percent chance that padding is added, meaning that the palindrome
// is enclosed in an outer string
func CreatePalindromeSet(padPercent, oddPercent, sizeLimit, Count int) (retval []string) {
	for i := 0; i < Count; i++ {
		retval = append(retval, CreatePalindrome(padPercent, oddPercent, sizeLimit))
	}
	return
}
