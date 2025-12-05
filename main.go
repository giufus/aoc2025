package main

import (
	cmd1 "aoc2025/day1/cmd"
	cmd2 "aoc2025/day2/cmd"
	"fmt"
	"path/filepath"
)


func main() {
	path := filepath.Join("day1", "input")
	fmt.Println("Solution is", cmd1.Main(path))
	path = filepath.Join("day2", "input")
	fmt.Println("Solution is", cmd2.Main(path))
}