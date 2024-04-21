package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"passgen/pkg/utils"
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

	var f_n = flag.Int("n", 10, "number of passwords ≥ 1")
	var f_l = flag.Int("l", 32, "length of password, 8 ≤ length ≤ 32")
	var f_o = flag.Bool("o", false, "print only one last password")
	var f_c = flag.Bool("c", false, "copy to clipboard")

	flag.Parse()

	number, length, one, clipboard := *f_n, *f_l, *f_o, *f_c

	number = utils.Bound(number, 1, math.MaxInt)
	length = utils.Bound(length, 8, 32)

	if clipboard && !one {
		number = 1
	}

	fmt.Fprint(os.Stderr, "Master Password: ")

	seed, err := term.ReadPassword(int(syscall.Stdin))

	if err != nil {
		panic(err)
	}

	var mutable = seed

	fmt.Println()

	for line := range number {

		mutable = utils.Hash(seed, mutable)

		for i, symbol := range mutable {
			mutable[i] = symbols[symbol%42]
		}

		if !one || number == line+1 {
			if !clipboard {
				fmt.Printf("#%d %s\n", line+1, mutable[:length])
			} else {
				if err := utils.Pbcopy(string(mutable[:length])); err != nil {
					panic(err)
				}
			}
		}

	}

}
