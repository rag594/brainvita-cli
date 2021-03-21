package main

import (
	"fmt"
)

type Board [7][7]int

type Player struct {
	name string
}

type Cell struct {
	row int
	col int
}

type Game struct {
	player      Player
	board       Board
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
	if game.board[dest.row][dest.col] == -1 {
		return false
	}

	if game.board[dest.row][dest.col] == 0 {
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
				if game.board[source.row-1][dest.col-1] == -1 {
					return false
				}
			} else {
				if game.board[source.row+1][dest.col+1] == -1 {
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
		game.board[source.row][source.col] = 0
		game.board[dest.row][dest.col] = 1

		// Mark the middle cell as 0

		// vertically aligned
		if source.col == dest.col {
			if source.row > dest.row {
				game.board[source.row-1][source.col] = 0
			}

			if source.row < dest.row {
				game.board[source.row+1][source.col] = 0
			}
		}

		// horizontally aligned
		if source.row == dest.row {
			if source.col > dest.col {
				game.board[source.row][source.col-1] = 0
			}

			if source.col < dest.col {
				game.board[source.row][source.col+1] = 0
			}
		}

		// check if 2 cells are diagnol
		if source.row > dest.row && source.col > dest.col {
			game.board[source.row-1][source.col-1] = 0
		}

		if source.row < dest.row && source.col < dest.col {
			game.board[source.row+1][source.col+1] = 0
		}
	}
}

func (b Board) format() {
	fmt.Println()
	fmt.Printf("      %d  %d  %d        ", b[0][2], b[0][3], b[0][4])
	fmt.Println()
	fmt.Printf("      %d  %d  %d        ", b[1][2], b[1][3], b[1][4])
	fmt.Println()
	fmt.Printf("%d  %d  %d  %d  %d  %d  %d", b[2][0], b[2][1], b[2][2], b[2][3], b[2][4], b[2][5], b[2][6])
	fmt.Println()
	fmt.Printf("%d  %d  %d  %d  %d  %d  %d", b[3][0], b[3][1], b[3][2], b[3][3], b[3][4], b[3][5], b[3][6])
	fmt.Println()
	fmt.Printf("%d  %d  %d  %d  %d  %d  %d", b[4][0], b[4][1], b[4][2], b[4][3], b[4][4], b[4][5], b[4][6])
	fmt.Println()
	fmt.Printf("      %d  %d  %d        ", b[5][2], b[5][3], b[5][4])
	fmt.Println()
	fmt.Printf("      %d  %d  %d        ", b[6][2], b[6][3], b[6][4])
	fmt.Println()
}

func main() {

	var name string
	var srow, scol, drow, dcol int

	fmt.Println("Enter your name to play")
	fmt.Scanf("%s", &name)
	player := Player{name: name}

	board := Board{
		{-1, -1, 1, 1, 1, -1, -1},
		{-1, -1, 1, 1, 1, -1, -1},
		{1, 1, 1, 1, 1, 1, 1},
		{1, 1, 1, 0, 1, 1, 1},
		{1, 1, 1, 1, 1, 1, 1},
		{-1, -1, 1, 1, 1, -1, -1},
		{-1, -1, 1, 1, 1, -1, -1},
	}

	game := Game{player: player, board: board}

	fmt.Println("********Staring the Game**********")
	for {
		//Source
		fmt.Println("Enter the cell loc for source peg")
		fmt.Scanf("%d", &srow)
		fmt.Scanf("%d", &scol)
		source := Cell{srow, scol}

		//Destination
		fmt.Println("Enter the cell loc for dest peg")
		fmt.Scanf("%d", &drow)
		fmt.Scanf("%d", &dcol)
		destination := Cell{drow, dcol}

		// validate the source and dest cells for the move
		isValid := source.isValidMove(destination, game)

		game.isValidMove = isValid

		source.move(destination, &game)

		game.board.format()

	}
}
