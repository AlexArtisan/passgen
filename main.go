package main

import (
	"crypto/sha256"
	"flag"
	"fmt"
	"math"
	"os"
	"syscall"

	"golang.org/x/term"
)

var symbols = [42]byte{
	'(', 'E', 'H', '7', '4', 'A', 'l',
	'F', '0', 's', '3', 'k', '1', '#',
	'+', 't', 'x', '-', 'm', '8', '?',
	'I', 'y', 'n', '^', 'w', '%', 'D',
	'_', '5', '!', 'C', 'z', '$', 'J',
	'6', '*', 'B', '2', ')', 'G', '9',
}

func main() {

	var num = flag.Int("n", 10, "number of passwords ≥ 1")
	var width = flag.Int("l", 32, "length of password, 8 ≤ length ≤ 32")
	var last = flag.Bool("o", false, "print only one last password")

	flag.Parse()

	*num = int(math.Max(float64(*num), 1))
	*width = int(math.Min(math.Max(float64(*width), 8), 32))

	fmt.Fprint(os.Stderr, "Master Password: ")

	seed, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		panic(err)
	}

	var mutable = append([]byte{}, seed...)

	fmt.Println()

	for line := range *num {

		sum := sha256.Sum256(append(seed, mutable...))

		for i, code := range sum {
			sum[i] = symbols[code%42]
		}

		mutable = append([]byte{}, sum[:]...)

		if !*last || *num == line+1 {
			fmt.Printf("#%d %s\n", line+1, sum[:*width])
		}

	}

}
