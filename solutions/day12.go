package solutions

import (
	"fmt"

	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/types"
	"github.com/wazeemwoz/advent2022/utils"
)

var Adjacent = []Coord{
	{1, 0}, {0, 1}, {-1, 0}, {0, -1},
}

func FromBestStartingPosition(visited [][]int, grid [][]rune, end Coord) int {
	bestA := visited[end.Y][end.X] - 1
	for y, row := range visited {
		for x, e := range row {
			if grid[y][x] == 'a' {
				bestA = utils.Min(bestA, e)
			}
		}
	}
	return bestA
}

func FromStartingPosition(visited [][]int, grid [][]rune, end Coord) int {
	return visited[end.Y][end.X]
}

func Solution12(fnResult func([][]int, [][]rune, Coord) int) func(string) int {
	fileToGrid := func(filepath string) ([][]rune, Coord, Coord, [][]int) {
		fileStream := file.NewStream(filepath)
		grid := make([][]rune, 0)
		visited := make([][]int, 0)
		var start, end Coord
		fileStream.ForEach(func(line string) {
			grid = append(grid, make([]rune, len(line)))
			visited = append(visited, make([]int, len(line)))
			y := len(grid) - 1
			for x, r := range line {
				grid[y][x] = r
				visited[y][x] = int(^uint(0) >> 1)
				switch r {
				case 'S':
					end = Coord{X: x, Y: y}
					grid[y][x] = 'a'
				case 'E':
					start = Coord{X: x, Y: y}
					grid[y][x] = 'z'
					visited[y][x] = 0
				}
			}
		})
		return grid, start, end, visited
	}

	return func(filepath string) int {
		grid, start, end, visited := fileToGrid(filepath)

		queue := make([]Coord, 1)
		queue[0] = start

		for len(queue) > 0 {
			current := queue[0]
			queue = queue[1:]
			for _, dir := range Adjacent {
				next := current.Add(dir)
				if !(next.Y < 0 || next.Y >= len(grid) || next.X < 0 || next.X >= len(grid[next.Y])) {
					elevation := grid[next.Y][next.X] - grid[current.Y][current.X]
					if elevation >= -1 {
						if visited[next.Y][next.X] == int(^uint(0)>>1) {
							queue = append(queue, next)
						}
						visited[next.Y][next.X] = utils.Min(visited[current.Y][current.X]+1, visited[next.Y][next.X])
					}
				}
			}
		}
		print(visited, grid)

		return fnResult(visited, grid, end)
	}
}

func print(visibility [][]int, elevations [][]rune) {
	for y, row := range elevations {
		for x, e := range row {
			if visibility[y][x] == 0 {
				fmt.Printf(".")
			} else {
				fmt.Printf("%s", string(e))
			}
		}
		fmt.Printf("\n")
	}
}
