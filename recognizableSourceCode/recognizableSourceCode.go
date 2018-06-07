package recognizableSourceCode

import "fmt"

// PrintNames will print the names
func PrintNames() {
	names := []string{
		"jack",
		"bob",
		"john",
		"elvis",
		"piper",
		"wren",
	}
	for i, name := range names {
		fmt.Println(name, i)
	}
}
