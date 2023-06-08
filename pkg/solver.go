package pkg

type Solver struct {
	grid *Grid
}

func NewSolver(grid *Grid) *Solver {
	return &Solver{grid: grid}
}

// DfsBacktrackSolve solves the Sudoku puzzle using a backtracking algorithm
func (s *Solver) DfsBacktrackSolve() bool {
	row, col, possibleValues := s.findMostConstrainingCell()

	// If there are no empty cells in the grid, the puzzle is solved
	if row == -1 && col == -1 {
		return true
	}

	for _, num := range possibleValues {
		if s.grid.IsValueValid(row, col, num) {
			s.grid.SetValue(row, col, num)

			// Recursively try to solve the rest of the grid
			if s.DfsBacktrackSolve() {
				return true
			}

			// If no valid configuration was found, backtrack and reset the cell
			s.grid.SetValue(row, col, 0)
		}
	}

	return false
}

// findMostConstrainingCell returns the empty cell with the fewest valid values
func (s *Solver) findMostConstrainingCell() (int, int, []int) {
	minOptions := 10
	row := -1
	col := -1
	var possibleValues []int

	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {

			// If the cell is not empty, skip it
			if s.grid.GetValue(i, j) != 0 {
				continue
			}

			// Count the number of valid values for this cell
			count := 0
			var cellPossibleValues []int
			for num := 1; num <= 9; num++ {
				if s.grid.IsValueValid(i, j, num) {
					count++
					cellPossibleValues = append(cellPossibleValues, num)
				}
			}

			// If the number of valid values is less than the current minimum, update the minimum
			if count < minOptions {
				minOptions = count
				row = i
				col = j
				possibleValues = cellPossibleValues
			}

			// If the number of valid values is 1, return immediately
			if minOptions == 1 {
				return row, col, possibleValues
			}
		}
	}

	return row, col, possibleValues
}
