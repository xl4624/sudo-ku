# Sudo-ku

This project is a command-line program written in Go that solves Sudoku puzzles. It fetches a puzzle from [Sudoku.com.au](https://sudoku.com.au/), solves it, and displays the original puzzle and its solution in the console.

## Getting Started

To run the project, you'll need to have Go installed on your machine. If you don't have Go installed, you can download it from the official [Go download page](https://golang.org/dl/).

After Go is installed, you can download and run the project as follows:

```bash
# Clone the repository
git clone https://github.com/xl4624/Sudo-ku.git

# Go to the project directory
cd sudo-ku

# Run the project, replacing "difficulty" with the difficulty level of the puzzle you want to solve
go run main.go difficulty
```

The possible difficulty levels are "easy", "medium", "hard", "tough", and "parents".

## Implementation Details

The program uses a backtracking algorithm to solve the puzzles. It begins by finding the cell with the fewest valid numbers (the "most constraining" cell). Then it tries all possible valid numbers for that cell, and for each one, it recursively tries to solve the rest of the grid. If it eventually finds a number for the cell that leads to a contradiction (i.e., there's no valid number for a future cell), it backtracks and tries a different number. If it has tried all possible numbers for the current cell and none of them work, it backtracks to the previous cell and tries the next number for that cell. This process continues until it finds a solution or determines that no solution is possible.

## License

This project is licensed under the terms of the MIT license. See the [LICENSE](LICENSE) file for details.
