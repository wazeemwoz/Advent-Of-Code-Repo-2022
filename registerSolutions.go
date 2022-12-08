package main

import (
	"strconv"

	"github.com/wazeemwoz/advent2022/solutions"
)

var Solutions = registerSolutions()

type Answer interface {
	string | int
}

func asStr(f func(string) int) func(string) string {
	return func(s string) string {
		return strconv.Itoa(f(s))
	}
}

func registerSolutions() map[string](func(string) string) {
	registered := make(map[string](func(string) string))

	registered["1"] = asStr(solutions.Solution1(1))
	registered["1.1"] = asStr(solutions.Solution1(3))

	registered["2"] = asStr(solutions.Solution2(solutions.Part1Strategy()))
	registered["2.1"] = asStr(solutions.Solution2(solutions.Part2Strategy()))

	registered["3"] = asStr(solutions.Solution3)
	registered["3.1"] = asStr(solutions.Solution3_1)

	registered["4"] = asStr(solutions.Solution4(solutions.CountContained))
	registered["4.1"] = asStr(solutions.Solution4(solutions.CountOverlapping))

	registered["5"] = solutions.Solution5(solutions.MoveOneByOne)
	registered["5.1"] = solutions.Solution5(solutions.MoveTogether)

	registered["6"] = asStr(solutions.Solution6(4))
	registered["6.1"] = asStr(solutions.Solution6(14))

	registered["7"] = asStr(solutions.Solution7(solutions.FilteredSize))
	registered["7.1"] = asStr(solutions.Solution7(solutions.FindUnder))

	registered["8"] = asStr(solutions.Solution8)
	registered["8.1"] = asStr(solutions.Solution8_1)

	return registered
}
