package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

func loadInput(filename string) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		// TODO use line
	}
	utils.CheckErr(scanner.Err())
	// TODO return something
}

func part1(filename string) int {
	return 0
}

func part2(filename string) int {
	return 0
}
