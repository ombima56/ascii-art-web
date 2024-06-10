package Ascii

import (
	"fmt"
	"os"
	"strings"
)

// PrintBanner prints the input string using the loaded banner characters.
func PrintBanner(line, filename string) string {
	outPut := make([][]string, 8) // Output slice to store the banner lines.

	banner := LoadBanner(filename) // Load the banner characters

	for _, char := range line {

		if char < 32 || char > 126 {
			fmt.Printf("Character out of range:%q\n", char)
			os.Exit(1)
		}
		if ascii, Ok := banner[char]; Ok {

			// If the character is found, split it into lines and append to the output
			asciiLines := strings.Split(ascii, "\n")
			for i := 0; i < len(asciiLines); i++ {
				outPut[i] = append(outPut[i], asciiLines[i])
			}

		} else {
			// If the character is not found, print an error message and continue
			fmt.Printf("Charachter not found: %q\n", char)
			continue
		}
	}
	result := string('\n')
	// Print the assembled output lines
	for _, line := range outPut {
		result += (strings.Join(line, "")) + string('\n')
	}

	return result
}
