package solutions

import (
	"strings"

	"github.com/wazeemwoz/advent2022/file"
)

func Part1Strategy() map[string](map[string]int) {
	mapPoints := make(map[string](map[string]int))
	mapPoints["X"] = map[string]int{"A": 4, "B": 1, "C": 7}
	mapPoints["Y"] = map[string]int{"A": 8, "B": 5, "C": 2}
	mapPoints["Z"] = map[string]int{"A": 3, "B": 9, "C": 6}

	return mapPoints

}

func Part2Strategy() map[string](map[string]int) {
	mapPoints := make(map[string](map[string]int))
	mapPoints["X"] = map[string]int{"A": 3, "B": 1, "C": 2}
	mapPoints["Y"] = map[string]int{"A": 4, "B": 5, "C": 6}
	mapPoints["Z"] = map[string]int{"A": 8, "B": 9, "C": 7}

	return mapPoints
}

func Solution2(strategy map[string](map[string]int)) func(string) int {
	return func(filepath string) int {
		score := 0

		file.WithFileDo(filepath, func(entry string) {
			moves := strings.Split(entry, " ")

			score += strategy[moves[1]][moves[0]]
		})

		return score
	}
}
