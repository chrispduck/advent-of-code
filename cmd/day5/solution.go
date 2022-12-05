package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
}

func part1(filename string) string {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)

	var stackA, stackB, stackC []rune
	var instructions [][]int
	stateRead := false
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		if (line == "") || (line[1] == '1') {
			// deal with split between state input and instructions
			//fmt.Println("now reading instructions")
			stateRead = true
		} else {
			if !stateRead {
				if line[1] != ' ' {
					prepend(&stackA, rune(line[1]))
				}
				if len(line) >= 6 && line[5] != ' ' {
					prepend(&stackB, rune(line[5]))
				}
				if len(line) >= 10 && line[9] != ' ' {
					prepend(&stackC, rune(line[9]))
				}
				// prepare stacks
				//fmt.Printf("stackA %c\n", stackA)
				//fmt.Printf("stackB %c\n", stackB)
				//fmt.Printf("stackC %c\n", stackC)
			} else {
				instructions = append(instructions, []int{int(line[5] - '0'), int(line[12] - '0'), int(line[17] - '0')})
			}
		}
	}
	checkErr(scanner.Err())
	stackA, stackB, stackC = execute(instructions, stackA, stackB, stackC)
	return getTopCrates(stackA, stackB, stackC)
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func pop(arr *[]rune) rune {
	//fmt.Printf("stack pop %c\n", arr)
	res := (*arr)[len(*arr)-1]
	*arr = (*arr)[:len(*arr)-1]
	//fmt.Printf("stack pop %c, item %c\n", arr, res)

	return res
}

func push(arr *[]rune, item rune) {
	*arr = append(*arr, item)
}

func prepend(arr *[]rune, item rune) {
	*arr = append([]rune{item}, *arr...)
}

func execute(instructions [][]int, stackA, stackB, stackC []rune) ([]rune, []rune, []rune) {
	stacks := []*[]rune{&stackA, &stackB, &stackC}
	for _, instruction := range instructions {
		repeats, from, to := instruction[0], instruction[1], instruction[2]
		//fmt.Println(repeats, from, to)
		for i := 0; i < repeats; i++ {
			item := pop(stacks[from-1])
			push(stacks[to-1], item)
			//fmt.Printf("stackA %c\n", stackA)
			//fmt.Printf("stackB %c\n", stackB)
			//fmt.Printf("stackC %c\n", stackC)
		}

	}
	return *stacks[0], *stacks[1], *stacks[2]
}

func getTopCrates(stackA, stackB, stackC []rune) string {
	s1 := string(pop(&stackA))
	s2 := string(pop(&stackB))
	s3 := string(pop(&stackC))
	return s1 + s2 + s3
}
