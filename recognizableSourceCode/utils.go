package recognizableSourceCode

import (
	"fmt"
)

// PrintArray will print an array of strings
func PrintArray(things []string) {
	for _, thing := range things {
		fmt.Println(thing)
	}
	// also there is a comment here too
}
