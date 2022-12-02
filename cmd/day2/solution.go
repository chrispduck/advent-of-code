package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	filename := "example_input.txt"
	//filename := "input.txt"
	part1, part2 := solve(filename)
	fmt.Println(part1)
	fmt.Println(part2)
}

func solve(filename string) (int, int) {
	// A Rock B Paper C Scissors
	// X Rock Y Paper Z scissors
	// X = 1
	// Y = 2
	// Z - 3
	// Loss = 0
	// Draw = 3
	// Win = 6

	file, err := os.Open(filename)
	checkError(err)

	score1 := 0
	score2 := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		shapes := strings.Split(line, " ")
		if cap(shapes) != 2 {
			log.Fatalln("unrecognised input ", shapes)
		}
		score1 += judge(shapes[0], shapes[1])
		shape2 := determineMove(shapes[0], shapes[1])
		score2 += judge(shapes[0], shape2)
	}

	return score1, score2
}

func checkError(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// judge
// shape a by player one
// shape b by player two
func judge(a string, b string) int {
	score := 0
	// score for playing a type
	switch b {
	case "X":
		score += 1
	case "Y":
		score += 2
	case "Z":
		score += 3
	default:
		log.Fatalln("unexpected move", b)
	}

	// score for winning
	switch a + b {
	case "AX":
		score += 3
	case "AY":
		score += 6
	case "AZ":
		score += 0
	case "BX":
		score += 0
	case "BY":
		score += 3
	case "BZ":
		score += 6
	case "CX":
		score += 6
	case "CY":
		score += 0
	case "CZ":
		score += 3
	default:
		log.Fatalln("unexpected moves", a, b)
	}
	return score
}

// a in (A, B, C) rock/paper/scissors
// outcome in (X, Y, Z) lose/draw/win
// returns X,Y,Z rock/paper/scissors
func determineMove(a string, outcome string) string {
	switch a + outcome {
	// losing cases
	case "AX":
		return "Z"
	case "BX":
		return "X"
	case "CX":
		return "Y"
	// draw cases:
	case "AY":
		return "X"
	case "BY":
		return "Y"
	case "CY":
		return "Z"
	// win cases
	case "AZ":
		return "Y"
	case "BZ":
		return "Z"
	case "CZ":
		return "X"
	default:
		log.Fatalln("unexpected case", a+outcome)
		return ""
	}
}
