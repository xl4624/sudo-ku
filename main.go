package main

import (
	"fmt"
	"log"
	"os"

	"github.com/xl4624/Sudoku/api"
	"github.com/xl4624/Sudoku/pkg"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a difficulty level: easy, medium, hard, tough, parents")
	}
	difficulty := os.Args[1]

	switch difficulty {
	case "easy", "medium", "hard", "tough", "parents":
	default:
		log.Fatalf("Invalid difficulty level: %v. Please provide a difficulty level: easy, medium, hard, tough, parents", difficulty)
	}

	// handler := api.NewHandler()
	// http.HandleFunc("/api/solve", handler.SolveHandler)

	fetcher := api.NewPuzzleFetcher()
	fetcher.SetDifficulty(difficulty)
	grid, err := fetcher.FetchPuzzle()
	if err != nil {
		log.Fatalf("Failed to fetch puzzle: %v", err)
	}

	fmt.Println("Puzzle:")
	grid.Display()

	solver := pkg.NewSolver(&grid)
	if solved := solver.DfsBacktrackSolve(); !solved {
		fmt.Println("Could not solve Sudoku puzzle")
	} else {
		fmt.Println("Solution:")
		grid.Display()
	}

	// log.Fatal(http.ListenAndServe(":8000", nil))
}
