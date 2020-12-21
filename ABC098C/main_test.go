package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHonestly(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		n := 5
		row := "WEEWW"
		assert.Equal(t, 1, Honestly(n, row))
	})
	t.Run("Example 2", func(t *testing.T) {
		n := 12
		row := "WEWEWEEEWWWE"
		assert.Equal(t, 4, Honestly(n, row))
	})
	t.Run("Example 3", func(t *testing.T) {
		n := 8
		row := "WWWWWEEE"
		assert.Equal(t, 3, Honestly(n, row))
	})
}

func TestCumulativeSum(t *testing.T) {
	t.Run("Example 1", func(t *testing.T) {
		n := 5
		row := "WEEWW"
		assert.Equal(t, 1, CumulativeSum(n, row))
	})
	t.Run("Example 2", func(t *testing.T) {
		n := 12
		row := "WEWEWEEEWWWE"
		assert.Equal(t, 4, CumulativeSum(n, row))
	})
	t.Run("Example 3", func(t *testing.T) {
		n := 8
		row := "WWWWWEEE"
		assert.Equal(t, 3, CumulativeSum(n, row))
	})
	t.Run("Test 1", func(t *testing.T) {
		n := 2
		row := "WE"
		assert.Equal(t, 0, CumulativeSum(n, row))
	})
}
