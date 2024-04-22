package utils

import (
	"slices"
	"testing"
)

func TestHash(t *testing.T) {

	var hash = []byte{
		23, 232, 141, 177, 135, 175, 214, 44, 22, 229,
		222, 191, 62, 101, 39, 205, 0, 107, 192, 18, 188,
		144, 181, 26, 129, 12, 216, 12, 45, 81, 31, 67,
	}

	if slices.Compare(Hash([]byte{0, 1, 2}, []byte{3, 4, 5}), hash) != 0 {
		t.Error("invalid hash")
	}

}

func TestBound(t *testing.T) {

	t.Run("within the bound", func(t *testing.T) {
		res := Bound(1, 0, 2)
		if res != 1 {
			t.Errorf("incorrect, got: %d, want: %d.", res, 1)
		}
	})

	t.Run("lower limit", func(t *testing.T) {
		res := Bound(0, 1, 2)
		if res != 1 {
			t.Errorf("incorrect, got: %d, want: %d.", res, 1)
		}
	})

	t.Run("upper limit", func(t *testing.T) {
		res := Bound(2, 0, 1)
		if res != 1 {
			t.Errorf("incorrect, got: %d, want: %d.", res, 1)
		}
	})

}

func TestPbcopy(t *testing.T) {
	if err := Pbcopy("one"); err != nil {
		t.Error(err)
	}
}
