package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {

	fmt.Println(part1("example_input.txt", 3))
	fmt.Println(part1("input.txt", 9))

	fmt.Println(part2("example_input.txt", 3))
	fmt.Println(part2("input.txt", 9))
}

func part1(filename string, n int) string {
	stacks, instructions := readInput(filename, n)
	executeA(instructions, &stacks)
	return getTopCrates(&stacks)
}

func part2(filename string, n int) string {
	stacks, instructions := readInput(filename, n)
	executeB(instructions, &stacks)
	return getTopCrates(&stacks)
}

func readInput(filename string, n int) (stacks [][]rune, instructions [][]int) {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)

	stacks = make([][]rune, n)
	stateRead := false
	for scanner.Scan() {
		line := scanner.Text()
		if (line == "") || (line[1] == '1') {
			// deal with split between state input and instructions
			stateRead = true
		} else {
			if !stateRead {
				parseInputState(&stacks, line, n)
			} else {
				instructions = append(instructions, parseInstruction(line))
			}
		}
	}
	checkErr(scanner.Err())
	return stacks, instructions
}

func parseInputState(stacks *[][]rune, line string, n int) {
	for i := 0; i < n; i++ {
		idx := 4*i + 1
		if len(line) < idx {
			return
		}
		if line[idx] != ' ' {
			prepend(&(*stacks)[i], rune(line[idx]))
		}
	}
}

func parseInstruction(line string) []int {
	rx := regexp.MustCompile(`\d+`)
	s := rx.FindAllString(line, -1)
	return []int{strToInt(s[0]), strToInt(s[1]), strToInt(s[2])}
}

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	checkErr(err)
	return i
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func pop(arr *[]rune) rune {
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	return res
}

func push(arr *[]rune, item rune) {
	*arr = append(*arr, item)
}

func prepend(arr *[]rune, item rune) {
	*arr = append([]rune{item}, *arr...)
}

func executeA(instructions [][]int, stacks *[][]rune) {
	for _, instruction := range instructions {
		repeats, from, to := instruction[0], instruction[1], instruction[2]
		for i := 0; i < repeats; i++ {
			item := pop(&(*stacks)[from-1])
			push(&(*stacks)[to-1], item)
		}
	}
}

func executeB(instructions [][]int, stacks *[][]rune) {
	for _, instruction := range instructions {
		repeats, from, to := instruction[0], instruction[1], instruction[2]
		_ = to
		size := len((*stacks)[from-1])
		items := (*stacks)[from-1][size-repeats : size]
		// remove them from the previous
		(*stacks)[from-1] = (*stacks)[from-1][:size-repeats]
		// add them to other
		(*stacks)[to-1] = append((*stacks)[to-1], items...)
	}
}

func getTopCrates(stacks *[][]rune) string {
	res := ""
	for i := 0; i < len(*stacks); i++ {
		res += string((*stacks)[i][len((*stacks)[i])-1])

	}
	return res
}
