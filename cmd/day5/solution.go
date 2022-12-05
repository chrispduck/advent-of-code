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
}

func part1(filename string, n int) string {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)

	stacks := make([][]rune, n)
	var instructions [][]int
	stateRead := false
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if (line == "") || (line[1] == '1') {
			// deal with split between state input and instructions
			//fmt.Println("now reading instructions")
			stateRead = true
		} else {
			if !stateRead {
				//fmt.Println("parsing input")

				parseInputState(&stacks, line, n)
				//fmt.Printf("stacks main: %+c\n", stacks)

			} else {
				instructions = append(instructions, parseInstruction(line))
			}
		}
	}
	checkErr(scanner.Err())
	//fmt.Println(instructions)
	execute(instructions, &stacks)
	//fmt.Printf("stacks main: %+c\n", stacks)
	return getTopCrates(&stacks)
}

func parseInputState(stacks *[][]rune, line string, n int) {
	for i := 0; i < n; i++ {
		idx := 4*i + 1
		if len(line) < idx {
			return
		}
		if line[idx] != ' ' {
			//fmt.Printf("input letter: %c", line[idx])
			prepend(&(*stacks)[i], rune(line[idx]))
		}
	}
	//fmt.Printf("%c\n, %c\n, %c\n", (*stacks)[0], (*stacks)[1], (*stacks)[2])
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
	//fmt.Printf("stack pop %c\n", arr)
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	//fmt.Printf("stack pop %c, item %c\n", arr, res)
	return res
}

func push(arr *[]rune, item rune) {
	*arr = append(*arr, item)
}

func prepend(arr *[]rune, item rune) {
	*arr = append([]rune{item}, *arr...)
}

func execute(instructions [][]int, stacks *[][]rune) {
	for _, instruction := range instructions {
		repeats, from, to := instruction[0], instruction[1], instruction[2]
		//fmt.Printf("stack state: %+c\n", stacks)
		//fmt.Println(repeats, from, to)
		for i := 0; i < repeats; i++ {
			item := pop(&(*stacks)[from-1])
			push(&(*stacks)[to-1], item)
		}
	}
}

func getTopCrates(stacks *[][]rune) string {
	res := ""
	for i := 0; i < len(*stacks); i++ {
		res += string((*stacks)[i][len((*stacks)[i])-1])

	}
	return res
}
