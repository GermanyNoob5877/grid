package main

import (
	"fmt"
	"grid/grid"
	"grid/tictactoe"
)

// Dieses Programm spielt eine Partie Tic Tac Toe aus einer Liste von Zügen nach,
// z. B. um eine Partie zu analysieren oder zu testen. Es verwendet die Pakete
// `grid` und `tictactoe`; das Programm `play` ist ähnlich aufgebaut.
// Die Züge in `main` sind vorgegeben; die Aufgabe ist `replay`.

func main() {
	replay([]int{5, 1, 9, 3, 2, 8, 7, 4, 6})
}

// replay erwartet eine Liste von Feldnummern und spielt sie als Züge nach,
// abwechselnd für "X" und "O", beginnend mit "X". Ist ein Zug nicht möglich,
// gibt sie z. B. "Ungültiger Zug: 5" aus, und mit dem nächsten Zug der Liste
// ist dasselbe Symbol noch einmal an der Reihe. Ist das Spiel vorbei, werden
// weitere Züge nicht mehr gespielt. Am Ende gibt sie das Spielfeld aus und
// darunter das Ergebnis: "X hat gewonnen!", "O hat gewonnen!",
// "Unentschieden." oder "Das Spiel ist noch nicht vorbei.".
//
// TODO
func replay(moves []int) {
	spielfeld := tictactoe.NewBoard()
	player := "X"
	Error := false
	for i := range len(moves) {
		if grid.Get(spielfeld, tictactoe.FieldToRow(moves[i]), tictactoe.FieldToCol(moves[i])) == "." && moves[i] < 10 && moves[i] > 0 {
			grid.Set(spielfeld, tictactoe.FieldToRow(moves[i]), tictactoe.FieldToCol(moves[i]), player)
			player = tictactoe.NextPlayer(player)
		} else {
			if grid.Get(spielfeld, tictactoe.FieldToRow(moves[i]), tictactoe.FieldToCol(moves[i])) != "." || moves[i] < 10 || moves[i] > 0 {
				fmt.Println("Ungültiger Zug:", moves[i])
				if Error == false {
					Error = true
				}
			}
		}
		if tictactoe.GameOver(spielfeld) == true && Error == false {
			if tictactoe.Wins(spielfeld, "X") {
				fmt.Println(grid.Format(spielfeld))
				fmt.Println("X hat gewonnen!")
				break
			}
			if tictactoe.Wins(spielfeld, "O") {
				fmt.Println(grid.Format(spielfeld))
				fmt.Println("O hat gewonnen!")
				break
			}
			if tictactoe.IsFull(spielfeld) {
				fmt.Println(grid.Format(spielfeld))
				fmt.Println("Unentschieden.")
				break
			}
		} else {
			if i == len(moves) {
				fmt.Println("Das Spiel ist noch nicht vorbei.")
			}
		}
	}
	if Error == true {
		fmt.Println(grid.Format(spielfeld))
		fmt.Println("Das Spiel ist noch nicht vorbei.")
	}
}
