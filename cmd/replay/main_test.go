package main

func Example_replay_win() {
	replay([]int{1, 4, 2, 5, 3, 6})

	// Output:
	// X X X
	// O O .
	// . . .
	// X hat gewonnen!
}

func Example_replay_draw() {
	replay([]int{1, 2, 3, 5, 4, 6, 8, 7, 9})

	// Output:
	// X O X
	// X O O
	// O X X
	// Unentschieden.
}

func Example_replay_invalidMoves() {
	replay([]int{5, 5, 1, 0, 2, 3})

	// Output:
	// Ungültiger Zug: 5
	// Ungültiger Zug: 0
	// O X O
	// . X .
	// . . .
	// Das Spiel ist noch nicht vorbei.
}
