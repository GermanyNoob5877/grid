// Package tictactoe enthält die Regeln von Tic Tac Toe.
// Es verwendet die Bibliothek `grid` für das Spielfeld.
package tictactoe

import "grid/grid"

// Diese Aufgaben setzen die Regeln von Tic Tac Toe um. Das Spielfeld hat drei
// Zeilen und drei Spalten; zwei Personen setzen abwechselnd ihr Symbol
// "X" oder "O" in ein freies Feld. Wer zuerst eine ganze Zeile, Spalte
// oder Diagonale mit dem eigenen Symbol füllt, gewinnt.
// Die Felder sind für die Eingabe von `1` bis `9` nummeriert:
//
//	1 2 3
//	4 5 6
//	7 8 9
//
// Das Spielfeld wird nur über die Funktionen der Bibliothek `grid`
// gelesen und geändert.

// Empty ist der Inhalt eines freien Felds.
const Empty = "."

// NewBoard liefert ein neues, leeres Spielfeld für Tic Tac Toe,
// in dem jedes Feld `Empty` enthält.
//
// TODO
func NewBoard() [][]string {
	return grid.New(3, 3, Empty)
}

//FieldToRow erwarte ein Field als Eingabe und gibt die Reihe des Feldes zurück
func FieldToRow(field int) int {
	if field == 1 || field == 2 || field == 3 {
		return 0
	}
	if field == 4 || field == 5 || field == 6 {
		return 1
	}
	if field == 7 || field == 8 || field == 9 {
		return 2
	}
	return 0
}

//FieldToCol erwarte ein Field als Eingabe und gibt die Spalte des Feldes zurück
func FieldToCol(field int) int {
	if field == 1 || field == 4 || field == 7 {
		return 0
	}
	if field == 2 || field == 5 || field == 8 {
		return 1
	}
	if field == 3 || field == 6 || field == 9 {
		return 2
	}
	return 0
}

// Place erwartet ein Spielfeld, die Nummer `field` eines Felds von `1` bis `9`
// und das Symbol `player` und setzt das Symbol in dieses Feld.
// Sie liefert `true`, wenn das geklappt hat, und `false`, wenn es das Feld
// nicht gibt oder es schon besetzt ist. Dann bleibt das Spielfeld unverändert.
//
// TODO
func Place(cells [][]string, field int, player string) bool {
	if field > 9 || field < 1 {
		return false
	}
	if cells[FieldToRow(field)][FieldToCol(field)] == "." {
		cells[FieldToRow(field)][FieldToCol(field)] = player
		return true
	}
	return false
}

// NextPlayer erwartet das Symbol, das gerade am Zug war, und liefert
// das Symbol, das als Nächstes am Zug ist: "O" nach "X" und "X" nach "O".
//
// TODO
func NextPlayer(player string) string {
	if player == "X" {
		return "O"
	}
	if player == "O" {
		return "X"
	}
	return "."
}

// Wins erwartet ein Spielfeld und ein Symbol und liefert `true`, wenn eine
// Zeile, eine Spalte oder eine Diagonale ganz mit diesem Symbol gefüllt ist.
//
// TODO
func Wins(cells [][]string, player string) bool {
	for i := range len(cells) {
		if grid.AllEqual(grid.Row(cells, i), player) {
			return true
		}
		if grid.AllEqual(grid.Col(cells, i), player) {
			return true
		}
	}
	if grid.AllEqual(grid.Diagonal(cells), player) {
		return true
	}
	if grid.AllEqual(grid.AntiDiagonal(cells), player) {
		return true
	}
	return false
}

// Winner erwartet ein Spielfeld und liefert das Symbol, das gewonnen hat,
// oder "", wenn noch niemand gewonnen hat.
//
// TODO
func Winner(cells [][]string) string {
	if cells[0][0] == cells[0][1] && cells[0][1] == cells[0][2] && cells[0][0] != "." { //Reihe 1 selbes symbol?
		return cells[0][0]
	}
	if cells[1][0] == cells[1][1] && cells[1][1] == cells[1][2] && cells[1][0] != "." { //Reihe 2 selbes Symbol?
		return cells[1][0]
	}
	if cells[2][0] == cells[2][1] && cells[2][1] == cells[2][2] && cells[2][0] != "." { //Reihe 3 selbes Symbol?
		return cells[2][0]
	}
	if cells[0][0] == cells[1][0] && cells[1][0] == cells[2][0] && cells[0][0] != "." { //Spalte 1 selbes Symbol?
		return cells[0][0]
	}
	if cells[0][1] == cells[1][1] && cells[1][1] == cells[2][1] && cells[0][1] != "." { //Spalte 2 selbes Symbol?
		return cells[0][1]
	}
	if cells[0][2] == cells[1][2] && cells[1][2] == cells[2][2] && cells[0][2] != "." { //Spalte 3 selbes Symbol?
		return cells[0][2]
	}
	if cells[0][0] == cells[1][1] && cells[1][1] == cells[2][2] && cells[0][0] != "." { //Diagonale Links Oben -> Rechts unten?
		return cells[0][0]
	}
	if cells[2][0] == cells[1][1] && cells[1][1] == cells[0][2] && cells[2][0] != "." { //Diagonale Links UNten -> Rechts Oben?
		return cells[2][0]
	}
	return ""
}

// IsFull erwartet ein Spielfeld und liefert `true`, wenn kein Feld mehr frei ist.
//
// TODO
func IsFull(cells [][]string) bool {
	for i := range 3 {
		for j := range 3 {
			if cells[i][j] == "." {
				return false
			}
		}
	}
	return true
}

// GameOver erwartet ein Spielfeld und liefert `true`, wenn das Spiel vorbei ist,
// weil jemand gewonnen hat oder kein Feld mehr frei ist.
//
// TODO
func GameOver(cells [][]string) bool {
	if Winner(cells) == "X" || Winner(cells) == "O" || IsFull(cells) == true {
		return true
	}
	return false
}
