package grid

import "fmt"

func ExampleAllEqual() {
	fmt.Println(AllEqual([]string{"X", "X", "X"}, "X"))
	fmt.Println(AllEqual([]string{"X", "O", "X"}, "X"))
	fmt.Println(AllEqual([]string{"X", "X", "X"}, "O"))
	fmt.Println(AllEqual([]string{}, "X"))

	// Output:
	// true
	// false
	// false
	// true
}

func ExampleContains() {
	cells := [][]string{{"X", ".", "O"}, {".", "X", "."}}
	fmt.Println(Contains(cells, "O"))
	fmt.Println(Contains(cells, "."))
	fmt.Println(Contains(cells, "Y"))

	// Output:
	// true
	// true
	// false
}
