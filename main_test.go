package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMax(t *testing.T) {

	t.Run("Test positif case to Find Max", func(t *testing.T) {

		a := []int{1, 2, 3, 8, 9, 3, 2, 1}
		max := FindMax(a)

		n := 3
		if max != n {
			t.Errorf("findMax was incorrect, got: %d, want: %d.", max, n)
		}
		assert.Equal(t, n, max)
	})
	t.Run("Test negative case to Find Max", func(t *testing.T) {

		a := []int{1, 2, 3, 8, 9, 3, 2, 1}
		max := FindMax(a)

		n := 2
		assert.NotEqual(t, n, max)
	})

}
