package main

import (
	"fmt"
)

func main() {

	// create a tic tac toe board
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}
	// we're predesigning a game for them to play turn by turn
	//ok lets just randomly assign numbers over here we need to take user input for now

	//for i := 0; i < len(board); i++ {
	//	fmt.Printf("%s\n", board)
	//}
	fmt.Printf("%s\n", board)

}
