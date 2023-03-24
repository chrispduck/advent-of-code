package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	fmt.Println(part1("example_input.txt", 10))
	fmt.Println(part1("input.txt", 2000000))
	fmt.Println(part2("example_input.txt", 20, 20))
	fmt.Println(part2("input.txt", 4_000_000, 4_000_000))
}

func loadInput(filename string) (sensors []utils.Coordinate, beacons []utils.Coordinate) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		r := regexp.MustCompile("-?\\d+")
		utils.CheckErr(err)
		matches := r.FindAllString(line, -1)
		sensor := utils.Coordinate{
			X: utils.StrToInt(matches[0]),
			Y: utils.StrToInt(matches[1]),
		}
		beacon := utils.Coordinate{
			X: utils.StrToInt(matches[2]),
			Y: utils.StrToInt(matches[3]),
		}
		sensors = append(sensors, sensor)
		beacons = append(beacons, beacon)
	}
	utils.CheckErr(scanner.Err())
	return sensors, beacons
}

func part1(filename string, y int) int {
	sensors, beacons := loadInput(filename)
	invalidPositions := make(map[utils.Coordinate]bool)
	nSensors := len(sensors)
	for s := 0; s < nSensors; s++ {
		distToBeacon := sensors[s].L1Distance(beacons[s])
		delta := distToBeacon - utils.Abs(y-sensors[s].Y)
		for x := sensors[s].X - delta; x <= sensors[s].X+delta; x++ {
			// search along the valid section
			if !(beacons[s].X == x && beacons[s].Y == y) {
				// do not count if it is that beacon
				invalidPositions[utils.Coordinate{X: x, Y: y}] = true
			}
		}
	}
	return len(invalidPositions)
}

func part2(filename string, ymax, xmax int) int {
	sensors, beacons := loadInput(filename)
	for y := 0; y <= ymax; y++ {
	loop:
		for x := 0; x <= xmax; x++ {
			for idx, sensor := range sensors {
				distToBeacon := sensor.L1Distance(beacons[idx])
				distToNewPosition := sensor.L1Distance(utils.Coordinate{X: x, Y: y})
				if distToNewPosition <= distToBeacon {
					// invalid position - how far can we skip
					x += sensor.X - x                         // centre below sensor
					x += distToBeacon - utils.Abs(sensor.Y-y) // skip as far as the L1 norm allows
					continue loop
				}
			}
			// no beacons closer
			return x*4000000 + y
		}
	}
	log.Fatalln("no solution")
	return 0
}
