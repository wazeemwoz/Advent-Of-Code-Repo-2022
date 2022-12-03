package file

import (
	"bufio"
	"log"
	"os"
)

func ForEachLine(pathToFile string, consumer func(string)) {
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

func ForGroupLines(pathToFile string, groupSize int, consumer func([]string)) {
	lineCount := 0
	lines := make([]string, groupSize)
	ForEachLine(pathToFile, func(entry string) {
		index := lineCount % len(lines)
		lines[index] = entry
		if index == len(lines)-1 {
			consumer(lines)
		}
		lineCount++
	})
}
