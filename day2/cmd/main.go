package cmd2

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync/atomic"
)


func Main(path string) int {
	content, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	// Convert the byte slice to a string.
	fileContent := string(content)
	intervals := strings.Split(fileContent, ",")
	var solution atomic.Int64
	doneChan := make(chan struct{})
	
	for _, interval := range(intervals) {
		intervalSlice := strings.Split(interval, "-")
		start := intervalSlice[0]
		end := intervalSlice[1]

		if err != nil {
			log.Println(err)
		}

		go func(s string, e string, sol *atomic.Int64) {
			beginning, err := strconv.Atoi(s)
			if err != nil {
				log.Println("not valid beginning")
			}
			ending, err := strconv.Atoi(e)
			if err != nil {
				log.Println("not valid ending")
			}
			for i:=beginning; i<=ending; i++ {
				backToString := fmt.Sprintf("%d", i)
				if len(backToString) < 2 {
					continue
				}
				firstHalf := backToString[0:len(backToString)/2]
				secondHalf := backToString[len(backToString)/2:]
				if firstHalf == secondHalf {
					log.Printf("found an illegal %d\n", i)
					sol.Add((int64)(i))
				}
			}
			doneChan <- struct{}{}
		}(start, end, &solution)
	}

	for i := 0; i < len(intervals); i++ {
		<-doneChan
	}

	log.Printf("Solution 2 %v\n", solution.Load())

	return int(solution.Load())
 }