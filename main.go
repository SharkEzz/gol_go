package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var width, height, sleepTime uint

	flag.UintVar(&width, "w", 40, "The map width (default 40)")
	flag.UintVar(&height, "h", 20, "The map height (default 20)")
	flag.UintVar(&sleepTime, "s", 50, "The delay between two generation in ms (default 50ms)")

	flag.Parse()

	board := NewBoard(int(width), int(height), false)

	fmt.Print("\033[?25l") // Disable cursor
	fmt.Print("\033[2J")   // Clear terminal

	for i := 0; true; i++ {

		fmt.Printf("\033[%d;%dH", 0, 0) // Set cursor position
		fmt.Printf("Génération : %v", i)
		board.DrawBoard()
		board.cells = *board.ComputeNextGeneration()
		time.Sleep(time.Millisecond * time.Duration(sleepTime))
	}

	fmt.Print("\033[?25h") // Enable cursor
}
