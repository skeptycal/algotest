package algotest

import (
	"fmt"
	"math/rand"
	"strings"
)

func RandomLower() byte {
	return byte(rand.Intn(26) + 97)
}

func RandomString(size int) string {
	sb := strings.Builder{}
	defer sb.Reset()
	for i := 0; i < size; i++ {
		sb.WriteByte(RandomLower())
	}
	return sb.String()
}

func Flip(chance int) bool {
	return rand.Intn(100) > chance
}

func RandomStringSet(n int) (retval []string) {
	for i := 0; i < n; i++ {
		retval = append(retval, RandomString(i))
	}
	return
}

func CreateCatTable(n int) {
	fmt.Println("precalc := map[int]int{")
	for i := 0; i < n+1; i++ {
		fmt.Printf("	%v: %v,\n", i, Cat(i))
	}
	fmt.Println("}")
}

func ShowStrings(set []string) {
	for _, s := range set {
		fmt.Printf("%s - %t\n", s, IsPalindrome(s))
	}
}
