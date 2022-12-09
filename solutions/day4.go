package solutions

import (
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/types"
	. "github.com/wazeemwoz/advent2022/utils"
)

func CountContained(unions int, sizeLeft int, sizeRight int) int {
	if unions == Min(sizeLeft, sizeRight) {
		return 1
	}
	return 0
}

func CountOverlapping(unions int, sizeLeft int, sizeRight int) int {
	if unions > 0 {
		return 1
	}
	return 0
}

func Solution4(scoringStrategy func(int, int, int) int) func(string) int {
	return func(filepath string) int {
		score := 0
		file.NewStream(filepath).ForEach(func(entry string) {
			left, right := lineToRanges(entry)

			union := left.Unions(right)

			score += scoringStrategy(union, left.Size(), right.Size())
		})
		return score
	}
}

func lineToRanges(line string) (Range, Range) {
	words := strings.Split(line, ",")
	return wordToRange(words[0]), wordToRange(words[1])
}

func wordToRange(word string) Range {
	numsAsStrings := strings.Split(word, "-")
	return Range{ToInt(numsAsStrings[0]), ToInt(numsAsStrings[1])}
}
