package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
	"strings"
)

var KeyCycles = []int{20, 60, 100, 140, 180, 220}

func main() {
	fmt.Println(part1("example_input2.txt"))
	fmt.Println(part1("input.txt"))
	part2("example_input2.txt")
	part2("input.txt")

}

type instruction struct {
	name string
	arg  int
}

func loadInput(filename string) []instruction {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	var instructions []instruction
	for scanner.Scan() {
		line := scanner.Text()
		_ = line
		sArr := strings.Split(line, " ")
		var item instruction
		item.name = sArr[0]
		if len(sArr) > 1 {
			item.arg = utils.StrToInt(sArr[1])
		}
		instructions = append(instructions, item)
	}
	utils.CheckErr(scanner.Err())
	return instructions
}

func computeCycleToX(instructions []instruction) map[int]int {
	cycle := 1
	X := 1
	cycleToX := map[int]int{1: 1}
	for _, inst := range instructions {
		switch inst.name {
		case "noop":
			cycleToX[cycle+1] = X
			cycle += 1
		case "addx":
			cycleToX[cycle+1] = X
			X += inst.arg
			cycleToX[cycle+2] = X
			cycle += 2
		}
	}
	return cycleToX
}
func part1(filename string) int {
	instructions := loadInput(filename)
	cycleToX := computeCycleToX(instructions)
	result := 0
	for _, keyCycle := range KeyCycles {
		result += cycleToX[keyCycle] * keyCycle
	}
	return result
}

func part2(filename string) int {
	instructions := loadInput(filename)
	cycleToX := computeCycleToX(instructions)
	const nPixels = 240
	pixels := [nPixels]string{}
	for i := 0; i < nPixels; i++ {
		clockPixel := i % 40
		X := cycleToX[i+1]
		if (X-1 == clockPixel) || (X == clockPixel) || (X+1 == clockPixel) {
			pixels[i] = "#"
		} else {
			pixels[i] = " "
		}
	}
	printPixels(pixels)
	return 0
}

func printPixels(pixels [240]string) {
	for i, s := range pixels {
		if i%40 == 0 {
			fmt.Printf("\n")
		}
		fmt.Printf("%s", s)
	}
}
