package solutions

import (
	"github.com/wazeemwoz/advent2022/file"
	"github.com/wazeemwoz/advent2022/utils"
)

type coord struct {
	x int
	y int
}

func Solution8(filepath string) int {
	grid := fileToGrid(filepath)
	visibileMap := make(map[coord]bool)
	visibility := 0
	loopGrid(grid, func(current coord, stoppedAt coord) {
		if stoppedAt.x < 0 || stoppedAt.x == len(grid[0]) || stoppedAt.y < 0 || stoppedAt.y == len(grid) {
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
	loopGrid(grid, func(current coord, stoppedAt coord) {
		stoppedAt.x = utils.Max(stoppedAt.x, 0)
		stoppedAt.x = utils.Min(stoppedAt.x, len(grid[0])-1)
		stoppedAt.y = utils.Max(stoppedAt.y, 0)
		stoppedAt.y = utils.Min(stoppedAt.y, len(grid)-1)
		visibileMap[current.y][current.x] *= (utils.Abs(current.x-stoppedAt.x) + utils.Abs(current.y-stoppedAt.y))
		visibility = utils.Max(visibileMap[current.y][current.x], visibility)
	})
	return visibility
}

func loopGrid(grid [][]int, fnStopper func(current coord, next coord)) {
	expander := func(x int, y int, direction coord) {
		next := coord{direction.x + x, direction.y + y}
		for !(next.x < 0 || next.x >= len(grid[0]) || next.y < 0 || next.y >= len(grid)) {
			if grid[y][x] <= grid[next.y][next.x] {
				break
			}
			next = coord{direction.x + next.x, direction.y + next.y}
		}
		fnStopper(coord{x, y}, next)
	}

	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[y]); x++ {
			expander(x, y, coord{1, 0})
			expander(x, y, coord{-1, 0})
			expander(x, y, coord{0, 1})
			expander(x, y, coord{0, -1})
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
