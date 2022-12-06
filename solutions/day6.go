package solutions

import (
	"fmt"

	"github.com/wazeemwoz/advent2022/file"
)

func Solution6(size int) func(string) int {
	return func(filepath string) int {
		fileStream := file.NewStream(filepath)

		lastSet := make([]string, size)
		index := 0

		fileStream.ForEachChar(func(entry string) bool {
			lastSet[index%len(lastSet)] = entry
			index++
			return !isUnique(lastSet)
		})
		return index
	}
}

func isUnique(set []string) bool {
	tracking := make(map[string]int)

	for _, v := range set {
		if v == "" {
			return false
		}
		tracking[v]++
		if tracking[v] > 1 {
			return false
		}
	}
	fmt.Println(set)
	return true
}
