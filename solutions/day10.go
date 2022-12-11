package solutions

import (
	"fmt"
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/utils"
)

func draw(cycles []int) {
	for cycle, x := range cycles {
		if cycle%40 == 0 {
			fmt.Println()
		}
		if cycle%40 >= x-1 && cycle%40 <= x+1 {
			fmt.Print("#")
		} else {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func Solution10(filepath string) int {
	fileStream := file.NewStream(filepath)
	x_register := make([]int, 0)
	signal := 0
	last_x := 1

	fileStream.ForEach(func(line string) {
		instruction := strings.Split(line, " ")
		if instruction[0] != "noop" {
			x_register = append(x_register, last_x)
			x_register = append(x_register, last_x)
			last_x = last_x + ToInt(instruction[1])
		} else {
			x_register = append(x_register, last_x)
		}
	})
	x_register = append(x_register, last_x)

	for _, v := range []int{20, 60, 100, 140, 180, 220} {
		if v < len(x_register) {
			signal += (v * x_register[v-1])
		}
	}

	draw(x_register)

	return signal
}
