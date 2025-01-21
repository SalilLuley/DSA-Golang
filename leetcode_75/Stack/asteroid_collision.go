package leetcode_75

import "fmt"

func AsteroidCollision() {
	asteroids := []int{5, 10, -5}
	fmt.Printf("asteroidCollision(asteroids): %v\n", asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) []int {
	stack := make([]int, len(asteroids)-1)

	for i, v := range asteroids {
		if v < 0 {

		} else {
			stack = append(stack, v)
		}
	}

	return stack
}
