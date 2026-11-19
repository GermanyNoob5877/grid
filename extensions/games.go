package extensions

// Diese Aufgaben verwenden die Bibliothek `grid` und das Paket `tictactoe`
// für weitere Spiele: einen Überblick über die freien Felder,
// z. B. für einen Computergegner, und den ersten Schritt zu Vier Gewinnt.

// FreeFields erwartet ein Spielfeld für Tic Tac Toe und liefert die Nummern
// der freien Felder von `1` bis `9` in aufsteigender Reihenfolge.
// Für ein leeres Spielfeld sind das alle Nummern von `1` bis `9`.
//
// TODO
func FreeFields(cells [][]string) []int {
	Feldnummer := 1
	Ausgabe := []int{}
	for i := range len(cells) {
		for j := range len(cells[0]) {
			if cells[i][j] == "." {
				Ausgabe = append(Ausgabe, Feldnummer)
			}
			Feldnummer += 1
		}
	}
	return Ausgabe
}

// DropPiece erwartet ein Spielfeld für Vier Gewinnt, die Nummer `col` einer
// Spalte, gezählt ab `0`, und ein Symbol `piece`. Wie bei Tic Tac Toe enthalten freie Felder
// `tictactoe.Empty`. Bei Vier Gewinnt fällt ein Stein in einer Spalte nach
// unten, bis er auf den Boden oder einen anderen Stein trifft. Die Funktion
// setzt das Symbol deshalb in das unterste freie Feld der Spalte und liefert
// `true`. Ist die Spalte schon voll, liefert sie `false`.
//
// Weiterführende Aufgabe.
//
// TODO
func DropPiece(cells [][]string, col int, piece string) bool {
	for i := range len(cells) {
		if i == len(cells)-1 && cells[i][col] == "." {
			cells[i][col] = piece
			return true
		} else {
			if cells[i][col] != "." && i > 0 {
				if cells[i-1][col] == "." {
					cells[(i - 1)][col] = piece
					return true
				}
			}
		}
	}
	return false
}
