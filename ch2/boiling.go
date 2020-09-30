package ch2

import (
	"fmt"
	"os"
)

const boilingF = 212.0

func main() {
	var f = boilingF         // aaaa
	var c = (f - 32) * 5 / 9 // bbbb

	var a int
	var b int
	fmt.Fprintln(os.Stdout, &a, &b)

	fmt.Printf("boiling point = %gF or %gC\n", f, c)

}
