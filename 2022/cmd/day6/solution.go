package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const (
	tokenLengthPart1 = 4
	tokenLengthPart2 = 14
)

func main() {
	//fmt.Println(part1("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	//fmt.Println(part1("nppdvjthqldpwncqszvftbrmjlhg"))
	//fmt.Println(part1("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	//fmt.Println(part1("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	s := loadInput("input.txt")
	fmt.Println(part1(s))

	//fmt.Println(part2("mjqjpqmgbljsphdztnvjfqwrcgsmlb"))
	//fmt.Println(part2("bvwbjplbgvbhsrlpgdmjqwftvncz"))
	//fmt.Println(part2("nppdvjthqldpwncqszvftbrmjlhg"))
	//fmt.Println(part2("nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg"))
	//fmt.Println(part2("zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"))
	fmt.Println(part2(s))
}

func part1(s string) int {
	return solution(s, tokenLengthPart1)
}
func part2(s string) int {
	return solution(s, tokenLengthPart2)
}
func solution(s string, tokenLength int) int {
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
