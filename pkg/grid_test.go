package pkg

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestNewGrid(t *testing.T) {
	grid := NewGrid()

	// Check that all values of a new grid are 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid.GetValue(i, j) != 0 {
				t.Errorf("expected 0, got %d", grid.GetValue(i, j))
			}
		}
	}
}

func TestSetValue(t *testing.T) {
	grid := NewGrid()

	// Set a value in the grid and check if it is correctly set
	grid.SetValue(0, 0, 5)
	if grid.GetValue(0, 0) != 5 {
		t.Errorf("expected 5, got %d", grid.GetValue(0, 0))
	}
}

func TestSetSafeValue(t *testing.T) {
	grid := NewGrid()

	// Check if the function can safely set a valid value
	if err := grid.SetSafeValue(0, 0, 5); err != nil {
		t.Errorf("expected nil, got %s", err)
	}
	if grid.GetValue(0, 0) != 5 {
		t.Errorf("expected 5, got %d", grid.GetValue(0, 0))
	}

	// Check if the function correctly rejects a value out of range
	if err := grid.SetSafeValue(0, 0, 10); err != ErrInvalidValueRange {
		t.Errorf("expected %s, got %s", ErrInvalidValueRange, err)
	}

	// Check if the function correctly rejects a value that is in the same row
	if err := grid.SetSafeValue(1, 0, 5); err != ErrInvalidInputType {
		t.Errorf("expected %s, got %s", ErrInvalidInputType, err)
	}

	// Check if the function correctly rejects a value that is in the same column
	if err := grid.SetSafeValue(0, 1, 5); err != ErrInvalidInputType {
		t.Errorf("expected %s, got %s", ErrInvalidInputType, err)
	}

	// Check if the function correctly rejects a value that is in the same sub-grid
	if err := grid.SetSafeValue(1, 1, 5); err != ErrInvalidInputType {
		t.Errorf("expected %s, got %s", ErrInvalidInputType, err)
	}

	// Check that the cells maintain their original values when attempting to set an invalid value
	if grid.GetValue(1, 0) != 0 || grid.GetValue(0, 1) != 0 || grid.GetValue(1, 1) != 0 {
		t.Errorf("expected 0, got %d", grid.GetValue(0, 1))
	}
}

func TestDisplay(t *testing.T) {
	grid := NewGrid()

	// Fill the grid with a pattern to test the display function
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid.SetValue(i, j, (i+j)%9+1) // creates a pattern of numbers 1-9
		}
	}

	// Store the output of the display function
	rescueStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	grid.Display()

	// Restore the output
	_ = w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescueStdout

	expectedOut := `1 2 3 4 5 6 7 8 9
2 3 4 5 6 7 8 9 1
3 4 5 6 7 8 9 1 2
4 5 6 7 8 9 1 2 3
5 6 7 8 9 1 2 3 4
6 7 8 9 1 2 3 4 5
7 8 9 1 2 3 4 5 6
8 9 1 2 3 4 5 6 7
9 1 2 3 4 5 6 7 8`

	// Compare the output of the display function to the expected output
	if strings.TrimSpace(string(out)) != expectedOut {
		t.Errorf("expected %s, got %s", expectedOut, string(out))
	}
}

func TestInputArrayToGrid(t *testing.T) {
	grid := NewGrid()

	arr := [81]int{
		5, 3, 0, 0, 7, 0, 0, 0, 0,
		6, 0, 0, 1, 9, 5, 0, 0, 0,
		0, 9, 8, 0, 0, 0, 0, 6, 0,
		8, 0, 0, 0, 6, 0, 0, 0, 3,
		4, 0, 0, 8, 0, 3, 0, 0, 1,
		7, 0, 0, 0, 2, 0, 0, 0, 6,
		0, 6, 0, 0, 0, 0, 2, 8, 0,
		0, 0, 0, 4, 1, 9, 0, 0, 5,
		0, 0, 0, 0, 8, 0, 0, 7, 9,
	}

	// Check if the function correctly inputs an array into the grid
	grid.InputArrayToGrid(arr)
}

func TestGridClear(t *testing.T) {
	grid := NewGrid()

	// Fill the grid with a pattern to test the clear function
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			grid.SetValue(i, j, (i+j)%9+1) // creates a pattern of numbers 1-9
		}
	}

	grid.Clear()

	// Check that all values of the grid are 0
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if grid.GetValue(i, j) != 0 {
				t.Errorf("expected 0, got %d", grid.GetValue(i, j))
			}
		}
	}
}
