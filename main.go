package main

import (
	"fmt"
	"log"
	"os"
    "time"

	"github.com/xl4624/Sudo-ku/api"
	"github.com/xl4624/Sudo-ku/pkg"
)

func main() {
	if len(os.Args) < 2 {
		log.Fatal("Please provide a difficulty level: easy, medium, hard, tough, or test")
	}
	difficulty := os.Args[1]

	switch difficulty {
	case "easy", "medium", "hard", "tough":
    case "test":
        testSolveTime()
        return
	default:
		log.Fatalf("Invalid difficulty level: %v. Please provide a difficulty level: easy, medium, hard, tough", difficulty)
	}

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
}

func testSolveTime() {
    difficulties := []string{"easy", "medium", "hard", "tough"}
    sampleSize := 5

    for _, difficulty := range difficulties {
        var totalDuration time.Duration
        for i := 0; i < sampleSize; i++ {
            fetcher := api.NewPuzzleFetcher()
            fetcher.SetDifficulty(difficulty)
            grid, err := fetcher.FetchPuzzle()
            if err != nil {
                log.Fatalf("Failed to fetch puzzle: %v", err)
            }

            start := time.Now()
            solver := pkg.NewSolver(&grid)
            solver.DfsBacktrackSolve()
            totalDuration += time.Since(start)
        }

        averageDuration := totalDuration / time.Duration(sampleSize)
        fmt.Printf("Average time to solve %v puzzle: %v\n", difficulty, averageDuration)
    }
}
