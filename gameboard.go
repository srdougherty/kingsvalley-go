package main

import (
	"fmt"
	"strconv"
	"strings"
)

//    A     B     C     D     E
//          10        20        30
// 1234567890123456789012345678901
// -------------------------------01
// |     |     |  ^  |     |     |02   1
// |  O  |  O  |  O  |  O  |  O  |03
// -------------------------------04
// |     |     |     |     |     |05   2
// |     |     |     |     |     |06
// -------------------------------07
// |     |     ||---||     |     |08   3
// |     |     ||---||     |     |09
// -------------------------------10
// |     |     |     |     |     |11   4
// |     |     |     |     |     |12
// -------------------------------13
// |     |     |  ^  |     |     |14   5
// |  X  |  X  |  X  |  X  |  X  |15
// -------------------------------16

const player1Char = "X"
const player2Char = "O"
const crownChar = "^"
const horizontalBorderChar = "-"
const verticalBorderChar = "|"
const emptyChar = " "
const cellCountW = 5
const cellCountH = 5
const cellWidth = 5    // characters (excluding border)
const cellHeight = 2   // lines (excluding border)
const borderWidth = 1  // characters
const borderHeight = 1 // characters

type cellType int

const (
	cellTypeNormal cellType = iota
	cellTypeMaster
)

type cellValue int

const (
	cellValueEmpty cellValue = iota
	cellValuePlayer1Pawn
	cellValuePlayer1King
	cellValuePlayer2Pawn
	cellValuePlayer2King
)

type cell struct {
	value cellValue
	kind  cellType
}

type gameBoard struct {
	cellData [cellCountH][cellCountW]cell
}

func newBoard() gameBoard {
	var b gameBoard

	for r := range b.cellData {
		for _, cell := range b.cellData[r] {
			cell.kind = cellTypeNormal
			cell.value = cellValueEmpty
		}
	}

	b.cellData[2][2].kind = cellTypeMaster

	b.cellData[0][0].value = cellValuePlayer2Pawn
	b.cellData[0][1].value = cellValuePlayer2Pawn
	b.cellData[0][2].value = cellValuePlayer2King
	b.cellData[0][3].value = cellValuePlayer2Pawn
	b.cellData[0][4].value = cellValuePlayer2Pawn

	b.cellData[4][0].value = cellValuePlayer1Pawn
	b.cellData[4][1].value = cellValuePlayer1Pawn
	b.cellData[4][2].value = cellValuePlayer1King
	b.cellData[4][3].value = cellValuePlayer1Pawn
	b.cellData[4][4].value = cellValuePlayer1Pawn

	return b
}

func (b *gameBoard) Width() int { // grid width in characters
	return (cellWidth * cellCountW) + (borderWidth * (cellCountW + 1))
}

func (b *gameBoard) Height() int { // grid height in lines
	return (cellHeight * cellCountH) + (borderHeight * (cellCountH + 1))
}

func (b *gameBoard) Render() {
	borderLine := strings.Repeat(horizontalBorderChar, b.Width())
	emptyCell := strings.Repeat(emptyChar, cellWidth)

	for i := 0; i < cellCountW; i++ {
		fmt.Print(strings.Repeat(emptyChar, borderWidth+cellCountW%2+1))
		fmt.Print(string('A' + i))
		fmt.Print(strings.Repeat(emptyChar, cellCountW%2+1))
	}
	fmt.Println()

	fmt.Println(borderLine)
	for r := range b.cellData {
		topLine := verticalBorderChar
		bottomLine := verticalBorderChar

		for _, cell := range b.cellData[r] {
			cellTop := ""
			cellBottom := ""

			switch cell.kind {
			case cellTypeNormal:
				cellTop += emptyCell
				cellBottom += emptyCell
			case cellTypeMaster:
				cellTop += "|---|"
				cellBottom += "|---|"
			}

			switch cell.value {
			case cellValuePlayer1King, cellValuePlayer2King:
				cellTop = cellTop[:2] + crownChar + cellTop[2+1:]
			}

			switch cell.value {
			case cellValuePlayer1Pawn, cellValuePlayer1King:
				cellBottom = cellTop[:2] + player1Char + cellTop[2+1:]
			case cellValuePlayer2Pawn, cellValuePlayer2King:
				cellBottom = cellTop[:2] + player2Char + cellTop[2+1:]
			}

			topLine += cellTop + verticalBorderChar
			bottomLine += cellBottom + verticalBorderChar
		}

		topLine += strings.Repeat(emptyChar, 2) + strconv.Itoa(r+1)
		fmt.Println(topLine)
		fmt.Println(bottomLine)
		fmt.Println(borderLine)
	}
}

func (b *gameBoard) MovePieceUp(rowIndex int, colIndex int) {
	//fmt.Println("-- move up -- initialRowIndex=", rowIndex)
	for rowIndex > 0 {
		curRow := &b.cellData[rowIndex]
		curCell := &curRow[colIndex]
		if curCell.value == cellValueEmpty {
			fmt.Println("curCell empty [", rowIndex, ",", colIndex, "]")
			break
		}

		nextRow := &b.cellData[rowIndex-1]
		nextCell := &nextRow[colIndex]
		if nextCell.value != cellValueEmpty {
			//fmt.Println("nextCell not empty [", rowIndex-1, ",", colIndex, "]")
			break
		}

		nextCell.value = curCell.value
		curCell.value = cellValueEmpty

		rowIndex--
		//fmt.Println("rowIndex=", rowIndex)
	}
}

func (b *gameBoard) MovePieceDown(rowIndex int, colIndex int) {
	//fmt.Println("-- move down -- initialRowIndex=", rowIndex)
	for rowIndex < len(b.cellData)-1 {
		curRow := &b.cellData[rowIndex]
		curCell := &curRow[colIndex]
		if curCell.value == cellValueEmpty {
			fmt.Println("curCell empty [", rowIndex, ",", colIndex, "]")
			break
		}

		nextRow := &b.cellData[rowIndex+1]
		nextCell := &nextRow[colIndex]
		if nextCell.value != cellValueEmpty {
			//fmt.Println("nextCell not empty [", rowIndex+1, ",", colIndex, "]")
			break
		}

		nextCell.value = curCell.value
		curCell.value = cellValueEmpty

		rowIndex++
		//fmt.Println("rowIndex=", rowIndex)
	}
}

func (b *gameBoard) MovePieceLeft(rowIndex int, colIndex int) {
	//fmt.Println("-- move left -- initialColIndex=", colIndex)
	for colIndex > 0 {
		curRow := &b.cellData[rowIndex]
		curCell := &curRow[colIndex]
		if curCell.value == cellValueEmpty {
			fmt.Println("curCell empty [", rowIndex, ",", colIndex, "]")
			break
		}

		nextRow := &b.cellData[rowIndex]
		nextCell := &nextRow[colIndex-1]
		if nextCell.value != cellValueEmpty {
			//fmt.Println("nextCell not empty [", rowIndex, ",", colIndex-1, "]")
			break
		}

		nextCell.value = curCell.value
		curCell.value = cellValueEmpty

		colIndex--
		//fmt.Println("colIndex=", colIndex)
	}
}

func (b *gameBoard) MovePieceRight(rowIndex int, colIndex int) {
	//fmt.Println("-- move right -- initialColIndex=", colIndex)
	for colIndex < len(b.cellData[rowIndex])-1 {
		curRow := &b.cellData[rowIndex]
		curCell := &curRow[colIndex]
		if curCell.value == cellValueEmpty {
			fmt.Println("curCell empty [", rowIndex, ",", colIndex, "]")
			break
		}

		nextRow := &b.cellData[rowIndex]
		nextCell := &nextRow[colIndex+1]
		if nextCell.value != cellValueEmpty {
			//fmt.Println("nextCell not empty [", rowIndex, ",", colIndex+1, "]")
			break
		}

		nextCell.value = curCell.value
		curCell.value = cellValueEmpty

		colIndex++
		//fmt.Println("colIndex=", colIndex)
	}
}

func (b *gameBoard) printStatus() {
	for r := range b.cellData {
		for c, cell := range b.cellData[r] {
			fmt.Println("[", r, ",", c, "]=", cell.value)
		}
	}
	fmt.Println("-------------------------------------")
}

func (b *gameBoard) isWinConditionFulfilled() bool {
	return (b.cellData[2][2].value == cellValuePlayer1King) || (b.cellData[2][2].value == cellValuePlayer2King)
}
