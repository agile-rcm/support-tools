package tools_test

import (
	"fmt"
	"git.agiletech.de/AgileRCM/support-tools/tools"
)

func ExampleLevenshtein() {
	fmt.Println(tools.Levenshtein([]rune(string("GE (German Restricted)")), []rune(string("GE (German Restricted )"))))
	// Output: 1
}
