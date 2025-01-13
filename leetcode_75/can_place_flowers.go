package leetcode_75

import "fmt"

// https://leetcode.com/problems/can-place-flowers/description/?envType=study-plan-v2&envId=leetcode-75
func CanPlaceFlowers() {
	arr := []int{1, 0, 0, 0, 1}
	n := 1
	fmt.Printf("canPlaceFlowers(): %v\n", canPlaceFlowers(arr, n))
}

// Time Complexity: O(n)
// Space Complexity: O(1)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && (i == 0 || flowerbed[i-1] == 0) && (i == len(flowerbed)-1 || flowerbed[i+1] == 0) {
			flowerbed[i] = 1
			n--
		}
		if n == 0 {
			break
		}
	}

	if n == 0 {
		return true
	} else {
		return false
	}

}
