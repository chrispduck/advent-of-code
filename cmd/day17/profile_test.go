package main

import (
	"runtime"
	"testing"
)

func init() {
	runtime.SetCPUProfileRate(100_000)
}

func TestProfile(t *testing.T) {
	part1("example_input.txt", 2023)
}
