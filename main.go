package main

import (
	cmd1 "aoc2025/day1/cmd"
	"fmt"
	"path/filepath"
)


func main() {
	path := filepath.Join("day1", "input1")
	fmt.Println("Solution is", cmd1.Main(path))
}