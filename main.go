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

func hash(seed, mutable []byte) []byte {
	sum := sha256.Sum256(append(seed, mutable...))
	return sum[:]
}

func bound(value, min, max int) int {
	return int(math.Min(math.Max(float64(value), float64(min)), float64(max)))
}

func main() {

	var f_n = flag.Int("n", 10, "number of passwords ≥ 1")
	var f_l = flag.Int("l", 32, "length of password, 8 ≤ length ≤ 32")
	var f_o = flag.Bool("o", false, "print only one last password")

	flag.Parse()

	number, length, one := *f_n, *f_l, *f_o

	number = bound(number, 1, math.MaxInt)
	length = bound(length, 8, 32)

	fmt.Fprint(os.Stderr, "Master Password: ")

	seed, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		panic(err)
	}

	var mutable = seed

	fmt.Println()

	for line := range number {

		mutable = hash(seed, mutable)

		for i, symbol := range mutable {
			mutable[i] = symbols[symbol%42]
		}

		if !one || number == line+1 {
			fmt.Printf("#%d %s\n", line+1, mutable[:length])
		}

	}

}
