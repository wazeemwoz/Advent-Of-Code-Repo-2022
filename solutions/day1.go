package solutions

import (
	"strconv"

	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/utils"
)

func Solution1(topK int) func(string) int {

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

		file.NewStream(filepath).ForEach(func(entry string) {
			if len(entry) > 0 {
				cals, _ := strconv.Atoi(entry)
				runningCount += cals
			} else {
				UpdateTop(top, runningCount)
				runningCount = 0
			}
		})
		UpdateTop(top, runningCount)

		return sum(top)
	}
}
