package cmd3

import (
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
	lines := strings.Split(fileContent, "\n")
	var solution atomic.Int64	
	doneChan := make(chan struct{})
	defer close(doneChan)

	for _,line := range lines {
		
		go func(line string, tot *atomic.Int64)  {
			defer func() {
				doneChan <- struct{}{}
			}()
			
			 power := joltageV2(line, tot)
			 (*tot).Add(int64(power))
			 
		}(line, &solution)
	}

	for range lines {
		<-doneChan
	}

	return int(solution.Load())

}


func joltage(line string, total *atomic.Int32) int {

	
	slice := strings.Split(line, "")
	var firstBatteryPosition int
	var firstBatteryValue string
	
	//var secondBatteryPosition int
	var secondBatteryValue string

	for i,v := range slice {
		if v > firstBatteryValue && i < len(slice)-1 {
			firstBatteryValue = v
			firstBatteryPosition = i
		}
	}

	for _,v := range slice[firstBatteryPosition+1:] {
		if v > secondBatteryValue {
			secondBatteryValue = v
			//secondBatteryPosition = i
		}
	}

	sol, err := strconv.Atoi(firstBatteryValue+secondBatteryValue)
	if err != nil {
		log.Printf("can't convert solution to int")
	} else {
		log.Printf("joltage is %d", sol)
	}

	return sol
}

// greedy
func joltageV2(line string, total *atomic.Int64) int64 {

	
	const k = 12 // The desired length of the number

	if len(line) < k {
		log.Printf("line is too short to find a %d-digit number: %s", k, line)
		return 0
	}

	var resultBuilder strings.Builder
	resultBuilder.Grow(k)

	currentSearchIndex := 0

	for i := 0; i < k; i++ {
		// Determine the end of the search window for the current digit.
		// We must leave enough characters for the remaining digits.
		remainingNeeded := k - (i + 1)
		endSearchIndex := len(line) - 1 - remainingNeeded

		// Find the best character ('0' through '9') in the current window
		bestChar := byte('0' - 1) // Start with a character smaller than '0'
		bestCharIndex := -1

		for j := currentSearchIndex; j <= endSearchIndex; j++ {
			if line[j] > bestChar {
				bestChar = line[j]
				bestCharIndex = j
			}
		}

		// Append the best character found to our result
		resultBuilder.WriteByte(bestChar)

		// The next search must start after the character we just picked
		currentSearchIndex = bestCharIndex + 1
	}

	resultStr := resultBuilder.String()
	sol, err := strconv.Atoi(resultStr)
	if err != nil {
		log.Printf("can't convert result to int: %s", resultStr)
		return 0
	}

	log.Printf("joltageV2 found: %d", sol)
	return (int64)(sol)
}