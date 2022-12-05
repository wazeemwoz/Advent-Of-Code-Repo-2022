package file

import (
	"bufio"
	"log"
	"os"
	"regexp"
)

type Stream struct {
	file    *os.File
	scanner *bufio.Scanner
}

func NewStream(pathToFile string) Stream {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}

	return Stream{file, bufio.NewScanner(file)}
}

func (stream Stream) NextUntil(regex *regexp.Regexp) []string {
	lines := make([]string, 0)
	for stream.scanner.Scan() {
		line := stream.scanner.Text()
		lines = append(lines, line)
		if regex.MatchString(line) {
			break
		}
	}
	return lines
}

func (stream Stream) ForEach(consumer func(string)) {
	defer stream.file.Close()

	scanner := stream.scanner

	for scanner.Scan() {
		consumer(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func (stream Stream) ForGroup(groupSize int, consumer func([]string)) {
	lineCount := 0
	lines := make([]string, groupSize)
	stream.ForEach(func(entry string) {
		index := lineCount % len(lines)
		lines[index] = entry
		if index == len(lines)-1 {
			consumer(lines)
		}
		lineCount++
	})
}
