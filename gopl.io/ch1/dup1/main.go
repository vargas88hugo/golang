// Dup1 prints the text of each line that appears more than
// once in the standard input, proceded by its count
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // create a new empty map

	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		// line := input.Text()
		// counts[line] = counts[line] + 1
		counts[input.Text()]++
	}

	for line, n := range counts {
		if n >= 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
