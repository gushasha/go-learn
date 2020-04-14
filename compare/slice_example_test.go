package compare_test

import (
	"fmt"
	"github.com/gushasha/go-learn/compare"
)

func ExampleStrings() {
	strs1 := []string{"Hello", "world"}
	strs2 := []string{"Hello", "Go"}
	if compare.Strings(strs1, strs2) {
		fmt.Println("equal")
	} else {
		fmt.Println("not equal")
	}
}
