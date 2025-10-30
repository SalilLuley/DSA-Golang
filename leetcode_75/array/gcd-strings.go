package array

import "fmt"

// https://leetcode.com/problems/greatest-common-divisor-of-strings/description/?envType=study-plan-v2&envId=leetcode-75

func GcdOfStrings() {
	str1 := "ABABAB"
	str2 := "ABAB"
	fmt.Printf("gcdOfStrings() %v\n", gcdOfStrings(str1, str2))
}

func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 != str2+str1 {
		return ""
	}
	gcdLength := gcd(len(str1), len(str2))
	return str2[:gcdLength]
}

func gcd(a int, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
