package utils

import (
	"crypto/sha256"
	"io"
	"math"
	"os/exec"
)

func Hash(seed, mutable []byte) []byte {
	sum := sha256.Sum256(append(seed, mutable...))
	return sum[:]
}

func Bound(value, min, max int) int {
	return int(math.Min(math.Max(float64(value), float64(min)), float64(max)))
}

func Pbcopy(text string) error {
	cmd := exec.Command("pbcopy")

	var pipe io.WriteCloser
	var err error

	if pipe, err = cmd.StdinPipe(); err != nil {
		return err
	}

	if err = cmd.Start(); err != nil {
		return err
	}

	if _, err = pipe.Write([]byte(text)); err != nil {
		return err
	}

	if err = pipe.Close(); err != nil {
		return err
	}

	return cmd.Wait()
}
