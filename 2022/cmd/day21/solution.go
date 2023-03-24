package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type instruction struct {
	number       *int
	firstMonkey  string
	secondMonkey string
	operation    string
}

func loadInput(filename string) (d map[string]instruction) {
	d = make(map[string]instruction)
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	r := regexp.MustCompile("[a-z]{4}")
	rOperation := regexp.MustCompile("[+/*-]")
	rNumber := regexp.MustCompile("\\d+")
	for scanner.Scan() {
		line := scanner.Text()
		matches := r.FindAllString(line, -1)
		if len(matches) == 1 {
			numberStr := rNumber.FindString(line)
			number := utils.StrToInt(numberStr)
			d[matches[0]] = instruction{
				number: &number,
			}
		} else if len(matches) == 3 {
			op := rOperation.FindString(line)
			if op == "" {
				panic("no operation found")
			}
			inst := instruction{
				firstMonkey:  matches[1],
				secondMonkey: matches[2],
				operation:    op,
			}
			d[matches[0]] = inst
		} else {
			panic("Invalid input: " + line + " " + matches[0])
		}
	}
	utils.CheckErr(scanner.Err())
	return d
}

func part1(filename string) int {
	d := loadInput(filename)
	return evaluate("root", d)
}

func evaluate(monkeyName string, d map[string]instruction) int {
	monkey, ok := d[monkeyName]
	if !ok {
		panic("monkey not found")
	}
	if monkey.number != nil {
		return *monkey.number
	}
	firstMonkey := evaluate(monkey.firstMonkey, d)
	secondMonkey := evaluate(monkey.secondMonkey, d)
	switch monkey.operation {
	case "+":
		return firstMonkey + secondMonkey
	case "-":
		return firstMonkey - secondMonkey
	case "*":
		return firstMonkey * secondMonkey
	case "/":
		return firstMonkey / secondMonkey
	}
	panic("Invalid operation")
}

func part2(filename string) int {
	input := loadInput(filename)
	f := func(humn int) int {
		input["humn"] = instruction{number: &humn}
		rootInstruction := input["root"]
		lhs := evaluate(rootInstruction.firstMonkey, input)
		rhs := evaluate(rootInstruction.secondMonkey, input)
		return lhs - rhs
	}
	return binarySearch(f, 1, 10e12, 1)
}

func binarySearch(f func(int) int, a int, b int, eps int) int {
	fA := f(a)
	fB := f(b)

	if !isSignChange(fA, fB) {
		panic(fmt.Sprintf("No sign change. Poor initial values. f(%d)=%d, f(%d)=%d", a, fA, b, fB))
	}

	for !isConverged(a, b, eps) {
		mid := (a + b) / 2
		fMid := f(mid)
		if isSignChange(fA, fMid) {
			b = mid
			fB = fMid
		} else {
			a = mid
			fA = fMid
		}
	}
	if fA == 0 {
		return a
	} else if fB == 0 {
		return b
	} else {
		panic(fmt.Sprintf("No zero found f(%d):%d f(%d):%d", a, fA, b, fB))
	}
}

func isSignChange(a, b int) bool {
	return (a <= 0 && b >= 0) || (a >= 0 && b <= 0)
}

func isConverged(a, b, eps int) bool {
	return abs(a-b) <= eps
}

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}
