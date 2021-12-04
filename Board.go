package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Board struct {
	height, width int
	cells         [][]uint
}

func NewBoard(width int, height int, isEmpty bool) *Board {
	board := Board{
		width:  width,
		height: height,
	}

	rand.Seed(time.Now().Unix())
	cells := make([][]uint, height)
	for y := 0; y < height; y++ {
		cells[y] = make([]uint, width)
		for x := 0; x < width; x++ {
			if !isEmpty && rand.Intn(100) > 80 {
				cells[y][x] = 1
			} else {
				cells[y][x] = 0
			}
		}
	}

	board.cells = cells

	return &board
}

func (b *Board) DrawBoard() {
	for y, YRange := range b.cells {
		for x := range YRange {
			if b.cells[y][x] == 1 {
				fmt.Print("*")
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}

func (b *Board) AliveNeighbors(x int, y int) uint {
	alive := uint(0)

	for i := -1; i <= 1; i++ {
		for j := -1; j <= 1; j++ {
			if x+i < 0 || x+i >= b.width {
				continue
			}

			if y+j < 0 || y+j >= b.height {
				continue
			}
			if x+i == x && y+j == y {
				continue
			}

			alive += b.cells[j+y][x+i]
		}
	}

	return alive
}

func (b *Board) ComputeNextGeneration() *[][]uint {
	currentBoard := b
	newBoard := make([][]uint, b.height)

	for y := 0; y < currentBoard.height; y++ {
		newBoard[y] = make([]uint, b.width)
		for x := 0; x < currentBoard.width; x++ {
			currentCell := b.cells[y][x]
			neighbors := b.AliveNeighbors(x, y)

			if currentCell == 1 && (neighbors < 2 || neighbors > 3) {
				newBoard[y][x] = 0
			} else if currentCell == 0 && neighbors == 3 {
				newBoard[y][x] = 1
			} else {
				newBoard[y][x] = currentBoard.cells[y][x]
			}
		}
	}

	return &newBoard
}
