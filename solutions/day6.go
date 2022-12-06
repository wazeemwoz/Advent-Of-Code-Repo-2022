package solutions

import (
	"github.com/wazeemwoz/advent2022/file"
)

func Solution6(size int) func(string) int {
	return func(filepath string) int {
		fileStream := file.NewStream(filepath)

		counts := make(map[string]int)
		index := 0
		text := ""
		runningSum := 0

		fileStream.ForEachChar(func(entry string) bool {
			if len(text) >= size {
				runningSum -= counts[text[index:index+1]]
				counts[text[index:index+1]]--
				text += entry
				counts[entry]++
				runningSum += counts[entry]
				index++
				return runningSum != size
			} else {
				text += entry
				counts[entry]++
				runningSum += counts[entry]
			}
			return true
		})
		return len(text)
	}
}
