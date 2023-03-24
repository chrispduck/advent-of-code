package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	filename := "day1/input_test1.txt"
	part1, part2 := solution(filename)
	fmt.Println(part1, part2)
	filename = "day1/input.txt"
	part1, part2 = solution(filename)
	fmt.Println(part1, part2)

}

func solution(fname string) (maxCalories int, caloriesTop3Elves int) {

	// open file
	f, err := os.Open(fname)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer func(f *os.File) {
		err := f.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(f)

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	var elvesCalories []int
	elfCalories := 0

	for scanner.Scan() {
		//fmt.Printf("line: %s\n", scanner.Text())
		lineOfText := scanner.Text()
		if scanner.Text() == "" {
			// add to list and reset counter
			elvesCalories = append(elvesCalories, elfCalories)
			elfCalories = 0
		} else {
			// sum calories for this elf
			calories, err := strconv.Atoi(lineOfText)
			if err != nil {
				log.Fatal(err)
			}
			elfCalories += calories
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	//fmt.Println(elvesCalories)
	maxCalories, idx := findMax(elvesCalories)
	elvesCalories[idx] = 0

	// part 2
	maxCalories2, idx := findMax(elvesCalories)
	elvesCalories[idx] = 0
	maxCalories3, idx := findMax(elvesCalories)
	elvesCalories[idx] = 0
	caloriesTop3Elves = maxCalories + maxCalories2 + maxCalories3
	return maxCalories, caloriesTop3Elves
}

func findMax(s []int) (max int, index int) {
	max = s[0]
	index = 0
	for idx, value := range s {
		if value > max {
			max = value
			index = idx
		}
	}
	return
}
