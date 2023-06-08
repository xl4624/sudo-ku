package pkg

import (
	"testing"
)

func TestSolve(t *testing.T) {
	grid := NewGrid()

	// Create a simple Sudoku puzzle
	arr := [81]int{
		5, 3, 0, 0, 7, 0, 0, 0, 0,
		6, 0, 0, 0, 9, 0, 0, 0, 0,
		0, 9, 8, 0, 0, 0, 0, 6, 0,
		8, 0, 0, 0, 6, 0, 0, 0, 3,
		0, 0, 0, 8, 0, 3, 0, 0, 1,
		7, 0, 0, 0, 2, 0, 0, 0, 6,
		0, 6, 0, 0, 0, 0, 2, 8, 0,
		0, 0, 0, 4, 1, 9, 0, 0, 5,
		0, 0, 0, 0, 8, 0, 0, 7, 9,
	}

	// Use the InputArrayToGrid function to populate the grid with the array
	grid.InputArrayToGrid(arr)

	solver := NewSolver(&grid)

	if !solver.DfsBacktrackSolve() {
		t.Errorf("expected DfsBacktrackSolve to return %v, got %v", true, false)
	}

	// Check if the puzzle was solved correctly
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid.GetValue(i, j) == 0 {
				t.Errorf("expected grid at (%d, %d) to be filled, got %d", i, j, grid.GetValue(i, j))
			}
		}
	}
}
