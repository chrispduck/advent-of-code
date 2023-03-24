package utils

import (
	"log"
	"strconv"
)

func CheckErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

func StrToInt(s string) int {
	i, err := strconv.Atoi(s)
	CheckErr(err)
	return i
}

func Max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

type Coordinate struct {
	X, Y int
}

func (c Coordinate) Add(to Coordinate) Coordinate {
	return Coordinate{
		c.X + to.X,
		c.Y + to.Y,
	}
}

func (c Coordinate) Subtract(to Coordinate) Coordinate {
	return Coordinate{
		c.X - to.X,
		c.Y - to.Y,
	}
}

func (c Coordinate) L1Distance(to Coordinate) int {
	return Abs(c.X-to.X) + Abs(c.Y-to.Y)
}

func (c Coordinate) WrapAround(width, height int) Coordinate {
	return Coordinate{
		(c.X%width + width) % width,
		(c.Y%height + height) % height,
	}
}

type SetOfStrings struct {
	items map[string]interface{}
}

type Set interface {
	Add(string)
}

func IdempotentRemove(items []string, item string) []string {
	for idx, x := range items {
		if x == item {
			return append(items[:idx], items[idx+1:]...)
		}
	}
	return items
}

func IdempotentAdds(items []string, itemsToAdd []string) []string {
	for _, item := range itemsToAdd {
		items = IdempotentAdd(items, item)
	}
	return items
}

func IdempotentAdd(items []string, item string) []string {
	for _, x := range items {
		if x == item {
			return items
		}
	}
	return append(items, item)
}
