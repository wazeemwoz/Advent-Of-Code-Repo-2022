package solutions

import (
	"fmt"
	"strings"

	"github.com/wazeemwoz/advent2022/file"
	. "github.com/wazeemwoz/advent2022/utils"
)

type sMonkey struct {
	label                       int
	items                       []*sItem
	counter                     int
	divisibleBy                 int
	monkeyOnTrue, monkeyOnFalse int
	left, operator, right       string
}

type sItem struct {
	original, current int
}

func Modulo(monkey *sMonkey, others []*sMonkey) {
	divider := 1
	for _, monkey := range others {
		divider *= monkey.divisibleBy
	}
	monkey.items[monkey.counter].current %= divider
}

func Divide(div int) func(*sMonkey, []*sMonkey) {
	return func(monkey *sMonkey, others []*sMonkey) {
		monkey.items[monkey.counter].current /= div
	}
}

func Solution11(rounds int, stressRelief func(*sMonkey, []*sMonkey)) func(string) int {
	turn := func(monkey *sMonkey, others []*sMonkey) {
		for ; monkey.counter < len(monkey.items); monkey.counter++ {

			item := monkey.items[monkey.counter]

			left := ToIntDefault(monkey.left, item.current)
			right := ToIntDefault(monkey.right, item.current)

			switch monkey.operator {
			case "*":
				item.current = left * right
			case "+":
				item.current = left + right
			}

			stressRelief(monkey, others)

			if item.current%monkey.divisibleBy == 0 {
				other := others[monkey.monkeyOnTrue]
				other.items = append(other.items, item)
			} else {
				other := others[monkey.monkeyOnFalse]
				other.items = append(other.items, item)
			}
		}
	}

	return func(filepath string) int {
		fileStream := file.NewStream(filepath)
		monkeys := make([]*sMonkey, 0)
		fileStream.ForGroup(7, func(notes []string) {
			monkey := sMonkey{0, make([]*sItem, 0), 0, 1, 0, 0, "", "", ""}
			for _, strnum := range strings.Split(notes[1][18:], ", ") {
				num := ToInt(strnum)
				monkey.items = append(monkey.items, &sItem{num, num})
			}

			fmt.Sscanf(notes[0], "Monkey %d:", &monkey.label)
			fmt.Sscanf(notes[2], "  Operation: new = %s %s %s", &monkey.left, &monkey.operator, &monkey.right)
			fmt.Sscanf(notes[3], "  Test: divisible by %d", &monkey.divisibleBy)
			fmt.Sscanf(notes[4], "    If true: throw to monkey %d", &monkey.monkeyOnTrue)
			fmt.Sscanf(notes[5], "    If false: throw to monkey %d", &monkey.monkeyOnFalse)

			monkeys = append(monkeys, &monkey)
		})

		for i := 0; i < rounds; i++ {
			for _, monkey := range monkeys {
				turn(monkey, monkeys)
			}
			if i == 0 || i == 19 || (i+1)%1000 == 0 {
				fmt.Printf("== After round %d ==\n", i+1)
				topM(2, monkeys)
				fmt.Printf("--------------------\n")
			}
		}

		result := 1
		for _, val := range topM(2, monkeys) {
			result *= val
		}

		return result
	}
}

func topM(topK int, monkeys []*sMonkey) []int {
	topM := make([]int, topK)
	for i, monkey := range monkeys {
		fmt.Printf("Monkey %d inspected items %d times \n", i, monkey.counter)
		UpdateTop(topM, monkey.counter)
	}
	return topM
}
