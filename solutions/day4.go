package solutions

import (
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	"github.com/wazeemwoz/advent2022/utils"
)

type indexRange struct {
	left, right int
}

func CountContained(unions int, sizeLeft int, sizeRight int) int {
	if unions == utils.Min(sizeLeft, sizeRight) {
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

			union := left.unions(right)

			score += scoringStrategy(union, left.size(), right.size())
		})
		return score
	}
}

func lineToRanges(line string) (indexRange, indexRange) {
	words := strings.Split(line, ",")
	return wordToRange(words[0]), wordToRange(words[1])
}

func wordToRange(word string) indexRange {
	numsAsStrings := strings.Split(word, "-")
	return indexRange{utils.ToInt(numsAsStrings[0]), utils.ToInt(numsAsStrings[1])}
}

func (r indexRange) size() int {
	return r.right - (r.left - 1)
}

func (r indexRange) unions(other indexRange) int {
	span := indexRange{utils.Min(r.left, other.left), utils.Max(r.right, other.right)}

	return utils.Max((r.size()+other.size())-span.size(), 0)
}
