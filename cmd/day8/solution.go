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
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))

}

func part1(filename string) int {
	grid := loadInput(filename)
	total := 0
	yLen := len(grid)
	xLen := len(grid[0])

	// interior trees
	for y := 1; y < yLen-1; y++ {
		for x := 1; x < xLen-1; x++ {
			if visibleInRowCol(&grid, y, x) {
				total += 1
			}
		}
	}
	exteriorTotal := 2*yLen + 2*xLen - 4
	total += exteriorTotal
	return total
}

func part2(filename string) int {
	grid := loadInput(filename)
	yLen := len(grid)
	xLen := len(grid[0])

	topScore := 0
	// all trees, including edges
	for y := 0; y < yLen; y++ {
		for x := 0; x < xLen; x++ {
			score := scenicScoreSingleTree(&grid, y, x)
			if score > topScore {
				topScore = score
			}
		}
	}
	return topScore
}

func visibleInRowCol(grid *[][]int, y, x int) bool {
	height := (*grid)[y][x]
	yLen := len(*grid)
	xLen := len((*grid)[0])
	if maxHeightInRange(grid, 0, x-1, y, y) < height {
		return true
	}
	if maxHeightInRange(grid, x+1, xLen-1, y, y) < height {
		return true
	}
	if maxHeightInRange(grid, x, x, 0, y-1) < height {
		return true
	}
	if maxHeightInRange(grid, x, x, y+1, yLen-1) < height {
		return true
	}
	return false
}

func maxHeightInRange(grid *[][]int, xLower, xUpper, yLower, yUpper int) int {
	max := 0
	for j := yLower; j <= yUpper; j++ {
		for i := xLower; i <= xUpper; i++ {
			if (*grid)[j][i] > max {
				max = (*grid)[j][i]
			}
		}
	}
	return max
}

func scenicScoreSingleTree(grid *[][]int, y, x int) int {
	height := (*grid)[y][x]
	yLen := len(*grid)
	xLen := len((*grid)[0])

	// Walk <-X,Y
	s1 := 0
	for i := x - 1; i >= 0; i-- {
		s1 += 1
		if (*grid)[y][i] >= height {
			break
		}
	}
	// Walk X,Y->
	s2 := 0
	for i := x + 1; i <= xLen-1; i++ {
		s2 += 1
		if (*grid)[y][i] >= height {
			break
		}
	}
	// Walk up
	s3 := 0
	for j := y - 1; j >= 0; j-- {
		s3 += 1
		if (*grid)[j][x] >= height {
			break
		}
	}
	// Walk down
	s4 := 0
	for j := y + 1; j <= yLen-1; j++ {
		s4 += 1
		if (*grid)[j][x] >= height {
			break
		}
	}
	return s1 * s2 * s3 * s4
}

func loadInput(filename string) [][]int {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)
	var grid [][]int
	for scanner.Scan() {
		var row []int
		line := scanner.Text()
		for _, r := range line {
			row = append(row, int(r-'0'))
		}
		grid = append(grid, row)
	}
	return grid
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
