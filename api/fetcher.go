package api

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/chromedp/chromedp"
	"github.com/xl4624/Sudoku/pkg"
)

type PuzzleFetcher struct {
	Difficulty string
}

func NewPuzzleFetcher() *PuzzleFetcher {
	return &PuzzleFetcher{
		Difficulty: "easy",
	}
}

func (f *PuzzleFetcher) SetDifficulty(difficulty string) {
	f.Difficulty = difficulty
}

func (f *PuzzleFetcher) FetchPuzzle() (pkg.Grid, error) {
	grid := pkg.NewGrid()
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	var res string
	err := chromedp.Run(ctx,
		chromedp.Navigate(fmt.Sprintf("https://sudoku.com.au/%s.aspx", f.Difficulty)),
		chromedp.WaitVisible("#sudokutable", chromedp.ByID),
		chromedp.OuterHTML("#sudokutable", &res, chromedp.NodeVisible, chromedp.ByID),
	)
	if err != nil {
		log.Fatal(err)
	}

	doc, err := goquery.NewDocumentFromReader(strings.NewReader(res))
	if err != nil {
		log.Fatal(err)
	}

	doc.Find("td").Each(func(index int, element *goquery.Selection) {
		id, _ := element.Attr("id")
		value := element.Text()

		if strings.HasPrefix(id, "T") {
			idx, err := strconv.Atoi(id[1:])
			if err != nil {
				return
			}

			if value != "&nbsp;" {
				val, err := strconv.Atoi(value)
				if err != nil {
					return
				}

				row := idx / 9
				col := idx % 9
				if err := grid.SetSafeValue(row, col, val); err != nil {
					return
				}
			}
		}
	})

	return grid, nil
}
