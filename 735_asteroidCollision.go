package main

import (
	"container/list"
	"math"
)

func asteroidCollision(asteroids []int) []int {

	if len(asteroids) == 0 {
		return []int{}
	}
	stack := list.New()
	stack.PushBack(asteroids[0])

	for i := 1; i < len(asteroids); i++ {

		top := stack.Back()

		if stack.Len() != 0 && top.Value.(int) > 0 && asteroids[i] < 0 {
			//撞击
			helper(stack, asteroids[i])
		} else {
			stack.PushBack(asteroids[i])
		}
	}

	results := make([]int, stack.Len())

	for i := 0; i < len(results); i++ {

		head := stack.Front()

		results[i] = head.Value.(int)
		stack.Remove(head)
	}
	return results
}

func helper(stack *list.List, weight int) {
	abs := func(number int) int {
		return int(math.Abs(float64(number)))
	}

	for stack.Len() != 0 && stack.Back().Value.(int) > 0 && weight < 0 {

		top := stack.Back()

		if abs(top.Value.(int)) == abs(weight) {
			stack.Remove(top)
			return
		} else if abs(top.Value.(int)) < abs(weight) {
			stack.Remove(top)
		} else {
			return
		}
	}
	stack.PushBack(weight)
}
