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
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	//fmt.Println(fullyContains(4, 60, 3, 4))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

func part1(filename string) int {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		rx := regexp.MustCompile(`-|,`)
		arr := rx.Split(line, -1)
		//fmt.Println(arr)
		if fullyContains(strToint(arr[0]), strToint(arr[1]), strToint(arr[2]), strToint(arr[3])) {
			score += 1
		}
	}
	checkErr(scanner.Err())
	return score
}

func part2(filename string) int {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)
	score := 0
	for scanner.Scan() {
		line := scanner.Text()
		rx := regexp.MustCompile(`-|,`)
		arr := rx.Split(line, -1)
		if hasOverlap(strToint(arr[0]), strToint(arr[1]), strToint(arr[2]), strToint(arr[3])) {
			score += 1
		}
	}
	checkErr(scanner.Err())
	return score
}

func strToint(s string) int {
	i, err := strconv.Atoi(s)
	checkErr(err)
	return i
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func fullyContains(a1, a2, b1, b2 int) bool {
	if a1 <= b1 && a2 >= b2 {
		return true
	}
	if b1 <= a1 && b2 >= a2 {
		return true
	}
	return false
}

func hasOverlap(a1, a2, b1, b2 int) bool {
	if a1 < b1 && a2 < b1 {
		return false
	}
	if a1 > b2 && a2 > b2 {
		return false
	}
	return true
}
