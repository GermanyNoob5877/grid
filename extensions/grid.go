package extensions

// Diese Aufgaben ergänzen Funktionen für Spielfelder. Verwenden Sie die
// Funktionen der Bibliothek `grid`, statt sie nachzubauen.

// CountValue erwartet ein Spielfeld und einen Wert und liefert,
// wie viele Zellen diesen Wert enthalten.
//
// TODO
func CountValue(cells [][]string, value string) int {
	counter := 0
	for i := range len(cells) {
		for j := range len(cells[0]) {
			if cells[i][j] == value {
				counter = counter + 1
			}
		}
	}
	return counter
}

// Transpose erwartet ein Spielfeld und liefert ein neues Spielfeld,
// in dem die Spalten zu Zeilen geworden sind: Die erste Spalte von `cells`
// ist die erste Zeile des Ergebnisses und so weiter. Aus zwei Zeilen
// mit drei Spalten werden so drei Zeilen mit zwei Spalten.
//
// TODO
func Transpose(cells [][]string) [][]string {
	TransposedArray := make([][]string, len(cells[0]))
	for i := range TransposedArray {
		TransposedArray[i] = make([]string, len(cells))
	}
	for i := range len(TransposedArray) {
		for j := range len(TransposedArray[0]) {
			TransposedArray[i][j] = cells[j][i]
		}
	}
	return TransposedArray
}

// ContainsRun erwartet eine Linie, einen Wert und eine Länge `length` ab `1`
// und liefert `true`, wenn die Linie mindestens `length` Zellen mit diesem
// Wert direkt hintereinander enthält. Bei Vier Gewinnt gewinnt z. B., wer
// vier eigene Steine hintereinander in einer Linie hat.
//
// TODO
func ContainsRun(line []string, value string, length int) bool {
	counter := 0
	for i := range len(line) {
		if line[i] == value {
			counter += 1
		} else {
			counter = 0
		}
		if counter == length {
			return true
		}
	}
	return false
}
