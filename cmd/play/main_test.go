package main

func Example_printResult() {
	printResult([][]string{{"X", "X", "X"}, {"O", "O", "."}, {".", ".", "."}})
	printResult([][]string{{"X", "O", "."}, {"X", "O", "."}, {".", "O", "X"}})
	printResult([][]string{{"X", "O", "X"}, {"X", "O", "O"}, {"O", "X", "X"}})

	// Output:
	// X hat gewonnen!
	// O hat gewonnen!
	// Unentschieden.
}
