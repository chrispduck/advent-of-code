package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type Coordinate3D struct {
	X, Y, Z int
}

func (c Coordinate3D) getNeighbours() (neighbours []Coordinate3D) {
	return []Coordinate3D{
		{
			X: c.X + 1,
			Y: c.Y,
			Z: c.Z,
		},
		{
			X: c.X - 1,
			Y: c.Y,
			Z: c.Z,
		},
		{
			X: c.X,
			Y: c.Y + 1,
			Z: c.Z,
		},
		{
			X: c.X,
			Y: c.Y - 1,
			Z: c.Z,
		},
		{
			X: c.X,
			Y: c.Y,
			Z: c.Z + 1,
		},
		{
			X: c.X,
			Y: c.Y,
			Z: c.Z - 1,
		},
	}
}

func loadInput(filename string) (grid3D map[Coordinate3D]bool) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	grid3D = make(map[Coordinate3D]bool)
	for scanner.Scan() {
		line := scanner.Text()
		rx := regexp.MustCompile(`\d+`)
		s := rx.FindAllString(line, -1)
		grid3D[Coordinate3D{
			X: utils.StrToInt(s[0]),
			Y: utils.StrToInt(s[1]),
			Z: utils.StrToInt(s[2]),
		}] = true
	}
	utils.CheckErr(scanner.Err())
	return grid3D
}

func getSurfaceArea(grid3D map[Coordinate3D]bool) int {
	totalSurfaceArea := 0
	for cube, _ := range grid3D {
		neighbours := cube.getNeighbours()
		for _, neighbour := range neighbours {
			_, exists := grid3D[neighbour]
			if !exists {
				totalSurfaceArea += 1
			}
		}
	}
	return totalSurfaceArea
}

func part1(filename string) int {
	grid3D := loadInput(filename)
	return getSurfaceArea(grid3D)
}

func getMaxCoords(grid3D map[Coordinate3D]bool) (min, max Coordinate3D) {
	max.X, max.Y, max.Z = math.MinInt, math.MinInt, math.MinInt
	min.X, min.Y, min.Z = math.MaxInt, math.MaxInt, math.MaxInt
	for coord, _ := range grid3D {
		if coord.X > max.X {
			max.X = coord.X
		}
		if coord.Y > max.Y {
			max.Y = coord.Y
		}
		if coord.Z > max.Z {
			max.Z = coord.Z
		}
		if coord.X < min.X {
			min.X = coord.X
		}
		if coord.Y < min.Y {
			min.Y = coord.Y
		}
		if coord.Z < min.Z {
			min.Z = coord.Z
		}
	}
	return
}

func withinGrid(coord, min, max Coordinate3D) bool {
	return coord.X <= max.X && coord.Y <= max.Y && coord.Z <= max.Z && coord.X >= min.X && coord.Y >= min.Y && coord.Z >= min.Z
}

func containsLava(coord Coordinate3D, grid3D *map[Coordinate3D]bool) bool {
	_, exists := (*grid3D)[coord]
	return exists
}

func getExternalSurfaceArea(grid3D map[Coordinate3D]bool) int {
	totalSurfaceArea := 0
	min, max := getMaxCoords(grid3D)
	outerMin, outerMax := Coordinate3D{
		X: min.X - 1,
		Y: min.Y - 1,
		Z: min.Z - 1,
	}, Coordinate3D{
		X: max.X + 1,
		Y: max.Y + 1,
		Z: max.Z + 1,
	}
	// start outside out min point
	toExplore := make(map[Coordinate3D]bool) // must be empty space, within the min-1, and max+1
	toExplore[outerMin] = true
	visited := make(map[Coordinate3D]bool)
	for len(toExplore) > 0 {
		nextToExplore := make(map[Coordinate3D]bool)
		for cube, _ := range toExplore {
			neighbours := cube.getNeighbours()
			for _, neighbour := range neighbours {
				if !withinGrid(neighbour, outerMin, outerMax) {
					//fmt.Println("out of grid", neighbour)
					continue
				}
				if _, alreadyVisited := visited[neighbour]; alreadyVisited {
					//fmt.Println("already visited", neighbour)
					continue
				}
				if containsLava(neighbour, &grid3D) {
					totalSurfaceArea += 1
					//fmt.Println("containers lava", neighbour)
				} else {
					visited[neighbour] = true
					nextToExplore[neighbour] = true
				}
			}
		}
		toExplore = nextToExplore
	}
	return totalSurfaceArea
}

func part2(filename string) int {
	grid3D := loadInput(filename)
	return getExternalSurfaceArea(grid3D)
}
