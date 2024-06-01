package Ascii

import (
	"fmt"
	"os"
	"strings"
)

func HandleSpecialCase(s string) string {
	cases := []string{"\\a", "\\b", "\\t", "\\f", "\\r", "\\v"}

	for i := 0; i < len(cases); i++ {
		if strings.Contains(s, cases[i]) {
			fmt.Printf("Special Case \"%s\"  is not supported\n", cases[i])
			os.Exit(1)
		}
	}
	return s
}
