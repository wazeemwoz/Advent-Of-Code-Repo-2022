package solutions

import (
	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/types"
	"github.com/wazeemwoz/advent2022/utils"
)

func Solution8(filepath string) int {
	grid := fileToGrid(filepath)
	visibileMap := make(map[Coord]bool)
	visibility := 0
	loopGrid(grid, func(current Coord, stoppedAt Coord) {
		if stoppedAt.X < 0 || stoppedAt.X == len(grid[0]) || stoppedAt.Y < 0 || stoppedAt.Y == len(grid) {
			if !visibileMap[current] {
				visibility++
				visibileMap[current] = true
			}
		}
	})
	return visibility
}

func Solution8_1(filepath string) int {
	grid := fileToGrid(filepath)
	visibileMap := make([][]int, len(grid))
	for y, row := range grid {
		visibileMap[y] = make([]int, len(row))
		for x, _ := range row {
			visibileMap[y][x] = 1
		}
	}
	visibility := 0
	loopGrid(grid, func(current Coord, stoppedAt Coord) {
		stoppedAt.X = utils.Max(stoppedAt.X, 0)
		stoppedAt.X = utils.Min(stoppedAt.X, len(grid[0])-1)
		stoppedAt.Y = utils.Max(stoppedAt.Y, 0)
		stoppedAt.Y = utils.Min(stoppedAt.Y, len(grid)-1)
		visibileMap[current.Y][current.X] *= (utils.Abs(current.X-stoppedAt.X) + utils.Abs(current.Y-stoppedAt.Y))
		visibility = utils.Max(visibileMap[current.Y][current.X], visibility)
	})
	return visibility
}

func loopGrid(grid [][]int, fnStopper func(current Coord, next Coord)) {
	expander := func(x int, y int, direction Coord) {
		next := Coord{direction.X + x, direction.Y + y}
		for !(next.X < 0 || next.X >= len(grid[0]) || next.Y < 0 || next.Y >= len(grid)) {
			if grid[y][x] <= grid[next.Y][next.X] {
				break
			}
			next = Coord{direction.X + next.X, direction.Y + next.Y}
		}
		fnStopper(Coord{x, y}, next)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			expander(x, y, Coord{1, 0})
			expander(x, y, Coord{-1, 0})
			expander(x, y, Coord{0, 1})
			expander(x, y, Coord{0, -1})
		}
	}
}

func fileToGrid(filepath string) [][]int {
	fileStream := file.NewStream(filepath)
	treeGrid := make([][]int, 0)

	fileStream.ForEach(func(line string) {
		treeLine := make([]int, len(line))
		for i, ch := range line {
			treeLine[i] = utils.ToInt(string(ch))
		}
		treeGrid = append(treeGrid, treeLine)
	})

	return treeGrid
}
