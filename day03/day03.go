package main

import (
	"adventOfGo2017/utils"
	"fmt"
)

const (
	RIGHT = 0
	UP    = 1
	LEFT  = 2
	DOWN  = 3
)

var next = [4]int{UP, LEFT, DOWN, RIGHT}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func p1(input int) int {
	directionList := [4]int{0, 0, 0, 0}
	if input > 1 {
		counter1 := 0
		counter2 := 1
		currentDirection := RIGHT
		for n := 2; n <= input; n++ {
			counter1++
			directionList[currentDirection]++
			if counter1 == counter2 {
				counter1 = 0
				currentDirection = next[currentDirection]
				if currentDirection == RIGHT || currentDirection == LEFT {
					counter2 += 1
				}
			}
		}
	}
	return abs(directionList[RIGHT]-directionList[LEFT]) + abs(directionList[UP]-directionList[DOWN])
}

type pair struct {
	x, y int
}

func p2(input int) int {
	var dirVal = [4]pair{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	counter1, counter2, currentDirection, position := 0, 1, RIGHT, pair{0, 0}
	grid := make(map[pair]int)
	grid[pair{0, 0}] = 1

	for grid[position] < input {
		counter1++
		position.x += dirVal[currentDirection].x
		position.y += dirVal[currentDirection].y
		grid[position] = gridValue(position, grid)
		if counter1 == counter2 {
			counter1 = 0
			currentDirection = next[currentDirection]
			if currentDirection == RIGHT || currentDirection == LEFT {
				counter2 += 1
			}
		}
	}
	return grid[position]
}

func gridValue(position pair, grid map[pair]int) int {
	sum := 0
	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if !(i == 0 && j == 0) {
				sum += grid[pair{position.x + i, position.y + j}]
			}
		}
	}
	return sum
}

func main() {
	input := utils.Atoi(utils.GetInput(3))
	fmt.Printf("Part 1: %v\n", p1(input))
	fmt.Printf("Part 2: %v\n", p2(input))
}
