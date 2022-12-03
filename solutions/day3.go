package solutions

import (
	"github.com/wazeemwoz/advent2022/file"
)

func Solution3(filepath string) int {
	score := 0
	file.ForEachLine(filepath, func(entry string) {
		score += findDuplicates(entry[0:len(entry)/2], entry[len(entry)/2:])
	})
	return score
}

func Solution3_1(filepath string) int {
	score := 0
	file.ForGroupLines(filepath, 3, func(entries []string) {
		score += findDuplicates(entries...)
	})
	return score
}

func findDuplicates(words ...string) int {
	score := 0
	elements := make([]int, 52)
	for section, word := range words {
		for _, ch := range word {
			pos := 0
			if ch >= rune('a') && ch <= rune('z') {
				pos = int(ch) - int(rune('a'))
			} else {
				pos = (int(ch) - int(rune('A'))) + 26
			}

			if elements[pos] == section {
				elements[pos]++
				if elements[pos] == len(words) {
					score += (pos + 1)
				}
			}
		}
	}
	return score
}
