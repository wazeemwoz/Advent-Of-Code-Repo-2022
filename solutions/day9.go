package solutions

import (
	"fmt"
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/types"
	. "github.com/wazeemwoz/advent2022/utils"
)

var sol8_supportedSteps = map[Coord]Coord{
	{X: -2, Y: 0}: {X: 1, Y: 0}, {X: 2, Y: 0}: {X: -1, Y: 0}, {X: 0, Y: -2}: {X: 0, Y: 1}, {X: 0, Y: 2}: {X: 0, Y: -1}, //LEFT, RIGHT, UP, DOWN
	{X: -2, Y: -1}: {X: 1, Y: 1}, {X: -1, Y: -2}: {X: 1, Y: 1}, //UP LEFT
	{X: -2, Y: 1}: {X: 1, Y: -1}, {X: -1, Y: 2}: {X: 1, Y: -1}, //DOWN LEFT
	{X: 2, Y: -1}: {X: -1, Y: 1}, {X: 1, Y: -2}: {X: -1, Y: 1}, //UP RIGHT
	{X: 2, Y: 1}: {X: -1, Y: -1}, {X: 1, Y: 2}: {X: -1, Y: -1}, //DOWN RIGHT
}

func Solution9(filepath string) int {
	fileStream := file.NewStream(filepath)
	visited := make(map[Coord]int)
	head := Coord{X: 0, Y: 0}
	tail := Coord{X: 0, Y: 0}

	visited[tail]++

	directions := map[string]Coord{"U": {X: 0, Y: -1}, "D": {X: 0, Y: 1}, "L": {X: -1, Y: 0}, "R": {X: 1, Y: 0}}

	move := func(direction Coord, steps int) {
		for i := 0; i < steps; i++ {
			head.X = head.X + direction.X
			head.Y = head.Y + direction.Y

			if step, ok := sol8_supportedSteps[Coord{X: tail.X - head.X, Y: tail.Y - head.Y}]; ok {
				tail = Coord{X: tail.X + step.X, Y: tail.Y + step.Y}
				visited[tail]++
				fmt.Printf("Head: (%d,%d), Tail: (%d,%d) => %d \n", head.X, head.Y, tail.X, tail.Y, visited[tail])
			}
		}
	}

	fileStream.ForEach(func(line string) {
		action := strings.Split(line, " ")
		move(directions[action[0]], ToInt(action[1]))
	})

	return len(visited)
}
