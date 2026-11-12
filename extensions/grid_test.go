package extensions

import "fmt"

func ExampleCountValue() {
	cells := [][]string{{"X", ".", "O"}, {".", "X", "."}, {".", ".", "X"}}
	fmt.Println(CountValue(cells, "X"))
	fmt.Println(CountValue(cells, "O"))
	fmt.Println(CountValue(cells, "."))
	fmt.Println(CountValue(cells, "Y"))

	// Output:
	// 3
	// 1
	// 5
	// 0
}

func ExampleTranspose() {
	fmt.Println(Transpose([][]string{{"a", "b", "c"}, {"d", "e", "f"}}))
	fmt.Println(Transpose([][]string{{"a", "b"}, {"c", "d"}}))

	// Output:
	// [[a d] [b e] [c f]]
	// [[a c] [b d]]
}

func ExampleContainsRun() {
	line := []string{"X", "X", ".", "X", "X", "X"}
	fmt.Println(ContainsRun(line, "X", 3))
	fmt.Println(ContainsRun(line, "X", 4))
	fmt.Println(ContainsRun(line, ".", 1))
	fmt.Println(ContainsRun(line, "O", 1))
	fmt.Println(ContainsRun([]string{}, "X", 1))

	// Output:
	// true
	// false
	// true
	// false
	// false
}
