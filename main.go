package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("=============")
	fmt.Println("King's Valley")
	fmt.Println("=============")
	fmt.Println("Move format: (C)olumn(R)ow(D)irection")
	fmt.Println("ex: A5U moves the piece in cell A5 (bottom left) in the Upward direction.")
	fmt.Println("Inputting 'quit' will end the game.")
	fmt.Println("Press enter to begin")
	reader.ReadString('\n')

	done := false
	for !done {
		g := newGame(*reader)
		g.play()

		fmt.Println("Play again? (Y/n)")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.ToLower(strings.TrimRight(userInput, "\n"))
		if userInput != "" && userInput[0] != 'y' {
			done = true
		}
	}
}
