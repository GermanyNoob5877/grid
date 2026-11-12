# Spielfeld ohne Structs

Ein Spielfeld aus Zeilen und Spalten ist hier eine Liste von Zeilen,
jede Zeile eine Liste von Zellen, also `[][]string`.
Alle Funktionen erhalten das Spielfeld als Parameter.

In dieser Einheit entsteht eine Bibliothek für solche Spielfelder und darauf
ein Paket mit den Regeln von Tic Tac Toe.
Außerdem gibt es Programme, die beides verwenden.

## Aufbau

- [`grid`](grid): die Bibliothek für Spielfelder.
- [`tictactoe`](tictactoe): die Regeln von Tic Tac Toe, aufgebaut auf `grid`.
- [`extensions`](extensions): weitere Funktionen, die auf beiden aufbauen.
- [`cmd`](cmd): die Programme, eines pro Verzeichnis.

## Aufgaben

Für jede Funktion sind Name, Parameter und eine Beschreibung im Kommentar
vorgegeben, die erwarteten Ergebnisse stehen in den Tests.
Zu schreiben ist jeweils der Rumpf der Funktion.
Empfohlene Reihenfolge:

1. [`grid/cells.go`](grid/cells.go): `New`, `Rows`, `Cols`, `Get`, `Set`
   und `Format`, ein Spielfeld anlegen, lesen, ändern und ausgeben.
2. [`grid/lines.go`](grid/lines.go): `Row`, `Col`, `Diagonal` und
   `AntiDiagonal`, eine Zeile, Spalte oder Diagonale als Liste.
3. [`grid/checks.go`](grid/checks.go): `AllEqual` und `Contains`.
4. [`tictactoe/tictactoe.go`](tictactoe/tictactoe.go): `NewBoard`, `Place`,
   `NextPlayer`, `Wins`, `Winner`, `IsFull` und `GameOver`.
5. [`cmd/play/main.go`](cmd/play/main.go): `printResult`.
   Der Ablauf des Spiels in `main` ist vorgegeben.
6. [`cmd/replay/main.go`](cmd/replay/main.go): `replay`, eine Partie
   aus einer Liste von Zügen nachspielen. Die Züge in `main` sind vorgegeben.
7. [`extensions/grid.go`](extensions/grid.go): `CountValue`, `Transpose`
   und `ContainsRun`.
8. [`extensions/games.go`](extensions/games.go): `FreeFields`,
   die freien Felder bei Tic Tac Toe.
9. Weiterführend: `DropPiece` in [`extensions/games.go`](extensions/games.go),
   einen Stein bei Vier Gewinnt einwerfen.

Die Tests späterer Aufgaben schlagen fehl, bis diese gelöst sind.
Einzelne Tests lassen sich gezielt ausführen,
z. B. mit `go test ./grid -run ExampleNew`.

## Hinweise

- Zeilen und Spalten werden ab `0` gezählt.
  Die Felder von Tic Tac Toe sind für die Eingabe von `1` bis `9` nummeriert.
- Eine Liste als Parameter verweist auf dieselben Zellen wie beim Aufrufer.
  Eine Funktion wie `Set` ändert deshalb das Spielfeld des Aufrufers,
  ohne etwas zurückzugeben.

## Ausführen

Im Verzeichnis dieser Einheit (dort, wo die Datei `go.mod` liegt):

- `go test ./...`: führt alle Tests aus.
- `go run ./cmd/play`: startet ein Spiel für zwei Personen.
- `go run ./cmd/replay`: spielt die Partie aus `main` nach.
