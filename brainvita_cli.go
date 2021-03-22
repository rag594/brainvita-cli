package main

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/rivo/tview"
)

type Player struct {
	name string
}

type Cell struct {
	row int
	col int
}

type Game struct {
	player      Player
	board       *tview.Table
	isValidMove bool
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (source Cell) isValidMove(dest Cell, game Game) bool {

	// dest should not be one of the dead cells
	if game.board.GetCell(dest.row, dest.col).Text == " " {
		return false
	}

	if game.board.GetCell(dest.row, dest.col).Text == "0" {
		// check if 2 cells are vertically aligned
		if source.col == dest.col {
			if Abs(source.row-dest.row) == 2 {
				return true
			} else {
				return false
			}
		}

		// check if 2 cells are horizontally aligned
		if source.row == dest.row {
			if Abs(source.col-dest.col) == 2 {
				return true
			} else {
				return false
			}
		}

		// check if 2 cells are diagnol
		if source.row != dest.row && source.col != dest.col {

			// Corner cases for cell jumping along diagnols containing -1 in b/w
			if source.row > dest.row {
				if game.board.GetCell(source.row-1, dest.col-1).Text == " " {
					return false
				}
			} else {
				if game.board.GetCell(source.row+1, dest.col+1).Text == " " {
					return false
				}
			}

			// for diagnol jumping the abs diff should be 2
			if Abs(source.row-dest.row) == 2 && Abs(source.col-dest.col) == 2 {
				return true
			} else {
				return false
			}
		}
	} else {
		return false
	}

	return true
}

func (source Cell) move(dest Cell, game *Game) {
	if game.isValidMove {
		// Peg is removed so mark it as 0
		game.board.GetCell(source.row, source.col).Text = "0"
		game.board.GetCell(dest.row, dest.col).Text = "1"

		// Mark the middle cell as 0

		// vertically aligned
		if source.col == dest.col {
			if source.row > dest.row {
				game.board.GetCell(source.row-1, source.col).Text = "0"
			}

			if source.row < dest.row {
				game.board.GetCell(source.row+1, source.col).Text = "0"
			}
		}

		// horizontally aligned
		if source.row == dest.row {
			if source.col > dest.col {
				game.board.GetCell(source.row, source.col-1).Text = "0"
			}

			if source.col < dest.col {
				game.board.GetCell(source.row, source.col+1).Text = "0"
			}
		}

		// check if 2 cells are diagnol
		if source.row > dest.row && source.col > dest.col {
			game.board.GetCell(source.row-1, source.col-1).Text = "0"
		}

		if source.row < dest.row && source.col < dest.col {
			game.board.GetCell(source.row+1, source.col+1).Text = "0"
		}
	}
}

func main() {

	//var name string
	//var srow, scol, drow, dcol int

	// fmt.Println("Enter your name to play")
	// fmt.Scanf("%s", &name)
	player := Player{name: "Rag"}

	app := tview.NewApplication()
	table := tview.NewTable().
		SetBorders(true)
	cols, rows := 7, 7
	for r := 0; r < rows; r++ {
		for c := 0; c < cols; c++ {
			color := tcell.ColorWhite

			if r == 3 && c == 3 {
				table.SetCell(r, c,
					tview.NewTableCell("0").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else if (r == 0 || r == 1) && (c == 0 || c == 1) {
				table.SetCell(r, c,
					tview.NewTableCell(" ").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else if (r == 5 || r == 6) && (c == 0 || c == 1 || c == 5 || c == 6) {
				table.SetCell(r, c,
					tview.NewTableCell(" ").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else if (r == 0 || r == 1) && (c == 5 || c == 6) {
				table.SetCell(r, c,
					tview.NewTableCell(" ").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			} else {
				table.SetCell(r, c,
					tview.NewTableCell("1").
						SetTextColor(color).
						SetAlign(tview.AlignCenter))
			}
		}
	}

	game := Game{player: player, board: table}

	cells := make([]Cell, 0, 2)

	fmt.Println("********Staring the Game**********")

	// validate the source and dest cells for the move
	game.board.Select(0, 0).SetFixed(1, 1).SetDoneFunc(func(key tcell.Key) {
		if key == tcell.KeyEscape {
			app.Stop()
		}
		if key == tcell.KeyEnter {
			//fmt.Println(game.board.GetSelection())
			game.board.SetSelectable(true, true)
		}
	}).SetSelectedFunc(func(row int, column int) {
		selectedCell := Cell{row: row, col: column}
		cells = append(cells, selectedCell)
		game.board.GetCell(row, column).SetTextColor(tcell.ColorRed)
		game.board.SetSelectable(false, false)
		fmt.Println(cells)

		if len(cells) == 2 {
			source, dest := cells[0], cells[1]
			isValid := source.isValidMove(dest, game)
			game.isValidMove = isValid
			if isValid {
				source.move(dest, &game)
			} else {
				fmt.Println(isValid)
			}
			cells = nil
			game.board.GetCell(source.row, source.col).SetTextColor(tcell.ColorWhite)
			game.board.GetCell(dest.row, dest.col).SetTextColor(tcell.ColorWhite)
		}
	})

	if err := app.SetRoot(table, true).EnableMouse(true).Run(); err != nil {
		panic(err)
	}
}
