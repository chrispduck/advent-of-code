package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	filename := "example_input.txt"
	//filename := "input.txt"
	part1, _ := solveRunes(filename)
	fmt.Println(part1)
}

func solveRunes(filename string) (int, int) {
	file, err := os.Open(filename)
	CheckError(err)

	score1 := 0
	score2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes := []rune(line)
		score1 += judgeRunes(runes[0], runes[2])

	}

	return score1, score2
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// judge
// shape a by player one
// shape b by player two
func judgeRunes(a rune, b rune) int {
	score := 0
	// score for playing a type
	switch b {
	case 'X':
		score += 1
	case 'Y':
		score += 2
	case 'Z':
		score += 3
	default:
		log.Fatalln("unexpected move", b)
	}

	// score for winning?
	modulo := mod(int(b-a), 3)
	switch modulo {
	case 2:
		//draw
		score += 3
	case 1:
		// lose
		score += 0
	case 0:
		// win
		score += 6
	default:
		log.Fatalln("unrecognised case", modulo)
	}
	return score
}

func mod(a, b int) int {
	return (a%b + b) % b
}
