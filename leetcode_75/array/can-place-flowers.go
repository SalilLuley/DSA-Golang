package array

import "fmt"

// https://leetcode.com/problems/can-place-flowers/description/?envType=study-plan-v2&envId=leetcode-75
func CanPlaceFlowers() {
	flowerbed := []int{0, 0, 0, 0, 0, 1, 0, 0}
	n := 0
	fmt.Printf("canPlaceFlowers() %v\n", canPlaceFlowers(flowerbed, n))
}

func canPlaceFlowers(flowerbed []int, n int) bool {
	counter := 0
	lenFlowerbed := len(flowerbed)
	for i := range flowerbed {
		if flowerbed[i] == 0 {
			empty_left_pot := i == 0 || flowerbed[i-1] == 0

			empty_right_pot := i == lenFlowerbed-1 || flowerbed[i+1] == 0

			if empty_left_pot && empty_right_pot {
				flowerbed[i] = 1
				counter++
			}
		}
	}
	return counter >= n
}
