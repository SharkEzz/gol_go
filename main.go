package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var width uint
	var height uint

	flag.UintVar(&width, "w", 40, "The map width (default 40)")
	flag.UintVar(&height, "h", 20, "The map height (default 20)")

	flag.Parse()

	board := NewBoard(int(width), int(height), false)

	fmt.Print("\033[?25l")
	fmt.Print("\033[2J")

	for i := 0; true; i++ {

		fmt.Printf("\033[%d;%dH", 0, 0)
		fmt.Printf("Génération : %v", i)
		board.DrawBoard()
		time.Sleep(time.Millisecond * 50)
		board = board.ComputeNextGeneration()
	}

	fmt.Print("\033[?25h")
}
