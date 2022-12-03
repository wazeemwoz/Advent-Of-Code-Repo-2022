package solutions

import (
	"strconv"

	"github.com/wazeemwoz/advent2022/file"
)

func Solution1(topK int) func(string) int {
	updateTop := func(list []int, val int) {
		newVal := val
		for i, v := range list {
			if newVal > v {
				list[i] = newVal
				newVal = v
			}
		}
	}

	sum := func(list []int) int {
		result := 0
		for _, v := range list {
			result += v
		}
		return result
	}

	return func(filepath string) int {
		top := make([]int, topK)
		runningCount := 0

		file.ForEachLine(filepath, func(entry string) {
			if len(entry) > 0 {
				cals, _ := strconv.Atoi(entry)
				runningCount += cals
			} else {
				updateTop(top, runningCount)
				runningCount = 0
			}
		})
		updateTop(top, runningCount)

		return sum(top)
	}
}
