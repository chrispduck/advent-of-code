package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	//fmt.Println(part1("input.txt"))
	//fmt.Println(part2("example_input.txt"))
	//fmt.Println(part2("input.txt"))
}

type snowflake struct {
	l, r, u, d int
	total int
	wall bool
}

func (s snowflake) String() string {
	if s.wall {
		return "#"
	}
	if s.total == 0 {
		return "."
	}
	return strconv.Itoa(s.total)
}
type grid struct {
	grid [][]snowflake
}

func (g grid) String() string {
	s := ""
	for _, row := range g.grid {
		for _, c := range row {
			s += c.String()
		}
		s += "\n"
	}
	return s
}

func (g grid) Iterate() (res grid)	{
	res.grid = copy(g, res)

	for y, row := range g.grid {
		for x, c := range row {
			if c.wall {
				res.grid[y] = append(res.grid[y], c}
			}

		}
	}
	return res
}

func loadInput(filename string) grid {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	g := grid{}
	for scanner.Scan() {
		line := scanner.Text()
		g.grid = append(g.grid, []rune(line))
	}
	utils.CheckErr(scanner.Err())
	return g
}

func part1(filename string) int {
	g := loadInput(filename)
	fmt.Println(g)
	return 0
}

func part2(filename string) int {
	return 0
}
