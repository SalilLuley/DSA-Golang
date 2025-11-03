package array

import "fmt"

func KidsWithCandies() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Printf("kidsWithCandies(candies, extraCandies) %v\n", kidsWithCandies(candies, extraCandies))
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	maxCandies := 0
	canCarryExtraCandies := []bool{}
	for i := 0; i < len(candies); i++ {
		maxCandies = max(maxCandies, candies[i])
	}

	for i := 0; i < len(candies); i++ {
		canCarryExtraCandies = append(canCarryExtraCandies, candies[i]+extraCandies >= maxCandies)
	}

	return canCarryExtraCandies
}
