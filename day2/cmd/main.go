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
	defer close(doneChan)
	
	for _, interval := range(intervals) {
		intervalSlice := strings.Split(interval, "-")
		start := intervalSlice[0]
		end := intervalSlice[1]

		if err != nil {
			log.Println(err)
		}

		go sumIllegals_1(start, end, &solution, doneChan)
	}

	for range intervals {
		<-doneChan
	}

	log.Printf("Solution 2 %v\n", solution.Load())

	return int(solution.Load())
 }

 func sumIllegals_1(s string, e string, sol *atomic.Int64, doneChan chan struct{}) {
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
}

func sumIllegals_2(s string, e string, sol *atomic.Int64, doneChan chan struct{}) {
	panic("not yet baby")
 }