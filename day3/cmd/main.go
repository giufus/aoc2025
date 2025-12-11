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
	var solution atomic.Int32
	doneChan := make(chan struct{})
	defer close(doneChan)

	for _,line := range lines {
		
		go func(line string, tot *atomic.Int32)  {
			defer func() {
				doneChan <- struct{}{}
			}()
			
			 power := joltage(line, tot)
			 (*tot).Add(int32(power))
			 
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