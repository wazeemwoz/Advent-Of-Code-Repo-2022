package main

import "github.com/wazeemwoz/advent2022/solutions"

var Solutions = registerSolutions()

func registerSolutions() map[string](func(string) int) {
	registered := make(map[string](func(string) int))
	registered["1"] = solutions.Solution1(1)
	registered["1.1"] = solutions.Solution1(3)

	registered["2"] = solutions.Solution2(solutions.Part1Strategy())
	registered["2.1"] = solutions.Solution2(solutions.Part2Strategy())

	registered["3"] = solutions.Solution3
	registered["3.1"] = solutions.Solution3_1

	return registered
}
