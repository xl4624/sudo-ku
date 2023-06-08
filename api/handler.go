package api

import (
	"encoding/json"
	"net/http"

	"github.com/xl4624/Sudoku/pkg"
)

type Handler struct{}

type SolveResponse struct {
	Grid  [81]int `json:"Grid"`
	Error string  `json:"Error,omitempty"`
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) SolveHandler(w http.ResponseWriter, r *http.Request) {
	type SolveRequest struct {
		Grid [81]int `json:"Grid"`
	}

	var req SolveRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	grid := pkg.NewGrid()
	grid.InputArrayToGrid(req.Grid)

	solver := pkg.NewSolver(&grid)
	if solved := solver.DfsBacktrackSolve(); !solved {
		resp := SolveResponse{Error: "Could not solve Sudoku puzzle"}
		w.WriteHeader(http.StatusUnprocessableEntity)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Could not solve Sudoku puzzle", http.StatusInternalServerError)
		}
		return
	}

	resp := SolveResponse{Grid: grid.ToOutputArray()}
	if err := json.NewEncoder(w).Encode(resp); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
