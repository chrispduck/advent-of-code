package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	part1MaxFileSize = 100_000
	part2MaxSize     = 40_000_000
)

func main() {
	fmt.Println(part1("example_input.txt"))
	fmt.Println(part1("input.txt"))
	fmt.Println(part2("example_input.txt"))
	fmt.Println(part2("input.txt"))
}

type command struct {
	cmd     string
	arg     string
	results [][]string
}

func part1(filename string) int {
	commands := loadInput(filename)
	dirToSize := computeDirSizes(commands)
	return sumAllDirsWithSizeLE(dirToSize, part1MaxFileSize)
}

func part2(filename string) int {
	commands := loadInput(filename)
	dirToSize := computeDirSizes(commands)
	return findSizeOfDirToDelete(dirToSize, part2MaxSize)
}

func loadInput(filename string) []command {
	f, err := os.Open(filename)
	checkErr(err)
	scanner := bufio.NewScanner(f)

	var parsedCommands []command
	c := command{}
	r, err := regexp.Compile(" ")
	checkErr(err)
	for scanner.Scan() {
		line := scanner.Text()
		items := r.Split(line, -1)
		if items[0] == "$" {
			parsedCommands = append(parsedCommands, c)
			c = command{}
			c.cmd = items[1]
			if len(items) == 3 {
				c.arg = items[2]
			}
		} else {
			c.results = append(c.results, items)
		}
	}
	checkErr(scanner.Err())
	parsedCommands = append(parsedCommands, c)
	return parsedCommands[1:]
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func computeDirSizes(cmds []command) map[string]int {
	cwd := ""
	dirToSize := map[string]int{}
	for _, cmd := range cmds {
		switch cmd.cmd {
		case "cd":
			if cmd.arg == "/" {
				cwd = "/"
			} else if cmd.arg == ".." {
				dirs := strings.Split(cwd, "/")
				cwd = strings.Join(dirs[:len(dirs)-2], "/") + "/"
			} else {
				cwd += cmd.arg + "/"
			}
		case "ls":
			for _, result := range cmd.results {
				if fileSize, err := strconv.Atoi(result[0]); err == nil {
					for _, subPath := range getAllSubPaths(cwd) {
						dirToSize[subPath] += fileSize
					}
				}
			}
		}
	}
	return dirToSize
}

// path e.g. "/a/b"
// returns ["", "/a", "/a/b"]
func getAllSubPaths(path string) []string {
	var subPaths []string
	folderNames := strings.Split(path, "/") // ["", "a", "b", ""]
	for i := 0; i < len(folderNames)-1; i++ {
		subPath := strings.Join(folderNames[:len(folderNames)-1-i], "/") + "/"
		subPaths = append(subPaths, subPath)
	}
	return subPaths
}

func sumAllDirsWithSizeLE(dirToSize map[string]int, le int) int {
	total := 0
	for dir, size := range dirToSize {
		if size <= le {
			total += dirToSize[dir]
		}
	}
	return total
}

func findSizeOfDirToDelete(dirToSize map[string]int, maxSize int) int {
	occupiedSpace := dirToSize["/"]
	minSpaceToFree := occupiedSpace - maxSize
	actualSpaceToFree := occupiedSpace // initialise this as max
	for _, val := range dirToSize {
		if val >= minSpaceToFree && val < actualSpaceToFree {
			actualSpaceToFree = val
		}
	}
	return actualSpaceToFree
}
