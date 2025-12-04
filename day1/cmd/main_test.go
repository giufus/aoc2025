package cmd1

import "testing"

func TestRing_moveLeft(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name           string
		startRing      Ring
		moveAmount     int
		startSol       int
		expectedRing   Ring
		expectedSol    int
	}{
		{
			name: "Simple move left, no wrap",
			startRing: Ring{size: 100, list:make([]int, 100), current: 50},
			moveAmount: 10,
			startSol: 0,
			expectedRing: Ring{size: 100, list:make([]int, 100), current: 40},
			expectedSol: 0,
		},
		{
			name: "Move left with wrap around",
			startRing: Ring{size: 100, list:make([]int, 100), current: 5},
			moveAmount: 10,
			startSol: 0,
			expectedRing: Ring{size: 100, list:make([]int, 100), current: 95},
			expectedSol: 0,
		},
		{
			name: "Move left lands on 0, increments solution",
			startRing: Ring{size: 100, list:make([]int, 100), current: 10},
			moveAmount: 10,
			startSol: 0,
			expectedRing: Ring{size: 100, list:make([]int, 100), current: 0},
			expectedSol: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Create a pointer to the solution to pass to the method
			sol := tc.startSol
			
			// Call the method
			tc.startRing.moveLeft(tc.moveAmount, &sol)

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


func TestRing_moveRight(t *testing.T) {
	// Define test cases in a table
	testCases := []struct {
		name           string
		startRing      Ring
		moveAmount     int
		startSol       int
		expectedRing   Ring
		expectedSol    int
	}{
		{
			name: "Simple move right, no wrap",
			startRing: Ring{size: 100, list:make([]int, 100), current: 50},
			moveAmount: 10,
			startSol: 0,
			expectedRing: Ring{size: 100, list:make([]int, 100), current: 60},
			expectedSol: 0,
		},
		{
			name: "Move right with wrap around",
			startRing: Ring{size: 100, list:make([]int, 100), current: 95},
			moveAmount: 10,
			startSol: 0,
			expectedRing: Ring{size: 100, list:make([]int, 100), current: 5},
			expectedSol: 0,
		},
		{
			name: "Move right lands on 0, increments solution",
			startRing: Ring{size: 100, list:make([]int, 100), current: 90},
			moveAmount: 10,
			startSol: 0,
			expectedRing: Ring{size: 100, list:make([]int, 100), current: 0},
			expectedSol: 1,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			sol := tc.startSol
			
			tc.startRing.moveRigth(tc.moveAmount, &sol)

			if tc.startRing.current != tc.expectedRing.current {
				t.Errorf("FAIL: current position incorrect. expected %d, got %d", tc.expectedRing.current, tc.startRing.current)
			}
			if sol != tc.expectedSol {
				t.Errorf("FAIL: solution incorrect. expected %d, got %d", tc.expectedSol, sol)
			}
		})
	}
}