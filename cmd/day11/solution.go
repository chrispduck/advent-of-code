package main

import (
	"fmt"
)

func main() {
	//fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type monkey struct {
	items     []int
	operation func(old int) int
	divisible int
	toTrue    int
	toFalse   int
}

func (m *monkey) throw(to *monkey) {
	to.items = append(to.items, m.items[0])
	m.items = m.items[1:]
}

// Manually create input
func loadExampleInput() []*monkey {
	m0 := monkey{
		[]int{79, 98},
		func(old int) int { return old * 19 },
		23,
		2,
		3,
	}
	m1 := monkey{
		[]int{54, 65, 75, 74},
		func(old int) int { return old + 6 },
		19,
		2,
		0,
	}
	m2 := monkey{
		items:     []int{79, 60, 97},
		operation: func(old int) int { return old * old },
		divisible: 13,
		toTrue:    1,
		toFalse:   3,
	}
	m3 := monkey{
		items:     []int{74},
		operation: func(old int) int { return old + 3 },
		divisible: 17,
		toTrue:    0,
		toFalse:   1,
	}
	return []*monkey{&m0, &m1, &m2, &m3}
}

func loadInput() []*monkey {
	m0 := monkey{
		[]int{57},
		func(old int) int { return old * 13 },
		11,
		3,
		2,
	}
	m1 := monkey{
		[]int{58, 93, 88, 81, 72, 73, 65},
		func(old int) int { return old + 2 },
		7,
		6,
		7,
	}
	m2 := monkey{
		items:     []int{65, 95},
		operation: func(old int) int { return old + 6 },
		divisible: 13,
		toTrue:    3,
		toFalse:   5,
	}
	m3 := monkey{
		items:     []int{58, 80, 81, 83},
		operation: func(old int) int { return old * old },
		divisible: 5,
		toTrue:    4,
		toFalse:   5,
	}

	m4 := monkey{
		[]int{58, 89, 90, 96, 55},
		func(old int) int { return old + 3 },
		3,
		1,
		7,
	}
	m5 := monkey{
		[]int{66, 73, 87, 58, 62, 67},
		func(old int) int { return old * 7 },
		17,
		4,
		1,
	}
	m6 := monkey{
		items:     []int{85, 55, 89},
		operation: func(old int) int { return old + 4 },
		divisible: 2,
		toTrue:    2,
		toFalse:   0,
	}
	m7 := monkey{
		items:     []int{73, 80, 54, 94, 90, 52, 69, 58},
		operation: func(old int) int { return old + 7 },
		divisible: 19,
		toTrue:    6,
		toFalse:   0,
	}
	return []*monkey{&m0, &m1, &m2, &m3, &m4, &m5, &m6, &m7}
}

func part1(filename string) int {
	//monkeys := loadExampleInput()
	monkeys := loadInput()
	nRounds := 20

	monkeyIdxToNThrows := make(map[int]int)
	for i := 0; i < nRounds; i++ {
		for monkeyIdx, m := range monkeys {
			nItems := len(m.items)
			for j := 0; j < nItems; j++ {
				// always front of list
				// divide by 3 and floor aka int division
				m.items[0] = m.operation(m.items[0]) / 3
				if m.items[0]%m.divisible == 0 {
					m.throw(monkeys[m.toTrue])
				} else {
					m.throw(monkeys[m.toFalse])
				}
				monkeyIdxToNThrows[monkeyIdx] += 1
			}
		}
	}
	return computeMonkeyBusiness(monkeyIdxToNThrows)
}

func computeMonkeyBusiness(idToNThrows map[int]int) int {

	max, maxIdx := 0, 0
	for idx, val := range idToNThrows {
		if val > max {
			max = val
			maxIdx = idx
		}
	}
	delete(idToNThrows, maxIdx)

	max2, maxIdx2 := 0, 0
	for idx, val := range idToNThrows {
		if val > max2 {
			max2 = val
			maxIdx2 = idx
		}
	}
	delete(idToNThrows, maxIdx2)

	return max * max2
}

func part2(filename string) int {
	//monkeys := loadExampleInput()
	monkeys := loadInput()
	nRounds := 10000
	monkeyIdxToNThrows := make(map[int]int)

	// compute lcm
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.divisible
	}

	for i := 0; i < nRounds; i++ {
		for monkeyIdx, m := range monkeys {
			nItems := len(m.items)
			for j := 0; j < nItems; j++ {
				// always front of list
				// divide by 3 and floor aka int division
				m.items[0] = m.operation(m.items[0]) % lcm
				if m.items[0]%m.divisible == 0 {
					m.throw(monkeys[m.toTrue])
				} else {
					m.throw(monkeys[m.toFalse])
				}
				monkeyIdxToNThrows[monkeyIdx] += 1
			}
		}

	}
	return computeMonkeyBusiness(monkeyIdxToNThrows)
}
