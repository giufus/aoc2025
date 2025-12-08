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

		go sumIllegals_2(start, end, &solution, doneChan)
	}

	for range intervals {
		<-doneChan
	}

	log.Printf("Solution 2 %v\n", solution.Load())

	return int(solution.Load())
 }

 func sumIllegals_1(s string, e string, sol *atomic.Int64, doneChan chan struct{}) {
	defer func() {
		doneChan <- struct{}{}
	}()

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
		if len(backToString) % 2 != 0 {
			continue
		}
		firstHalf := backToString[0:len(backToString)/2]
		secondHalf := backToString[len(backToString)/2:]
		if firstHalf == secondHalf {
			log.Printf("found an illegal %d\n", i)
			sol.Add((int64)(i))
		}
	}
}

func sumIllegals_2(s string, e string, sol *atomic.Int64, doneChan chan struct{}) {
	defer func() {
		doneChan <- struct{}{}
	}()

	beginning, err := strconv.Atoi(s)
	if err != nil {
		log.Printf("not valid beginning in interval %s-%s", s, e)
		return
	}
	ending, err := strconv.Atoi(e)
	if err != nil {
		log.Printf("not valid ending in interval %s-%s", s, e)
		return
	}

	for i := beginning; i <= ending; i++ {
		backToString := fmt.Sprintf("%d", i)
		n := len(backToString)

		if n < 2 {
			continue
		}

		// The length 'l' of the potential repeating pattern
		// must be a divisor of the total string length.
		for l := 1; l <= n/2; l++ {
			if n%l == 0 {
				pattern := backToString[0:l]
				expected := strings.Repeat(pattern, n/l)
				if backToString == expected {
					//log.Printf("found an illegal %d\n", i)
					sol.Add(int64(i))
					// Break the inner loop to avoid adding the same number multiple times
					// (e.g., 888888 is a repeat of "8", "88", and "888").
					break
				}
			}
		}
	}
}