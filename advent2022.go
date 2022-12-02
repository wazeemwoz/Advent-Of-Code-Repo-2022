package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	solution := os.Args[1]
	filePath := os.Args[2]

	solutions := make(map[string](func(string) int))
	solutions["1"] = solution1(1)
	solutions["1.1"] = solution1(3)

	fmt.Printf("Running solution %s with file %s \n", solution, filePath)

	answer := solutions[solution](filePath)

	fmt.Printf("Solution is: %d \n", answer)
}

func withFileDo(pathToFile string, consumer func(string)) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		consumer(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func solution1(topK int) func(string) int {
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

		withFileDo(filepath, func(entry string) {
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
