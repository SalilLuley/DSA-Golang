package leetcode_75

import (
	"fmt"
	"math"
)

func AsteroidCollision() {
	// asteroids := []int{5, 10, -5}
	asteroids := []int{10, 2, -5}
	fmt.Printf("asteroidCollision(asteroids): %v\n", asteroidCollision(asteroids))
}

func asteroidCollision(asteroids []int) []int {

	stack := Stack{}

	for i := range asteroids {
		if asteroids[i] > 0 {
			stack.Push(asteroids[i])
		} else {
			fmt.Printf("stack: %v\n", stack)
			result := stack.Pop()
			fmt.Printf("result: %v\n", result)
			fmt.Printf("(result+asteroids[i] == 0): %v\n", (result+asteroids[i] == 0))

			if (result+asteroids[i] == 0) != true {
				max := math.Max(float64(result), float64(asteroids[i]))
				stack.Push(int(max))
			}
		}

	}
	return stack.items
}

func (s *Stack) Push(data int) {
	s.items = append(s.items, data)
}

func (s *Stack) Pop() int {
	if s.IsEmpty() {
		return 0
	}
	item := s.items[len(s.items)-1]
	s.items = s.items[0 : len(s.items)-1]
	return item
}

func (s *Stack) IsEmpty() bool {
	if len(s.items) == 0 {
		return true
	}
	return false
}

type Stack struct {
	items []int
}
