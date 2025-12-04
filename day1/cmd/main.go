package cmd1

import (
	"bufio"
	"os"
	"path/filepath"
	"strconv"
	"go.uber.org/zap"
)

type Ring struct {
	size int
	list []int
	current int
}

const (
	RING_SIZE = 100
	INITIAL_POSITION = 50
)

func initLog() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()       
}

func Main() {

	logger := initLog()
	defer logger.Sync()

	var solution int
	path := filepath.Join("day1", "input1")
	file, err := os.Open(path)

	if err != nil {
		logger.Errorf("Cant read input file %s", path)
	}
	defer file.Close()

	ring := Ring{RING_SIZE, make([]int, RING_SIZE), INITIAL_POSITION} 
	
	for i := range(RING_SIZE) {
		ring.list[i] = i
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		switch line[0] {
			case 'L':
				num, err := strconv.Atoi(line[1:])
				if err != nil {
					logger.Errorf("invalid L input %s" + line[1:])
				}
				solution = (&ring).moveLeft(num, &solution)
			case 'R':
				num, err := strconv.Atoi(line[1:])
				if err != nil {
					logger.Errorf("invalid R input %s" + line[1:])
				}
				solution = (&ring).moveRigth(num, &solution)
		default:
			logger.Infof("input line %s will not be used", line)
		}
	}


	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		logger.Error(err)
	}

	logger.Infof("Solution is %d", solution)
}

func (r *Ring) moveLeft(n int, sol *int) (int) {
	effectiveN := n % r.size 
    r.current = (r.current - effectiveN + r.size) % r.size
	
	if r.current == 0 {
		*sol++
	}

	return *sol
}

func (r *Ring) moveRigth(n int, sol *int) (int) {
	r.current = (r.current + n) % r.size

	if r.current == 0 {
		*sol++
	}

	return *sol
}