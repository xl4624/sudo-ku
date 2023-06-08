package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xl4624/Sudo-ku/pkg"
)

func TestSolveHandler(t *testing.T) {
	t.Run("returns UnprocessableEntity for unsolvable puzzle", func(t *testing.T) {
		grid := pkg.NewGrid()
		var puzzle = [81]int{
			8, 8, 3, 9, 2, 1, 6, 5, 7,
			6, 5, 7, 3, 4, 8, 2, 1, 9,
			1, 2, 9, 6, 0, 5, 0, 3, 4,
			9, 1, 0, 0, 6, 2, 3, 7, 5,
			3, 6, 0, 7, 9, 4, 1, 0, 2,
			2, 7, 0, 5, 1, 3, 9, 4, 6,
			5, 8, 1, 2, 3, 9, 7, 6, 4,
			4, 9, 6, 1, 8, 0, 5, 2, 3,
			7, 3, 2, 4, 5, 0, 0, 0, 8,
		}

		grid.InputArrayToGrid(puzzle)

		reqBody, _ := json.Marshal(map[string][81]int{
			"Grid": grid.ToOutputArray(),
		})
		req, err := http.NewRequest("POST", "api/solve", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()

		handler := NewHandler()
		handlerFunc := http.HandlerFunc(handler.SolveHandler)

		handlerFunc.ServeHTTP(rr, req)
		assert.Equal(t, http.StatusUnprocessableEntity, rr.Code)

		var resp SolveResponse
		json.Unmarshal(rr.Body.Bytes(), &resp)

		assert.Equal(t, "Could not solve Sudoku puzzle", resp.Error)
	})
}
