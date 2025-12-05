package cmd1

import (
	"bufio"
	"os"
	"strconv"

	"go.uber.org/zap"
)

type RingA struct {
	size    int
	list    []int
	current int
}

type RingB RingA

type SolutionStrategy interface {
	MoveLeft(n int, sol *int) int
	MoveRight(n int, sol *int) int
}

const (
	RING_SIZE        = 100
	INITIAL_POSITION = 50
)

func initLog() *zap.SugaredLogger {
	logger, _ := zap.NewProduction()
	return logger.Sugar()
}

func Main(path string) int {

	logger := initLog()
	defer logger.Sync()

	var solution int

	file, err := os.Open(path)

	if err != nil {
		logger.Errorf("Cant read input file %s", path)
	}
	defer file.Close()

	ring := RingB{RING_SIZE, make([]int, RING_SIZE), INITIAL_POSITION}

	for i := range RING_SIZE {
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
			solution = (&ring).MoveLeft(num, &solution)
		case 'R':
			num, err := strconv.Atoi(line[1:])
			if err != nil {
				logger.Errorf("invalid R input %s" + line[1:])
			}
			solution = (&ring).MoveRigth(num, &solution)
		default:
			logger.Infof("input line %s will not be used", line)
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		logger.Error(err)
	}

	logger.Infof("Solution is %d", solution)
	return solution
}

func (r *RingA) MoveLeft(n int, sol *int) int {
	r.current = (r.current - n + r.size) % r.size

	if r.current == 0 {
		*sol++
	}

	return *sol
}

func (r *RingA) MoveRigth(n int, sol *int) int {
	r.current = (r.current + n) % r.size

	if r.current == 0 {
		*sol++
	}

	return *sol
}


func (r *RingB) MoveLeft(n int, sol *int) int {
	var rounds int
	if r.current == 0 {
		rounds = n / r.size
	}

	if r.current != 0 && n >= r.current {
		rounds = ((n - r.current) / r.size) + 1
	}

	r.current = (((r.current - n) % r.size) + r.size) % r.size
	
	*sol += rounds

	return *sol
}

func (r *RingB) MoveRigth(n int, sol *int) int {
	
	rounds := (r.current + n) / r.size
	*sol += rounds

	r.current = (r.current + n) % r.size

	return *sol
}