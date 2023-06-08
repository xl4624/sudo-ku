package pkg

import (
	"errors"
	"fmt"
)

var (
	ErrInvalidInputType  = errors.New("input must be an integer")
	ErrInvalidValueRange = errors.New("input must be an integer in range 1-9")
)

type Grid [9][9]int

func NewGrid() Grid {
	return Grid{}
}

// SetValue sets the value of a cell in the grid
func (g *Grid) SetValue(row, col, val int) {
	g[row][col] = val
}

// SetSafeValue sets the value of a cell in the grid if it is valid
func (g *Grid) SetSafeValue(row, col, val int) error {
	if val < 1 || val > 9 {
		return ErrInvalidValueRange
	}

	if !g.IsValueValid(row, col, val) {
		return ErrInvalidInputType
	}

	g.SetValue(row, col, val)
	return nil
}

// GetValue returns the value of a cell in the grid
func (g *Grid) GetValue(row, col int) int {
	return g[row][col]
}

// IsValueValid checks if the given value is valid at the given row and column
func (g *Grid) IsValueValid(row, col, val int) bool {
	// Check the row and column
	for i := 0; i < 9; i++ {
		if i != col && g.GetValue(row, i) == val || i != row && g.GetValue(i, col) == val {
			return false
		}
	}

	// Check the box
	startRow, startCol := (row/3)*3, (col/3)*3
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if startRow+i != row && startCol+j != col && g.GetValue(startRow+i, startCol+j) == val {
				return false
			}
		}
	}

	return true
}

// Clear resets all the values of the grid to 0
func (g *Grid) Clear() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			g.SetValue(i, j, 0)
		}
	}
}

// Display prints the grid to the console
func (g *Grid) Display() {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Print(g.GetValue(i, j))
			if j != 8 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}

// InputArrayToGrid converts an array of 81 integers and sets the values of the grid
func (g *Grid) InputArrayToGrid(arr [81]int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if arr[i*9+j] != 0 {
				g.SetValue(i, j, arr[i*9+j])
			}
		}
	}
}

// ToOutputArray converts the grid to an array of 81 integers and returns it
func (g *Grid) ToOutputArray() [81]int {
	var outputArray [81]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			outputArray[i*9+j] = g.GetValue(i, j)
		}
	}

	return outputArray
}
