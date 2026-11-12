package extensions

import (
	"fmt"

	"grid/grid"
)

func ExampleFreeFields() {
	fmt.Println(FreeFields([][]string{{"X", ".", "O"}, {".", "X", "."}, {".", ".", "X"}}))
	fmt.Println(FreeFields([][]string{{".", ".", "."}, {".", ".", "."}, {".", ".", "."}}))
	fmt.Println(FreeFields([][]string{{"X", "O", "X"}, {"X", "O", "O"}, {"O", "X", "X"}}))

	// Output:
	// [2 4 6 7 8]
	// [1 2 3 4 5 6 7 8 9]
	// []
}

func ExampleDropPiece() {
	cells := [][]string{{".", ".", ".", "."}, {".", ".", ".", "."}, {".", ".", ".", "."}}
	fmt.Println(DropPiece(cells, 1, "X"))
	fmt.Println(DropPiece(cells, 1, "O"))
	fmt.Println(DropPiece(cells, 2, "X"))
	fmt.Println(DropPiece(cells, 1, "X"))
	fmt.Println(DropPiece(cells, 1, "O"))
	fmt.Print(grid.Format(cells))

	// Output:
	// true
	// true
	// true
	// true
	// false
	// . X . .
	// . O . .
	// . X X .
}
