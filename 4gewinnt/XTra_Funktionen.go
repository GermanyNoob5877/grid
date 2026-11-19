package main

import (
	"grid/extensions"
	"grid/grid"
)

func GameOver(cells [][]string) bool {
	//if WinForO(cells) {
	//	return true
	//}
	//if WinForX(cells) {
	//	return true
	//}
	if SixBySevenIsFull(cells) {
		return true
	}
	return false
}

func WinForX(cells [][]string) bool {
	for i := range 7 {
		if extensions.ContainsRun(grid.Col(cells, i), "X", 4) {
			return true
		}
	}
	for j := range 6 {
		if extensions.ContainsRun(grid.Row(cells, j), "X", 4) {
			return true
		}
	}
	return false
}

func WinForO(cells [][]string) bool {
	for i := range 7 {
		if extensions.ContainsRun(grid.Row(cells, i), "O", 4) {
			return true
		}
		if extensions.ContainsRun(grid.Col(cells, i), "O", 4) {
			return true
		}
	}
	return false
}

func SixBySevenIsFull(cells [][]string) bool {
	for i := range 6 {
		for j := range 7 {
			if cells[i][j] == "." {
				return false
			}
		}
	}
	return true
}

func DiagonalWinner(cells [][]string, value string, length int) bool {
	return true
}
