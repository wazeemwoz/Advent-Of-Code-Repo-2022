package main

import (
	"fmt"
	"os"
)

func main() {
	solution := os.Args[1]
	filePath := os.Args[2]

	fmt.Printf("Running solution %s with file %s \n", solution, filePath)

	answer := Solutions[solution](filePath)

	fmt.Printf("Solution is: %d \n", answer)
}
