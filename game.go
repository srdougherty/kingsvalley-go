package main

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type game struct {
	reader         *bufio.Reader
	players        []player
	curPlayerIndex int
	curPlayer      *player
	board          gameBoard
}

func newGame(reader bufio.Reader) *game {
	var g game

	g.reader = &reader

	g.players = []player{player{"Player 1", player1Char}, player{"Player 2", player2Char}}
	g.curPlayerIndex = 0
	g.curPlayer = &g.players[g.curPlayerIndex]

	g.board = newBoard()

	return &g
}

func (g *game) changeToNextPlayer() {
	g.curPlayerIndex++
	if g.curPlayerIndex >= len(g.players) {
		g.curPlayerIndex = 0
	}
	g.curPlayer = &g.players[g.curPlayerIndex]
}

func (g *game) parseMoveInput(userInput string) (rowIndex int, colIndex int, directionLetter string, err error) {
	curMove := strings.ToLower(userInput)
	if len(curMove) < 3 {
		return -1, -1, "", errors.New("Invalid format, expected CRD (C)olumn(R)ow(D)irection, ex: A5U")
	}

	colLetter := curMove[0]
	rowNum := curMove[1:2]
	directionLetter = curMove[2:3]

	colIndex = int(colLetter - 'a')
	rowIndex, _ = strconv.Atoi(rowNum)
	rowIndex--

	if rowIndex < 0 || rowIndex >= len(g.board.cellData) {
		return -1, -1, "", errors.New("Invalid row=" + strconv.Itoa(rowIndex+1))
	}
	cellRow := g.board.cellData[rowIndex]

	if colIndex < 0 || colIndex >= len(cellRow) {
		return -1, -1, "", errors.New("Invalid column=" + strconv.Itoa(colIndex+1))
	}

	return rowIndex, colIndex, directionLetter, nil
}

func (g *game) moveCurPlayer(rowIndex int, colIndex int, directionLetter string) (err error) {
	b := &g.board

	cellRow := b.cellData[rowIndex]

	selectedCell := cellRow[colIndex]
	switch selectedCell.value {
	case cellValueEmpty:
		return errors.New("Cell is empty, select another cell")
	case cellValuePlayer1Pawn, cellValuePlayer1King:
		if g.curPlayerIndex != 0 {
			return errors.New("Cell occupied by other player, select another cell")
		}
	case cellValuePlayer2Pawn, cellValuePlayer2King:
		if g.curPlayerIndex != 1 {
			return errors.New("Cell occupied by other player, select another cell")
		}
	}

	switch directionLetter {
	case "u":
		b.MovePieceUp(rowIndex, colIndex)
	case "d":
		b.MovePieceDown(rowIndex, colIndex)
	case "l":
		b.MovePieceLeft(rowIndex, colIndex)
	case "r":
		b.MovePieceRight(rowIndex, colIndex)
	default:
		return errors.New("Invalid direction=" + directionLetter)
	}

	return nil
}

func (g *game) play() {
	b := &g.board
	reader := g.reader

	done := false
	for !done {
		b.Render()

		fmt.Print("[", g.curPlayer.name, " (", g.curPlayer.piece, ")] Enter move: ")
		userInput, _ := reader.ReadString('\n')
		userInput = strings.TrimRight(userInput, "\n")
		if userInput == "quit" {
			done = true
			break
		}

		rowIndex, colIndex, directionLetter, e := g.parseMoveInput(userInput)
		if e != nil {
			fmt.Println(e)
			continue
		}
		//fmt.Println("cell index=[", rowIndex, ",", colIndex, "], direction=", directionLetter)

		e = g.moveCurPlayer(rowIndex, colIndex, directionLetter)
		if e != nil {
			fmt.Println(e)
			continue
		}

		if b.isWinConditionFulfilled() {
			b.Render()
			fmt.Println(g.curPlayer.name, "wins")
			done = true
			break
		}

		g.changeToNextPlayer()
	}
}
