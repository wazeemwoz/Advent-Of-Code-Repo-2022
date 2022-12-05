package solutions

import (
	"regexp"
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	"github.com/wazeemwoz/advent2022/utils"
)

func Solution5(filepath string) string {
	fileStream := file.NewStream(filepath)
	chartLines := fileStream.NextUntil(regexp.MustCompile(`^\s1`))
	stacks := make([]string, 0)

	for i := 0; i < len(chartLines[len(chartLines)-1]); i++ {
		stack := ""
		for j := len(chartLines) - 1; j >= 0; j-- {
			if string(chartLines[j][i]) == " " {
				break
			}

			stack = stack + string(chartLines[j][i])
		}
		if len(stack) > 0 {
			stacks = append(stacks, stack)
		}
	}

	fileStream.ForEach(func(line string) {
		if len(line) == 0 {
			return
		}

		change := strings.Split(line, " ")

		move := utils.ToInt(change[1])
		from := utils.ToInt(change[3]) - 1
		to := utils.ToInt(change[5]) - 1

		for i := 0; i < move; i++ {
			stacks[to] = stacks[to] + stacks[from][(len(stacks[from]))-1:]
			stacks[from] = stacks[from][0 : (len(stacks[from]))-1]
		}
	})

	answer := ""
	for _, v := range stacks {
		answer += string(v[len(v)-1])
	}

	return answer
}
