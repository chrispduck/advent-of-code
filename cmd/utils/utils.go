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

func Abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}
