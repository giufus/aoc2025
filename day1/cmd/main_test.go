package cmd1

import (
	"path/filepath"
	"testing"
)

func TestRingA_moveLeft(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name         string
		startRing    RingA
		moveAmount   int
		startSol     int
		expectedRing RingA
		expectedSol  int
	}{
		{
			name:         "Simple move left, no wrap",
			startRing:    RingA{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingA{size: 100, list: make([]int, 100), current: 40},
			expectedSol:  0,
		},
		{
			name:         "Move left with wrap around",
			startRing:    RingA{size: 100, list: make([]int, 100), current: 5},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingA{size: 100, list: make([]int, 100), current: 95},
			expectedSol:  0,
		},
		{
			name:         "Move left lands on 0, increments solution",
			startRing:    RingA{size: 100, list: make([]int, 100), current: 10},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingA{size: 100, list: make([]int, 100), current: 0},
			expectedSol:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a pointer to the solution to pass to the method
			sol := tc.startSol

			// Call the method
			tc.startRing.MoveLeft(tc.moveAmount, &sol)

			// Assert the results
			if tc.startRing.current != tc.expectedRing.current {
				t.Errorf("FAIL: current position incorrect. expected %d, got %d", tc.expectedRing.current, tc.startRing.current)
			}
			if sol != tc.expectedSol {
				t.Errorf("FAIL: solution incorrect. expected %d, got %d", tc.expectedSol, sol)
			}
		})
	}
}

func TestRingA_moveRight(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name         string
		startRing    RingA
		moveAmount   int
		startSol     int
		expectedRing RingA
		expectedSol  int
	}{
		{
			name:         "Simple move right, no wrap",
			startRing:    RingA{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingA{size: 100, list: make([]int, 100), current: 60},
			expectedSol:  0,
		},
		{
			name:         "Move right with wrap around",
			startRing:    RingA{size: 100, list: make([]int, 100), current: 95},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingA{size: 100, list: make([]int, 100), current: 5},
			expectedSol:  0,
		},
		{
			name:         "Move right lands on 0, increments solution",
			startRing:    RingA{size: 100, list: make([]int, 100), current: 90},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingA{size: 100, list: make([]int, 100), current: 0},
			expectedSol:  1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sol := tc.startSol

			tc.startRing.MoveRigth(tc.moveAmount, &sol)

			if tc.startRing.current != tc.expectedRing.current {
				t.Errorf("FAIL: current position incorrect. expected %d, got %d", tc.expectedRing.current, tc.startRing.current)
			}
			if sol != tc.expectedSol {
				t.Errorf("FAIL: solution incorrect. expected %d, got %d", tc.expectedSol, sol)
			}
		})
	}
}

func TestRingB_moveLeft(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name         string
		startRing    RingB
		moveAmount   int
		startSol     int
		expectedRing RingB
		expectedSol  int
	}{
		{
			name:         "Simple move left",
			startRing:    RingB{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingB{size: 100, list: make([]int, 100), current: 40},
			expectedSol:  0,
		},
		{
			name:         "Single move left",
			startRing:    RingB{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   68,
			startSol:     0,
			expectedRing: RingB{size: 100, list: make([]int, 100), current: 82},
			expectedSol:  1,
		},
		{
			name:         "Double move left",
			startRing:    RingB{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   210,
			startSol:     0,
			expectedRing: RingB{size: 100, list: make([]int, 100), current: 40},
			expectedSol:  2,
		},
		
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a pointer to the solution to pass to the method
			sol := tc.startSol

			// Call the method
			tc.startRing.MoveLeft(tc.moveAmount, &sol)

			// Assert the results
			if tc.startRing.current != tc.expectedRing.current {
				t.Errorf("FAIL: current position incorrect. expected %d, got %d", tc.expectedRing.current, tc.startRing.current)
			}
			if sol != tc.expectedSol {
				t.Errorf("FAIL: solution incorrect. expected %d, got %d", tc.expectedSol, sol)
			}
		})
	}
}

func TestRingB_moveRight(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name         string
		startRing    RingB
		moveAmount   int
		startSol     int
		expectedRing RingB
		expectedSol  int
	}{
		{
			name:         "Simple move right",
			startRing:    RingB{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   10,
			startSol:     0,
			expectedRing: RingB{size: 100, list: make([]int, 100), current: 60},
			expectedSol:  0,
		},
		{
			name:         "Single move right",
			startRing:    RingB{size: 100, list: make([]int, 100), current: 95},
			moveAmount:   60,
			startSol:     0,
			expectedRing: RingB{size: 100, list: make([]int, 100), current: 55},
			expectedSol:  1,
		},
		{
			name:         "Double move right",
			startRing:    RingB{size: 100, list: make([]int, 100), current: 50},
			moveAmount:   210,
			startSol:     0,
			expectedRing: RingB{size: 100, list: make([]int, 100), current: 60},
			expectedSol:  2,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sol := tc.startSol

			tc.startRing.MoveRigth(tc.moveAmount, &sol)

			if tc.startRing.current != tc.expectedRing.current {
				t.Errorf("FAIL: current position incorrect. expected %d, got %d", tc.expectedRing.current, tc.startRing.current)
			}
			if sol != tc.expectedSol {
				t.Errorf("FAIL: solution incorrect. expected %d, got %d", tc.expectedSol, sol)
			}
		})
	}
}


func TestMain(t *testing.T) {
	ans := Main(filepath.Join("..", "..", "day1", "input1test"))
	expected := 6
	if ans != expected {
        t.Errorf("Main = %d; want %d", ans, expected)
    }
 }

func BenchmarkMain(b *testing.B) {
	for b.Loop() {
		Main(filepath.Join("..", "..", "day1", "input"))
	}
 }