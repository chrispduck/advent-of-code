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
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

func loadInput(filename string) (arr []Item) {
	f, err := os.Open(filename)
	utils.CheckErr(err)
	scanner := bufio.NewScanner(f)
	i := 0
	for scanner.Scan() {
		line := scanner.Text()
		x, err := strconv.Atoi(line)
		utils.CheckErr(err)
		arr = append(arr, Item{StartingIdx: i, Val: x})
		i++
	}
	utils.CheckErr(scanner.Err())
	return arr
}

type Item struct {
	Val         int
	StartingIdx int
}

func part1(filename string) int {
	arr := loadInput(filename)
	//fmt.Printf("starting array:     ")
	//printArray(arr)
	for startingIdx := 0; startingIdx < len(arr); startingIdx++ {
		currentIndex := findIndex(arr, startingIdx)
		arr = mixOnce(arr, currentIndex)
		//fmt.Printf("array after mixing: ")
		//printArray(arr)
	}

	zeroIdx := locateZero(arr)
	idxA := positiveModulo(zeroIdx+1000, len(arr))
	idxB := positiveModulo(zeroIdx+2000, len(arr))
	idxC := positiveModulo(zeroIdx+3000, len(arr))
	fmt.Println(arr[idxA].Val, arr[idxB].Val, arr[idxC].Val)

	return arr[idxA].Val + arr[idxB].Val + arr[idxC].Val
}

func mixOnce(arr []Item, idx int) []Item {
	length := len(arr)
	val := arr[idx].Val
	if val == 0 {
		//fmt.Println("n = 0, nothing to do")
		return arr
	}
	startIdx := idx
	endIdx := positiveModulo(idx+val, length-1)
	if endIdx == 0 && val < 0 {
		//fmt.Println("endIdx == 0, adding length")
		endIdx += length - 1
	}
	//fmt.Println("moving n:", val, "startIdx:", startIdx, "endIdx:", endIdx)
	if endIdx > startIdx {
		// moving right
		lhs := arr[:startIdx]
		rhs := arr[endIdx+1:]
		mid := arr[startIdx+1 : endIdx+1]
		toJoin := [][]Item{lhs, mid, {arr[startIdx]}, rhs}
		return join(toJoin)
	}

	if startIdx > endIdx {
		// moving left
		lhs := arr[:endIdx]
		rhs := arr[startIdx+1:]
		mid := arr[endIdx:startIdx]
		toJoin := [][]Item{lhs, {arr[startIdx]}, mid, rhs}
		return join(toJoin)
	}
	//fmt.Println("startIdx == endIdx", startIdx, endIdx)
	panic("should not happen startIdx == endIdx" + fmt.Sprintf("%v %v ", startIdx, endIdx))
}

func join(arrs [][]Item) []Item {
	//fmt.Printf("joining: ")
	//printArrays(arrs)
	var result []Item
	for _, arr := range arrs {
		for _, x := range arr {
			result = append(result, x)
		}
	}
	return result
}

//func printArray(arr []Item) {
//	for _, x := range arr {
//		fmt.Print(x.Val, " ")
//	}
//	fmt.Println()
//}

//func printArrays(arrs [][]Item) {
//	for _, arr := range arrs {
//		fmt.Printf("[")
//		for _, x := range arr {
//			fmt.Print(x.Val, " ")
//		}
//		fmt.Printf("] ")
//	}
//	fmt.Println()
//}

func findIndex(arr []Item, sIdx int) int {
	for i := range arr {
		if arr[i].StartingIdx == sIdx {
			return i
		}
	}
	panic("not found")
}

func locateZero(arr []Item) int {
	for i, x := range arr {
		if x.Val == 0 {
			return i
		}
	}
	panic("not found")
}

func part2(filename string) int {
	arr := loadInput(filename)
	//fmt.Printf("starting array:     ")
	//printArray(arr)

	// apply decryption key first
	for i := range arr {
		arr[i].Val *= 811589153
	}

	for i := 0; i < 10; i++ {
		// mix 10 times
		for startingIdx := 0; startingIdx < len(arr); startingIdx++ {
			currentIndex := findIndex(arr, startingIdx)
			arr = mixOnce(arr, currentIndex)
			//fmt.Printf("array after mixing: ")
			//printArray(arr)
		}
	}

	zeroIdx := locateZero(arr)
	idxA := positiveModulo(zeroIdx+1000, len(arr))
	idxB := positiveModulo(zeroIdx+2000, len(arr))
	idxC := positiveModulo(zeroIdx+3000, len(arr))
	fmt.Println(arr[idxA].Val, arr[idxB].Val, arr[idxC].Val)

	return arr[idxA].Val + arr[idxB].Val + arr[idxC].Val
}

func positiveModulo(x, n int) int {
	return (x%n + n) % n
}
