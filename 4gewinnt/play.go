package main

import (
	"fmt"
	"grid/extensions"
	"grid/grid"
	"grid/tictactoe"
)

func main() {
	spielfeld := grid.New(6, 7, ".")
	player := "X"
	var Eingabe = 0
	for GameOver(spielfeld) != true {
		fmt.Println("In welche Spalte möchtest du einen Stein (" + player + ") legen?")
		fmt.Scan(&Eingabe)
		extensions.DropPiece(spielfeld, Eingabe-1, player)
		player = tictactoe.NextPlayer(player)
		fmt.Println(grid.Format(spielfeld))
	}
	//Hier gewinnansagen und so reinhauen
}
