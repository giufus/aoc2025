package main

import (
	cmd1 "aoc2025/day1/cmd"
	cmd2 "aoc2025/day2/cmd"
	cmd3 "aoc2025/day3/cmd"
	"fmt"
	"os"
	"path/filepath"
)


func main() {
	switch {
		case os.Args[1] == "1":
			path := filepath.Join("day1", "input")
			fmt.Println("Solution is", cmd1.Main(path))
		case os.Args[1] == "2":
			path := filepath.Join("day2", "input")
			fmt.Println("Solution is", cmd2.Main(path))
			fmt.Println("Solution is", cmd1.Main(path))
		case os.Args[1] == "3":
			path := filepath.Join("day3", "input")
			fmt.Println("Solution is", cmd3.Main(path))
	}
}