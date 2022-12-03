package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"unicode"
)

func main() {
	fmt.Println(part1("input_example.txt"))
	fmt.Println(part2("input_example.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("input.txt"))
}

func part1(filename string) int {
	f, err := os.Open(filename)
	handleErr(err)
	scanner := bufio.NewScanner(f)
	var commonItems []rune
	for scanner.Scan() {
		rucksackItems := scanner.Text()
		c := len(rucksackItems)
		itemsA, itemsB := rucksackItems[:c/2], rucksackItems[c/2:]
		commonItems = append(commonItems, findCommon(itemsA, itemsB))
		//fmt.Println(commonItems)
	}
	return computeScore(commonItems)
}

func part2(filename string) int {
	f, err := os.Open(filename)
	handleErr(err)
	scanner := bufio.NewScanner(f)
	lines := make([]string, 3)
	var badges []rune
	i := 0
	for scanner.Scan() {
		lines[i%3] = scanner.Text()
		if i%3 == 2 {
			//fmt.Println(lines)
			badges = append(badges, findBadge(lines[0], lines[1], lines[2]))
		}
		i++
	}
	//fmt.Printf("%c\n", badges)
	return computeScore(badges)
}

func findCommon(s1 string, s2 string) rune {
	inS1 := make(map[rune]bool)
	for _, c := range s1 {
		inS1[c] = true
	}
	for _, c := range s2 {
		if inS1[c] == true {
			//fmt.Println(string(c))
			return c
		}
	}
	log.Fatalln("no common letter")
	return -1
}

func findBadge(a, b, c string) rune {
	inA, inB, inC := map[rune]bool{}, map[rune]bool{}, map[rune]bool{}
	for _, r := range a {
		inA[r] = true
	}
	for _, r := range b {
		inB[r] = true
	}
	for _, r := range c {
		inC[r] = true
	}
	for key := range inA {
		if inB[key] && inC[key] {
			return key
		}
	}
	log.Fatalln("did not find common letter")
	return -1
}

func computeScore(commonItems []rune) int {
	score := 0
	for _, item := range commonItems {
		if unicode.IsLower(item) {
			score += int(item) - 'a' + 1
		} else {
			score += int(item) - 'A' + 27
		}
	}
	return score
}

func handleErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
