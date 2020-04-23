// Experiment to measure the difference in running time between our
package exercise_1_3

import (
	"fmt"
	"io"
	"os"
	"strings"
)

var out io.Writer = os.Stdout

// One First approach
func One() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Fprintln(out, s)
}

// Two Second approach
func Two() {
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	fmt.Fprintln(out, s)
}

// Three Third approach
func Three() {
	fmt.Fprintln(out, strings.Join(os.Args[1:], " "))
}
