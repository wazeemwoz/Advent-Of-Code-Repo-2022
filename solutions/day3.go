package solutions

import (
	"github.com/wazeemwoz/advent2022/file"
)

func Solution3(filepath string) int {
	score := 0
	file.WithFileDo(filepath, func(entry string) {
		score += findDuplicates(entry[0:len(entry)/2], entry[len(entry)/2:])
	})
	return score
}

func Solution3_1(filepath string) int {
	score := 0
	lineCount := 0
	rucksacks := make([]string, 3)
	file.WithFileDo(filepath, func(entry string) {
		index := lineCount % len(rucksacks)
		rucksacks[index] = entry
		if index == len(rucksacks)-1 {
			score += findDuplicates(rucksacks...)
		}
		lineCount++
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
