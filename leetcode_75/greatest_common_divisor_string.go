package leetcode_75

import "fmt"

// Time complexity O(N+Log(Min(Len(Str1),Len(Str2))))
// Space complexity O(1)
func GreatestCommonDivisorOfStrings() {
	str1 := "ABABAB"
	str2 := "AB"
	fmt.Printf("gcdOfStrings(): %v\n", gcdOfStrings(str1, str2))
}
func gcdOfStrings(str1 string, str2 string) string {
	// Check if str1 + str2 equals str2 + str1
	if str1+str2 != str2+str1 {
		return ""
	}
	// Find the GCD of the lengths of str1 and str2
	gcdLength := gcd(len(str1), len(str2))
	// The GCD string is the substring of str1 up to gcdLength
	return str1[:gcdLength]
}

// Helper function to calculate GCD of two integers
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
