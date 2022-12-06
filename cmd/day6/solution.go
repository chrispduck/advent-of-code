package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//fmt.Println(part1("bvwbjplbgvbhsrlpgdmjqwftvncz", 4))
	//fmt.Println(part1("nppdvjthqldpwncqszvftbrmjlhg", 4))
	//fmt.Println(part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 4))
	//fmt.Println(part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 4))
	s := loadInput("input.txt")
	fmt.Println(part1(s, 4))

	//fmt.Println(part1("mjqjpqmgbljsphdztnvjfqwrcgsmlb", 14))
	//fmt.Println(part1("bvwbjplbgvbhsrlpgdmjqwftvncz", 14))
	//fmt.Println(part1("nppdvjthqldpwncqszvftbrmjlhg", 14))
	//fmt.Println(part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 14))
	//fmt.Println(part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 14))
	fmt.Println(part1(s, 14))
}

func part1(s string, tokenLength int) int {
	d := map[rune]int{}

	// initialise d as a counter of letters in sliding window
	for i := 0; i < tokenLength; i++ {
		d[rune(s[i])] += 1
	}
	if dictValuesLtTwo(d) {
		return tokenLength
	}

	// search
	for i := tokenLength; i <= len(s); i++ {
		// swap out letter
		d[rune(s[i-tokenLength])] -= 1
		d[rune(s[i])] += 1
		if dictValuesLtTwo(d) {
			return i + 1
		}
	}
	log.Fatalln("did not find token")
	return -1
}

func dictValuesLtTwo(d map[rune]int) bool {
	for _, val := range d {
		if val > 1 {
			return false
		}
	}
	return true
}

func loadInput(filename string) string {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalln(err)
	}
	reader := bufio.NewScanner(f)
	for reader.Scan() {
		return reader.Text()
	}
	log.Fatalln("couldn't read")
	return ""
}
