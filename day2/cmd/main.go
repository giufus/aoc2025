package cmd2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)


func Main(path string) int {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string.
	fileContent := string(content)
	intervals := strings.Split(fileContent, ",")
	twices := make(chan int)
	
	for _, interval := range(intervals) {
		intervalSlice := strings.Split(interval, "-")
		start := intervalSlice[0]
		end := intervalSlice[1]

		if err != nil {
			log.Println(err)
		}
		go func(s string, e string, t chan int) {
			s, err := str.Atoi(s)
			if err != null {
				log.Println("not valid start")
			}
			for i:=s; i<=e; i++ {
				
			}
		}(start, end, twices)

	}
	fmt.Scanln()

	return 0
 }