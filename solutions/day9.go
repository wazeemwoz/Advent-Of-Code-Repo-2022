package solutions

import (
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
	{X: 2, Y: 2}: {X: -1, Y: -1}, {X: -2, Y: -2}: {X: 1, Y: 1}, {X: -2, Y: 2}: {X: 1, Y: -1}, {X: 2, Y: -2}: {X: -1, Y: 1}, //Diagonals
}

func Solution9(ropeSize int) func(string) int {
	return func(filepath string) int {
		fileStream := file.NewStream(filepath)
		visited := make(map[Coord]int)
		rope := make([]Coord, ropeSize)
		for i := 0; i < ropeSize; i++ {
			rope[i] = Coord{X: 0, Y: 0}
		}

		visited[rope[len(rope)-1]]++

		directions := map[string]Coord{"U": {X: 0, Y: -1}, "D": {X: 0, Y: 1}, "L": {X: -1, Y: 0}, "R": {X: 1, Y: 0}}

		move := func(direction Coord, steps int) {
			for i := 0; i < steps; i++ {
				rope[0] = Coord{X: rope[0].X + direction.X, Y: rope[0].Y + direction.Y}
				prev := rope[0]
				for j := 1; j < len(rope); j++ {
					curr := rope[j]
					prevStep := Coord{X: curr.X - prev.X, Y: curr.Y - prev.Y}
					if step, ok := sol8_supportedSteps[prevStep]; ok {
						rope[j] = Coord{X: curr.X + step.X, Y: curr.Y + step.Y}
						if j == len(rope)-1 {
							visited[rope[j]]++
						}
					}
					prev = rope[j]
				}

			}
		}

		fileStream.ForEach(func(line string) {
			action := strings.Split(line, " ")
			move(directions[action[0]], ToInt(action[1]))
		})

		return len(visited)
	}
}
