package main

import (
	"advent-of-code/cmd/utils"
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type packetPair struct {
	left  any
	right any
}

func loadInputPart1(filename string) (res []packetPair) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	i := 0

	var pair packetPair
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		switch i % 3 {
		case 0:
			utils.CheckErr(json.Unmarshal(lineBytes, &pair.left))
		case 1:
			utils.CheckErr(json.Unmarshal(lineBytes, &pair.right))
			res = append(res, pair)
			pair = packetPair{}
		}
		i++
	}
	utils.CheckErr(scanner.Err())
	return res
}

func part1(filename string) int {
	packetPairs := loadInputPart1(filename)
	var idxValid []int
	for i, packetPair := range packetPairs {
		valid, stop := compare(packetPair.left, packetPair.right)
		if !stop {
			log.Fatalln("stop is false!")
		}
		if valid {
			idxValid = append(idxValid, i+1)
		}
	}
	result := 0
	for _, val := range idxValid {
		result += val
	}
	return result
}

func compare(left, right interface{}) (valid bool, stop bool) {
	leftNum, isLeftNum := left.(float64)
	rightNum, isRightNum := right.(float64)

	if isLeftNum && isRightNum {
		return compareNumbers(leftNum, rightNum)
	}

	leftList, isLeftList := left.([]interface{})
	rightList, isRightList := right.([]interface{})
	if isLeftList && isRightList {
		return compareLists(leftList, rightList)
	}

	if isLeftList && isRightNum {
		return compare(leftList, []interface{}{rightNum})

	}
	if isLeftNum && isRightList {
		return compare([]interface{}{leftNum}, rightList)
	}
	log.Fatalln("got unexpected typed", left, right)
	return false, false
}

func compareNumbers(left, right float64) (valid, stop bool) {
	if left < right {
		return true, true
	}
	if left > right {
		return false, true
	}
	return false, false
}

func compareLists(left, right []interface{}) (valid, stop bool) {
	for i := 0; i < len(left); i++ {
		rightOOB := i+1 > len(right)
		if rightOOB {
			// right ran out - INVALID!
			return false, true
		}
		valid, stop = compare(left[i], right[i])
		if stop {
			return valid, stop
		}
	}

	if len(left) == len(right) {
		// no conclusion
		return false, false
	}

	// left has more items - Valid
	return true, true
}

type packet interface{}

func loadInputPart2(filename string) (packets []packet) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		lineBytes := scanner.Bytes()
		if i%3 != 2 {
			var p packet
			utils.CheckErr(json.Unmarshal(lineBytes, &p))
			packets = append(packets, p)
		}
		i++
	}
	utils.CheckErr(scanner.Err())
	return packets
}

type packets []packet

func (x packets) Len() int { return len(x) }
func (x packets) Less(i, j int) bool {
	valid, stop := compare(x[i], x[j])
	if !stop {
		log.Fatalln("stop should be true")
	}
	return valid
}
func (x packets) Swap(i, j int) { x[i], x[j] = x[j], x[i] }

func locatePacketIdx(p packet, packets []packet) int {
	for idx, item := range packets {
		validLess, _ := compare(item, p)
		validMore, _ := compare(p, item)
		if !validLess && !validMore {
			return idx
		}
	}
	log.Fatalln("couldnt find packet")
	return 0
}

func part2(filename string) int {
	input := loadInputPart2(filename)
	var dividerA, dividerB interface{}
	dividerA = []interface{}{[]interface{}{float64(2)}}
	dividerB = []interface{}{[]interface{}{float64(6)}}
	input = append(input, dividerA)
	input = append(input, dividerB)
	sort.Sort(packets(input))
	locA := locatePacketIdx(dividerA, input) + 1
	locB := locatePacketIdx(dividerB, input) + 1
	return locA * locB
}
